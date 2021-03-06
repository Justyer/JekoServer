package world

import (
	"errors"
	"net"

	prtx "github.com/Justyer/JekoServer/tcp/model/protox"
)

const (
	// MaxRoom : 房间最大人数
	MaxRoom = 2
)

var (
	// JekoWorld : 世界实例
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

// World : 游戏世界
type World struct {
	RoomList []*prtx.RoomInfo
}

// XreateWorld : 实例化
func XreateWorld() *World {
	return &World{}
}

// GetRoomInfoByID : 获取房间信息
func (w *World) GetRoomInfoByID(id int32) (*prtx.RoomInfo, error) {
	for i := 0; i < len(w.RoomList); i++ {
		if w.RoomList[i].ID == id {
			return w.RoomList[i], nil
		}
	}
	return nil, errors.New("room not found")
}

// GetConnAtRoom : 获取房间里的用户的所有连接
func (w *World) GetConnAtRoom(id int32) ([]net.Conn, error) {
	ri, err := w.GetRoomInfoByID(id)
	if err != nil {
		return nil, errors.New("room has been disband")
	}
	var conns []net.Conn
	for _, r := range ri.UserList {
		conns = append(conns, r.Conn)
	}
	return conns, nil
}
