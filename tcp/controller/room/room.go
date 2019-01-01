package room

import (
	"github.com/Justyer/JekoServer/plugin/log"
	"github.com/Justyer/JekoServer/tcp/controller/base"
	"github.com/Justyer/JekoServer/tcp/model/prt"
	"github.com/Justyer/JekoServer/tcp/service/room"
	"github.com/Justyer/lingo/bytes"
	"github.com/golang/protobuf/proto"
)

type roomController struct {
	base.BaseController
}

func NewRoomController() *roomController {
	return &roomController{}
}

func (self *roomController) QueryRoomList() {
	var req prt.QueryRoomListReq
	proto.Unmarshal(self.DataPack.Data, &req)

	// var err error
	room := room.NewRoomService()
	room_list, _ := room.QueryRoomList()

	var resp prt.QueryRoomListResp
	resp.Code = 0
	resp.RoomList = room_list
	respByte, err := proto.Marshal(&resp)
	if err != nil {
		log.Err(err.Error())
	}
	log.Tx("[req_data]: %s", req.String())
	log.Tx("[resp_data]: %s", resp.String())
	len_byte := bytes.ToByteForLE(int32(len(respByte)))

	var resp_final_byte []byte

	resp_final_byte = bytes.Extend(resp_final_byte, bytes.ToByteForLE(self.DataPack.MsgType))
	resp_final_byte = bytes.Extend(resp_final_byte, bytes.ToByteForLE(uint16(prt.MsgCmd_value["Room_QueryListResp"])))
	resp_final_byte = bytes.Extend(resp_final_byte, len_byte)
	resp_final_byte = bytes.Extend(resp_final_byte, respByte)

	if _, err := self.Conn.Write(resp_final_byte); err != nil {
		log.Err("[write err]:", err)
	}
	log.Succ("[resp_final_byte]: %v", resp_final_byte)
}

func (self *roomController) GetIn() {
	var req prt.GetInReq
	proto.Unmarshal(self.DataPack.Data, &req)

	var err error
	room := room.NewRoomService()
	room_info, err := room.GetIn_insert(req.ID, self.Cache)

	var resp prt.GetInResp
	if err != nil {
		resp.Code = 1
	}
	resp.Room = room_info
	respByte, err := proto.Marshal(&resp)
	if err != nil {
		log.Err(err.Error())
	}
	log.Tx("[req_data]: %s", req.String())
	log.Tx("[resp_data]: %s", resp.String())
	len_byte := bytes.ToByteForLE(int32(len(respByte)))

	var resp_final_byte []byte

	resp_final_byte = bytes.Extend(resp_final_byte, bytes.ToByteForLE(self.DataPack.MsgType))
	resp_final_byte = bytes.Extend(resp_final_byte, bytes.ToByteForLE(uint16(prt.MsgCmd_value["Room_GetInResp"])))
	resp_final_byte = bytes.Extend(resp_final_byte, len_byte)
	resp_final_byte = bytes.Extend(resp_final_byte, respByte)

	room.GetIn_ff(req.ID, resp_final_byte)
}

func (self *roomController) EnterReady() {

}
