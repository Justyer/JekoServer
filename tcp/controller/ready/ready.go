package ready

import (
	"github.com/Justyer/JekoServer/plugin/log"
	"github.com/Justyer/JekoServer/tcp/controller/base"
	"github.com/Justyer/JekoServer/tcp/model/prt"
	"github.com/Justyer/JekoServer/tcp/service/ready"
	"github.com/Justyer/lingo/bytes"
	"github.com/golang/protobuf/proto"
)

type readyController struct {
	base.BaseController
}

func NewReadyController() *readyController {
	return &readyController{}
}

func (self *readyController) ReadyInfo() int32 {
	var req prt.GetReadyInfoReq
	proto.Unmarshal(self.DataPack.Data, &req)

	// var err error
	ready := ready.NewReadyService()
	user_list, _ := ready.ReadyInfo(self.Cache)

	var resp prt.GetReadyInfoResp
	resp.Code = 0
	resp.UserList = user_list
	respByte, err := proto.Marshal(&resp)
	if err != nil {
		log.Err(err.Error())
	}
	log.Tx("[req_data]: %s", req.String())
	log.Tx("[resp_data]: %s", resp.String())
	len_byte := bytes.ToByteForLE(int32(len(respByte)))

	var resp_final_byte []byte

	resp_final_byte = bytes.Extend(resp_final_byte, bytes.ToByteForLE(self.DataPack.MsgType))
	resp_final_byte = bytes.Extend(resp_final_byte, bytes.ToByteForLE(uint16(prt.MsgCmd_value["Ready_GetInfoResp"])))
	resp_final_byte = bytes.Extend(resp_final_byte, len_byte)
	resp_final_byte = bytes.Extend(resp_final_byte, respByte)

	if _, err := self.Conn.Write(resp_final_byte); err != nil {
		log.Err("[write err]:", err)
	}
	log.Succ("[resp_final_byte]: %v", resp_final_byte)

	return int32(-1)
}
