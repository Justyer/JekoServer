package ready

import (
	"github.com/Justyer/JekoServer/plugin/log"
	"github.com/Justyer/JekoServer/tcp/model/cache"
	"github.com/Justyer/JekoServer/tcp/model/prt"
	"github.com/Justyer/JekoServer/tcp/model/tool"
	"github.com/Justyer/JekoServer/tcp/model/weapon"
)

type readyServe struct {
}

func NewReadyService() *readyServe {
	return &readyServe{}
}

func (self *readyServe) ReadyInfo(c *tool.Cache) ([]*prt.UserInfoDTO, error) {
	users := cache.RoomMap[c.User.CurRoom].Users
	var us []*prt.UserInfoDTO

	for _, u := range users {
		var user prt.UserInfoDTO
		user.UserName = u.UserName
		user.IconURL = u.IconURL
		user.AttributeNum = int32(50)
		var wp prt.WeaponDTO
		wp.ID = 1
		wp.SN = 2
		user.Weapon = &wp
		us = append(us, &user)
	}
	return us, nil
}

func (self *readyServe) AddWeaponAddr(ws []*prt.WeaponExtraAttrDTO, c *tool.Cache) {
	var ea []*weapon.WeaponExtraAttr
	for _, w := range ws {
		var wea weapon.WeaponExtraAttr
		wea.AttrType = w.AttrType
		wea.Value = w.Value
		ea = append(ea, &wea)
	}
	c.User.Weapon.ExtraAttrs = ea
}

func (self *readyServe) Distribute(id int32, b []byte) {
	for _, u := range cache.RoomMap[id].Users {
		if _, err := u.Conn.Write(b); err != nil {
			log.Err("[write err]:", err)
		}
		log.Succ("[resp_final_byte]: %v", b)
	}
}

func (self *readyServe) GetRoomID(c *tool.Cache) int32 {
	return c.User.CurRoom
}
