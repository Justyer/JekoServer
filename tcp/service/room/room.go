package room

import (
	prt "github.com/Justyer/JekoServer/tcp/model/proto"
	prtx "github.com/Justyer/JekoServer/tcp/model/protox"
	"github.com/Justyer/JekoServer/tcp/model/world"
	"github.com/Justyer/JekoServer/tcp/service/base"
	"github.com/Justyer/jie"
)

type roomService struct {
	base.BaseService
}

func NewRoomService(c *jie.Context) *roomService {
	s := &roomService{}
	s.Ctx = c
	return s
}

// 查询房间列表
func (self *roomService) QueryRoomList() []*prt.RoomInfo {
	var ris []*prt.RoomInfo
	for _, room := range world.JekoWorld.RoomList {
		var ri prt.RoomInfo
		var ul []*prt.UserInfo
		for _, ui := range room.UserList {
			ul = append(ul, &ui.UserInfo)
		}
		ri = room.RoomInfo
		ri.UserList = ul
		ris = append(ris, &ri)
	}
	return ris
}

// 根据房间id将玩家放入房间
func (self *roomService) GetIn(id int32) (*prt.RoomInfo, error) {
	u := self.Ctx.Get("user").(prtx.UserInfo)
	world.JekoWorld.RoomList[id].UserList = append(world.JekoWorld.RoomList[id].UserList, &u)

	var ri prt.RoomInfo
	var us []*prt.UserInfo
	for _, ul := range world.JekoWorld.RoomList[id].UserList {
		us = append(us, &ul.UserInfo)
	}
	ri = world.JekoWorld.RoomList[id].RoomInfo
	ri.UserList = us
	return &ri, nil
}
