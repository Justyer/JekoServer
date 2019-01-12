package ready

import (
	"time"

	"github.com/Justyer/jie"

	"github.com/Justyer/JekoServer/plugin/log"
	prt "github.com/Justyer/JekoServer/tcp/model/proto"
	"github.com/Justyer/JekoServer/tcp/service/ready"
)

// 进入准备界面后，给房间内的每个人发一份所有人的信息
func ReadyInfoController(c *jie.Context) {
	var req prt.GetReadyInfoReq
	if err := c.BindProtoReq(&req); err != nil {
		log.Err("[ReadyInfoController bind_req_err]: %s", err.Error())
		return
	}
	log.Tx("[ReadyInfoController req]: %s", req.String())

	svc := ready.NewReadyService(c)
	user_list := svc.GetAllUserInRoom()

	var resp prt.GetReadyInfoResp
	resp.Code = 0
	resp.UserList = user_list
	data_b, err := c.PackProtoResp(&resp)
	if err != nil {
		log.Err("[ReadyInfoController pack_resp_err]: %s", err.Error())
		return
	}
	log.Tx("[ReadyInfoController resp]: %s", resp.String())

	final_b := svc.PackProtocol(uint16(prt.MsgCmd_value["Ready_GetInfoResp"]), data_b)

	c.Send(final_b)

	c.Redirect(uint16(prt.MsgType_value["Ready"]), uint16(prt.MsgCmd_value["Ready_PrepareCombatReq"]))
}

// 保存武器属性
func PrepareCombatController(c *jie.Context) {
	var req prt.PrepareCombatReq
	if err := c.BindProtoReq(&req); err != nil {
		log.Err("[PrepareCombatController bind_req_err]: %s", err.Error())
		return
	}
	log.Tx("[PrepareCombatController req]: %s", req.String())

	svc := ready.NewReadyService(c)
	svc.AddWeaponAddr(req.WeaponExtraAttrList)

	var resp prt.PrepareCombatResp
	data_b, err := c.PackProtoResp(&resp)
	if err != nil {
		log.Err("[PrepareCombatController pack_resp_err]: %s", err.Error())
		return
	}
	log.Tx("[PrepareCombatController resp]: %s", resp.String())

	final_b := svc.PackProtocol(uint16(prt.MsgCmd_value["Ready_PrepareCombatResp"]), data_b)

	room_id := svc.GetRoomID()
	// 某一个人提交天赋后，进入游戏的倒计时变为5秒
	time.AfterFunc(5*time.Second, func() {
		conns := svc.GetAllConnInRoom(room_id)
		c.Broadcast(final_b, conns)
	})
}
