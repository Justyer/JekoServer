package room

import (
	"errors"

	"github.com/Justyer/JekoServer/plugin/log"
	"github.com/Justyer/JekoServer/tcp/model/cache"
	"github.com/Justyer/JekoServer/tcp/model/prt"
	"github.com/Justyer/JekoServer/tcp/model/tool"
)

type roomServe struct {
}

func NewRoomService() *roomServe {
	return &roomServe{}
}

func (self *roomServe) QueryRoomList() ([]*prt.RoomInfoDTO, error) {
	var room_list []*prt.RoomInfoDTO
	for _, r := range cache.RoomMap {
		var r_dto prt.RoomInfoDTO

		r_dto.ID = r.ID
		for _, u := range r.Users {
			var user prt.UserInfoDTO
			user.UserName = u.UserName
			user.IconURL = u.IconURL
			r_dto.UserList = append(r_dto.UserList, &user)
		}

		room_list = append(room_list, &r_dto)
	}
	return room_list, nil
}

func (self *roomServe) GetIn_insert(id int32, c *tool.Cache) (*prt.RoomInfoDTO, error) {
	cache.RoomMap[id].Users = append(cache.RoomMap[id].Users, c.User)

	if r, ok := cache.RoomMap[id]; !ok {
		return nil, errors.New("room_id not exist")
	} else {
		var ri prt.RoomInfoDTO
		ri.ID = r.ID
		for _, u := range r.Users {
			var user prt.UserInfoDTO
			user.UserName = u.UserName
			user.IconURL = u.IconURL
			ri.UserList = append(ri.UserList, &user)
		}
		return &ri, nil
	}
}

func (self *roomServe) GetIn_ff(id int32, b []byte) {
	for _, u := range cache.RoomMap[id].Users {
		if _, err := u.Conn.Write(b); err != nil {
			log.Err("[write err]:", err)
		}
		log.Succ("[resp_final_byte]: %v", b)
	}
}

func (self *roomServe) EnterReady() {

}
