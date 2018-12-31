package router

import (
	"net"

	ptl "github.com/Justyer/JekoServer/plugin/protocol"
	"github.com/Justyer/JekoServer/tcp/controller/auth"
	"github.com/Justyer/JekoServer/tcp/model/cst"
)

type MzRouter struct {
	DataPack *ptl.DataPack_2_2_4
	Conn     net.Conn
}

func NewMzRouter() *MzRouter {
	return &MzRouter{}
}

func (self *MzRouter) MsgType() {
	switch int32(self.DataPack.MsgType) {
	case cst.MsgType_value["Login"]:
		self.MsgCmd_Login()
		// case cst.MsgType_value["Room"]:
		// 	RoomCmd(dp)
		// case cst.MsgType_value["Combat"]:
		// 	ComBatCmd(dp)
	}
}

func (self *MzRouter) MsgCmd_Login() {
	switch int32(self.DataPack.MsgCmd) {
	case cst.MsgCmd_value["Login_LoginReq"]:
		login := auth.NewLoginController()
		login.DataPack = self.DataPack
		login.Conn = self.Conn
		login.Login()
	}
}
