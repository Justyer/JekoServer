package auth

import (
	"github.com/Justyer/JekoServer/plugin/log"
	ptl "github.com/Justyer/JekoServer/plugin/protocol"
	"github.com/Justyer/JekoServer/tcp/model/auth"
	"github.com/Justyer/JekoServer/tcp/model/cst"
	"github.com/Justyer/JekoServer/tcp/model/user"
	auth_svc "github.com/Justyer/JekoServer/tcp/service/auth"
	"github.com/Justyer/lingo/bytes"
	"github.com/golang/protobuf/proto"
)

const (
	MacLogin  int32 = 0
	UserLogin int32 = 1
)

type LoginController struct {
	DataPack *ptl.DataPack_2_2_4
}

func NewLoginController() *LoginController {
	return &LoginController{}
}

func (self *LoginController) Login() []byte {
	var req auth.LoginReq
	proto.Unmarshal(self.DataPack.Data, &req)

	login := auth_svc.NewLoginService()
	var err error
	var u user.User

	switch req.Type {
	case MacLogin:
		u, err = login.MacLogin(req.MAC)
	case UserLogin:
		u, err = login.UserLogin(req.UserName, req.PassWord)
	}

	// 封装要相应的数据
	var resp auth.LoginResp
	if err != nil {
		resp.Code = 1
	}
	resp.UserID = u.UserID
	resp.UserName = u.UserName
	respByte, err := proto.Marshal(&resp)
	if err != nil {
		log.Err(err.Error())
	}
	len_byte := bytes.ToByteForLE(int32(len(respByte)))

	var resp_final_byte []byte

	resp_final_byte = bytes.Extend(resp_final_byte, bytes.ToByteForLE(self.DataPack.MsgType))
	resp_final_byte = bytes.Extend(resp_final_byte, bytes.ToByteForLE(uint16(cst.MsgCmd_value["Login_LoginResp"])))
	resp_final_byte = bytes.Extend(resp_final_byte, len_byte)
	resp_final_byte = bytes.Extend(resp_final_byte, respByte)

	return resp_final_byte
}
