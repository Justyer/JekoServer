package auth

import (
	"github.com/Justyer/JekoServer/plugin/log"
	"github.com/Justyer/JekoServer/tcp/controller/base"
	"github.com/Justyer/JekoServer/tcp/model/prt"
	"github.com/Justyer/JekoServer/tcp/model/user"
	auth_svc "github.com/Justyer/JekoServer/tcp/service/auth"
	"github.com/Justyer/lingo/bytes"
	"github.com/golang/protobuf/proto"
)

const (
	MacLogin  int32 = 0
	UserLogin int32 = 1
)

type loginController struct {
	base.BaseController
}

func NewLoginController() *loginController {
	return &loginController{}
}

func (self *loginController) Login() {
	var req prt.LoginReq
	proto.Unmarshal(self.DataPack.Data, &req)

	var err error
	var u user.User
	login := auth_svc.NewLoginService()

	switch req.Type {
	case MacLogin:
		u, err = login.MacLogin(req.MAC)
	case UserLogin:
		u, err = login.UserLogin(req.UserName, req.PassWord)
	}

	u.Conn = self.Conn

	self.Cache.User = &u

	// 封装要相应的数据
	var resp prt.LoginResp
	if err != nil {
		resp.Code = 1
	}
	resp.UserID = u.UserID
	resp.UserName = u.UserName
	respByte, err := proto.Marshal(&resp)
	if err != nil {
		log.Err(err.Error())
	}
	log.Tx("[req_data]: %s", req.String())
	log.Tx("[resp_data]: %s", resp.String())
	len_byte := bytes.ToByteForLE(int32(len(respByte)))

	var resp_final_byte []byte

	resp_final_byte = bytes.Extend(resp_final_byte, bytes.ToByteForLE(self.DataPack.MsgType))
	resp_final_byte = bytes.Extend(resp_final_byte, bytes.ToByteForLE(uint16(prt.MsgCmd_value["Login_LoginResp"])))
	resp_final_byte = bytes.Extend(resp_final_byte, len_byte)
	resp_final_byte = bytes.Extend(resp_final_byte, respByte)

	if _, err := self.Conn.Write(resp_final_byte); err != nil {
		log.Err("[write err]:", err)
	}
	log.Succ("[resp_final_byte]: %v", resp_final_byte)
}
