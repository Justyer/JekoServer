package router

import (
	"github.com/Justyer/JekoServer/tcp/controller/auth"
	"github.com/Justyer/JekoServer/tcp/controller/ready"
	"github.com/Justyer/JekoServer/tcp/controller/room"
	prt "github.com/Justyer/JekoServer/tcp/model/proto"
	"github.com/Justyer/jie"
)

// CreateJekoRouter : 加载路由
func CreateJekoRouter(j *jie.Engine) {
	// 登录
	j.Router.GET(auth.Login, uint16(prt.MsgType_value["Login"]), uint16(prt.MsgCmd_value["Login_LoginReq"]))

	// 房间
	j.Router.GET(room.QueryRoomList, uint16(prt.MsgType_value["Room"]), uint16(prt.MsgCmd_value["Room_QueryListReq"]))
	j.Router.GET(room.GetIn, uint16(prt.MsgType_value["Room"]), uint16(prt.MsgCmd_value["Room_GetInReq"]))
	j.Router.GET(room.EnterReady, uint16(prt.MsgType_value["Room"]), uint16(prt.MsgCmd_value["Room_EnterReadyReq"]))

	// 准备阶段
	j.Router.GET(ready.ReadyInfo, uint16(prt.MsgType_value["Ready"]), uint16(prt.MsgCmd_value["Ready_GetInfoReq"]))
	j.Router.GET(ready.PrepareCombat, uint16(prt.MsgType_value["Ready"]), uint16(prt.MsgCmd_value["Ready_PrepareCombatReq"]))
}
