package base

import (
	"net"

	"github.com/Justyer/JekoServer/tcp/model/world"
	"github.com/Justyer/jie"
	protocol "github.com/Justyer/jie/protocol/ptl_2_2_4"
	"github.com/Justyer/lingo/bytes"
)

// BaseService : 基础服务
type BaseService struct {
	Ctx *jie.Context
}

// NewBaseService : 实例化
func NewBaseService() *BaseService {
	return &BaseService{}
}

// PackProtocol : 封装数据包
func (svc *BaseService) PackProtocol(cmd uint16, dataB []byte) []byte {
	p := svc.Ctx.DP.(*protocol.Protocol)

	typeB := bytes.ToByteForLE(p.MsgType)
	cmdB := bytes.ToByteForLE(cmd)
	lenB := bytes.ToByteForLE(int32(len(dataB)))
	finalB := bytes.Merge(typeB, cmdB, lenB, dataB)
	return finalB
}

// GetAllConnInRoom : 获取房间内所有用户的连接
func (svc *BaseService) GetAllConnInRoom(id int32) []net.Conn {
	var ns []net.Conn
	for i := 0; i < len(world.JekoWorld.RoomList[id].UserList); i++ {
		ns = append(ns, world.JekoWorld.RoomList[id].UserList[i].Conn)
	}
	return ns
}

// GetRoomID : 根据上下文获取房间号
func (svc *BaseService) GetRoomID() int32 {
	return svc.Ctx.Get("room_id").(int32)
}
