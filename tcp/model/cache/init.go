package cache

import (
	"github.com/Justyer/JekoServer/tcp/model/room"
)

// 房间信息
var (
	RoomMap = make(map[int32]*room.RoomInfo)
)

func init() {
	initRoomMap()
}

// 初始化所有房间信息
func initRoomMap() {
	for i := 0; i < 10; i++ {
		var ri room.RoomInfo
		ri.ID = int32(i)

		RoomMap[int32(i)] = &ri
	}

}
