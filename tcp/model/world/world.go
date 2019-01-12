package world

import (
	"errors"
	"net"

	prtx "github.com/Justyer/JekoServer/tcp/model/protox"
)

const (
	// 房间最大人数
	MaxRoom = 2
)

var (
	JekoWorld *World
)

func init() {
	JekoWorld = XreateWorld()
	var ris []*prtx.RoomInfo
	for i := 0; i < 10; i++ {
		var ri prtx.RoomInfo
		ri.ID = int32(i)
		ris = append(ris, &ri)
	}
	JekoWorld.RoomList = ris
}

type World struct {
	RoomList []*prtx.RoomInfo
}

func XreateWorld() *World {
	return &World{}
}

// 获取房间信息
func (self *World) GetRoomInfoByID(id int32) (*prtx.RoomInfo, error) {
	for i := 0; i < len(self.RoomList); i++ {
		if self.RoomList[i].ID == id {
			return self.RoomList[i], nil
		}
	}
	return nil, errors.New("room not found")
}

// 获取房间里的用户的所有连接
func (self *World) GetConnAtRoom(id int32) ([]net.Conn, error) {
	ri, err := self.GetRoomInfoByID(id)
	if err != nil {
		return nil, errors.New("room has been disband")
	}
	var conns []net.Conn
	for _, r := range ri.UserList {
		conns = append(conns, r.Conn)
	}
	return conns, nil
}
