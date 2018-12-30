package router

import (
	"fmt"

	ptl "github.com/Justyer/JekoServer/plugin/protocol"
	"github.com/Justyer/JekoServer/tcp/controller/auth"
	"github.com/Justyer/JekoServer/tcp/model/cst"
)

type MzRouter struct {
	DataPack *ptl.DataPack_2_2_4
}

func NewMzRouter() *MzRouter {
	return &MzRouter{}
}

func (self *MzRouter) MsgType() []byte {
	fmt.Println("fdfdf", self.DataPack)
	var resp []byte
	switch int32(self.DataPack.MsgType) {
	case cst.MsgType_value["Login"]:
		resp = self.MsgCmd_Login()
		// case cst.MsgType_value["Room"]:
		// 	RoomCmd(dp)
		// case cst.MsgType_value["Combat"]:
		// 	ComBatCmd(dp)
	}
	return resp
}

func (self *MzRouter) MsgCmd_Login() []byte {
	var resp []byte
	switch int32(self.DataPack.MsgCmd) {
	case cst.MsgCmd_value["Login_LoginReq"]:
		login := auth.NewLoginController()
		login.DataPack = self.DataPack
		resp = login.Login()
	}
	return resp
}
