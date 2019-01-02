package router

import (
	"net"

	"github.com/Justyer/JekoServer/tcp/controller/ready"

	"github.com/Justyer/JekoServer/tcp/controller/room"
	"github.com/Justyer/JekoServer/tcp/model/prt"
	"github.com/Justyer/JekoServer/tcp/model/tool"

	"github.com/Justyer/JekoServer/plugin/log"
	ptl "github.com/Justyer/JekoServer/plugin/protocol"
	"github.com/Justyer/JekoServer/tcp/controller/auth"
)

type JekoRouter struct {
	DataPack *ptl.DataPack_2_2_4
	Conn     net.Conn
	Cache    *tool.Cache
}

func NewJekoRouter() *JekoRouter {
	return &JekoRouter{}
}

// 一级命令
func (self *JekoRouter) MsgType() {
	switch int32(self.DataPack.MsgType) {
	case prt.MsgType_value["Login"]:
		self.MsgCmd_Login()
	case prt.MsgType_value["Room"]:
		self.MsgCmd_Room()
	case prt.MsgType_value["Ready"]:
		self.MsgCmd_Ready()
	// case cst.MsgType_value["Combat"]:
	// 	ComBatCmd(dp)
	default:
		log.Err("[unkown msg_type]: %d %d %d %v", self.DataPack.MsgType, self.DataPack.MsgCmd, self.DataPack.DataLen, self.DataPack.Data)
	}
}

// 登录的二级命令
func (self *JekoRouter) MsgCmd_Login() {
	switch int32(self.DataPack.MsgCmd) {
	case prt.MsgCmd_value["Login_LoginReq"]:
		login := auth.NewLoginController()
		login.DataPack = self.DataPack
		login.Conn = self.Conn
		login.Cache = self.Cache
		login.Login()
	default:
		log.Err("[unkown msg_cmd_login]: %d %d %d %v", self.DataPack.MsgType, self.DataPack.MsgCmd, self.DataPack.DataLen, self.DataPack.Data)
	}
}

// 房间操作的二级命令
func (self *JekoRouter) MsgCmd_Room() {
	room := room.NewRoomController()
	room.DataPack = self.DataPack
	room.Conn = self.Conn
	room.Cache = self.Cache

	cmd := int32(self.DataPack.MsgCmd)
	for {
		switch cmd {
		case prt.MsgCmd_value["Room_QueryListReq"]:
			cmd = room.QueryRoomList()
		case prt.MsgCmd_value["Room_GetInReq"]:
			cmd = room.GetIn()
		case prt.MsgCmd_value["Room_EnterReadyReq"]:
			cmd = room.EnterReady()
		default:
			log.Err("[unkown msg_cmd_room]: %d %d %d %v", self.DataPack.MsgType, self.DataPack.MsgCmd, self.DataPack.DataLen, self.DataPack.Data)
		}
		if cmd == int32(-1) {
			break
		}
	}
}

// 准备界面的二级命令
func (self *JekoRouter) MsgCmd_Ready() {
	ready := ready.NewReadyController()
	ready.DataPack = self.DataPack
	ready.Conn = self.Conn
	ready.Cache = self.Cache

	cmd := int32(self.DataPack.MsgCmd)
	for {
		switch cmd {
		case prt.MsgCmd_value["Ready_GetInfoReq"]:
			cmd = ready.ReadyInfo()
		case prt.MsgCmd_value["Ready_PrepareCombatReq"]:
			// cmd = ready.GetIn()
		default:
			log.Err("[unkown msg_cmd_room]: %d %d %d %v", self.DataPack.MsgType, self.DataPack.MsgCmd, self.DataPack.DataLen, self.DataPack.Data)
		}
		if cmd == int32(-1) {
			break
		}
	}
}
