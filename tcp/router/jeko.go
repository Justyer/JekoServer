package router

import (
	"github.com/Justyer/JekoServer/tcp/controller/auth"
	"github.com/Justyer/JekoServer/tcp/controller/ready"
	"github.com/Justyer/JekoServer/tcp/controller/room"
	prt "github.com/Justyer/JekoServer/tcp/model/proto"
	"github.com/Justyer/jie"
)

func CreateJekoRouter(j *jie.Engine) {
	// 登录
	j.Router.GET(auth.LoginController, uint16(prt.MsgType_value["Login"]), uint16(prt.MsgCmd_value["Login_LoginReq"]))

	// 房间
	j.Router.GET(room.QueryRoomListController, uint16(prt.MsgType_value["Room"]), uint16(prt.MsgCmd_value["Room_QueryListReq"]))
	j.Router.GET(room.GetInController, uint16(prt.MsgType_value["Room"]), uint16(prt.MsgCmd_value["Room_GetInReq"]))
	j.Router.GET(room.EnterReadyController, uint16(prt.MsgType_value["Room"]), uint16(prt.MsgCmd_value["Room_EnterReadyReq"]))

	// 准备阶段
	j.Router.GET(ready.ReadyInfoController, uint16(prt.MsgType_value["Ready"]), uint16(prt.MsgCmd_value["Ready_GetInfoReq"]))
	j.Router.GET(ready.PrepareCombatController, uint16(prt.MsgType_value["Ready"]), uint16(prt.MsgCmd_value["Ready_PrepareCombatReq"]))
}
