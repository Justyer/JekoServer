package room

import (
	"time"

	"github.com/Justyer/jie"

	"github.com/Justyer/JekoServer/plugin/log"
	prt "github.com/Justyer/JekoServer/tcp/model/proto"
	"github.com/Justyer/JekoServer/tcp/model/world"
	"github.com/Justyer/JekoServer/tcp/service/room"
)

// 查询房间列表
func QueryRoomListController(c *jie.Context) {
	var req prt.QueryRoomListReq
	if err := c.BindProtoReq(&req); err != nil {
		log.Err("[QueryRoomListController bind_req_err]: %s", err.Error())
		return
	}
	log.Tx("[QueryRoomListController req]: %s", req.String())

	svc := room.NewRoomService(c)
	room_list := svc.QueryRoomList()

	var resp prt.QueryRoomListResp
	resp.Code = 0
	resp.RoomList = room_list
	data_b, err := c.PackProtoResp(&resp)
	if err != nil {
		log.Err("[QueryRoomListController pack_resp_err]: %s", err.Error())
		return
	}
	log.Tx("[QueryRoomListController resp]: %s", resp.String())

	final_b := svc.PackProtocol(uint16(prt.MsgCmd_value["Room_QueryListResp"]), data_b)

	c.Send(final_b)
}

// 进入房间
func GetInController(c *jie.Context) {
	var req prt.GetInRoomReq
	if err := c.BindProtoReq(&req); err != nil {
		log.Err("[GetInController bind_req_err]: %s", err.Error())
		return
	}
	log.Tx("[GetInController req]: %s", req.String())

	svc := room.NewRoomService(c)
	room_info, err := svc.GetIn(req.ID)

	var resp prt.GetInRoomResp
	if err != nil {
		resp.Code = 1
	}
	c.Put("room_id", req.ID)
	resp.Room = room_info
	data_byte, err := c.PackProtoResp(&resp)
	if err != nil {
		log.Err("[GetInController pack_resp_err]: %s", err.Error())
		return
	}
	log.Tx("[GetInController resp]: %s", resp.String())

	final_b := svc.PackProtocol(uint16(prt.MsgCmd_value["Room_GetInResp"]), data_byte)

	conns := svc.GetAllConnInRoom(req.ID)
	c.Broadcast(final_b, conns)

	if len(room_info.UserList) >= world.MaxRoom {
		c.Redirect(uint16(prt.MsgType_value["Room"]), uint16(prt.MsgCmd_value["Room_EnterReadyReq"]))
	}
}

// 进入准备阶段
func EnterReadyController(c *jie.Context) {
	log.Tx("[EnterReadyController req]: %s", "")

	svc := room.NewRoomService(c)
	room_id := svc.GetRoomID()

	var resp prt.EnterReadyResp
	resp.Code = 0
	data_byte, err := c.PackProtoResp(&resp)
	if err != nil {
		log.Err("[EnterReadyController pack_resp_err]: %s", err.Error())
		return
	}
	log.Tx("[EnterReadyController resp]: %s", resp.String())

	final_b := svc.PackProtocol(uint16(prt.MsgCmd_value["Room_EnterReadyResp"]), data_byte)

	// 所有人都进入房间后，2秒后进入准备界面
	time.AfterFunc(2*time.Second, func() {
		conns := svc.GetAllConnInRoom(room_id)
		c.Broadcast(final_b, conns)
	})
}
