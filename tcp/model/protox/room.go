package prtx

import (
	prt "github.com/Justyer/JekoServer/tcp/model/proto"
)

type RoomInfo struct {
	prt.RoomInfo

	UserList []*UserInfo
}

func NewRoomDao() *RoomInfo {
	return &RoomInfo{}
}
