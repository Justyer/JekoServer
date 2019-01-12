package base

import (
	"net"

	"github.com/Justyer/JekoServer/tcp/model/world"
	"github.com/Justyer/jie"
	protocol "github.com/Justyer/jie/protocol/ptl_2_2_4"
	"github.com/Justyer/lingo/bytes"
)

type BaseService struct {
	Ctx *jie.Context
}

func NewBaseService() *BaseService {
	return &BaseService{}
}

func (self *BaseService) PackProtocol(cmd uint16, data_b []byte) []byte {
	p := self.Ctx.DP.(*protocol.Protocol)

	type_b := bytes.ToByteForLE(p.MsgType)
	cmd_b := bytes.ToByteForLE(cmd)
	len_b := bytes.ToByteForLE(int32(len(data_b)))
	final_b := bytes.Merge(type_b, cmd_b, len_b, data_b)
	return final_b
}

func (self *BaseService) GetAllConnInRoom(id int32) []net.Conn {
	var ns []net.Conn
	for i, _ := range world.JekoWorld.RoomList[id].UserList {
		ns = append(ns, world.JekoWorld.RoomList[id].UserList[i].Conn)
	}
	return ns
}

func (self *BaseService) GetRoomID() int32 {
	return self.Ctx.Get("room_id").(int32)
}
