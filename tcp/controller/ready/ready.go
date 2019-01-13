package ready

import (
	"time"

	"github.com/Justyer/jie"

	"github.com/Justyer/JekoServer/plugin/log"
	prt "github.com/Justyer/JekoServer/tcp/model/proto"
	"github.com/Justyer/JekoServer/tcp/service/ready"
)

// ReadyInfo : 进入准备界面后，给房间内的每个人发一份所有人的信息
func ReadyInfo(c *jie.Context) {
	var req prt.GetReadyInfoReq
	if err := c.BindProtoReq(&req); err != nil {
		log.Err("[ReadyInfo bind_req_err]: %s", err.Error())
		return
	}
	log.Tx("[ReadyInfo req]: %s", req.String())

	svc := ready.NewReadyService(c)
	userList := svc.GetAllUserInRoom()

	var resp prt.GetReadyInfoResp
	resp.Code = 0
	resp.UserList = userList
	dataB, err := c.PackProtoResp(&resp)
	if err != nil {
		log.Err("[ReadyInfo pack_resp_err]: %s", err.Error())
		return
	}
	log.Tx("[ReadyInfo resp]: %s", resp.String())

	finalB := svc.PackProtocol(uint16(prt.MsgCmd_value["Ready_GetInfoResp"]), dataB)

	c.Send(finalB)

	// c.Redirect(uint16(prt.MsgType_value["Ready"]), uint16(prt.MsgCmd_value["Ready_GetInfoReq"]))
}

// PrepareCombat : 保存武器属性
func PrepareCombat(c *jie.Context) {
	var req prt.PrepareCombatReq
	if err := c.BindProtoReq(&req); err != nil {
		log.Err("[PrepareCombat bind_req_err]: %s", err.Error())
		return
	}
	log.Tx("[PrepareCombat req]: %s", req.String())

	svc := ready.NewReadyService(c)
	svc.AddWeaponAddr(req.WeaponExtraAttrList)

	var resp prt.PrepareCombatResp
	dataB, err := c.PackProtoResp(&resp)
	if err != nil {
		log.Err("[PrepareCombat pack_resp_err]: %s", err.Error())
		return
	}
	log.Tx("[PrepareCombat resp]: %s", resp.String())

	finalB := svc.PackProtocol(uint16(prt.MsgCmd_value["Ready_PrepareCombatResp"]), dataB)

	roomID := svc.GetRoomID()
	// 某一个人提交天赋后，进入游戏的倒计时变为5秒
	time.AfterFunc(5*time.Second, func() {
		conns := svc.GetAllConnInRoom(roomID)
		c.Broadcast(finalB, conns)
	})
}
