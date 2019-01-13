package room

import (
	"time"

	"github.com/Justyer/jie"

	"github.com/Justyer/JekoServer/plugin/log"
	prt "github.com/Justyer/JekoServer/tcp/model/proto"
	"github.com/Justyer/JekoServer/tcp/model/world"
	"github.com/Justyer/JekoServer/tcp/service/room"
)

// QueryRoomList : 查询房间列表
func QueryRoomList(c *jie.Context) {
	var req prt.QueryRoomListReq
	if err := c.BindProtoReq(&req); err != nil {
		log.Err("[QueryRoomList bind_req_err]: %s", err.Error())
		return
	}
	log.Tx("[QueryRoomList req]: %s", req.String())

	svc := room.NewRoomService(c)
	roomList := svc.QueryRoomList()

	var resp prt.QueryRoomListResp
	resp.Code = 0
	resp.RoomList = roomList
	dataB, err := c.PackProtoResp(&resp)
	if err != nil {
		log.Err("[QueryRoomList pack_resp_err]: %s", err.Error())
		return
	}
	log.Tx("[QueryRoomList resp]: %s", resp.String())

	finalB := svc.PackProtocol(uint16(prt.MsgCmd_value["Room_QueryListResp"]), dataB)

	c.Send(finalB)
}

// GetIn : 进入房间
func GetIn(c *jie.Context) {
	var req prt.GetInRoomReq
	if err := c.BindProtoReq(&req); err != nil {
		log.Err("[GetIn bind_req_err]: %s", err.Error())
		return
	}
	log.Tx("[GetIn req]: %s", req.String())

	svc := room.NewRoomService(c)
	roomInfo, err := svc.GetIn(req.ID)

	var resp prt.GetInRoomResp
	if err != nil {
		resp.Code = 1
	}
	c.Put("room_id", req.ID)
	resp.Room = roomInfo
	dataB, err := c.PackProtoResp(&resp)
	if err != nil {
		log.Err("[GetIn pack_resp_err]: %s", err.Error())
		return
	}
	log.Tx("[GetIn resp]: %s", resp.String())

	finalB := svc.PackProtocol(uint16(prt.MsgCmd_value["Room_GetInResp"]), dataB)

	conns := svc.GetAllConnInRoom(req.ID)
	c.Broadcast(finalB, conns)

	if len(roomInfo.UserList) >= world.MaxRoom {
		c.Redirect(uint16(prt.MsgType_value["Room"]), uint16(prt.MsgCmd_value["Room_EnterReadyReq"]))
	}
}

// EnterReady : 进入准备阶段
func EnterReady(c *jie.Context) {
	log.Tx("[EnterReady req]: %s", "")

	svc := room.NewRoomService(c)
	roomID := svc.GetRoomID()

	var resp prt.EnterReadyResp
	resp.Code = 0
	dataB, err := c.PackProtoResp(&resp)
	if err != nil {
		log.Err("[EnterReady pack_resp_err]: %s", err.Error())
		return
	}
	log.Tx("[EnterReady resp]: %s", resp.String())

	finalB := svc.PackProtocol(uint16(prt.MsgCmd_value["Room_EnterReadyResp"]), dataB)

	// 所有人都进入房间后，2秒后进入准备界面
	time.AfterFunc(2*time.Second, func() {
		conns := svc.GetAllConnInRoom(roomID)
		c.Broadcast(finalB, conns)
	})
}
