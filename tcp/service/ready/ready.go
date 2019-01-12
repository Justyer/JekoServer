package ready

import (
	prt "github.com/Justyer/JekoServer/tcp/model/proto"
	prtx "github.com/Justyer/JekoServer/tcp/model/protox"
	"github.com/Justyer/JekoServer/tcp/model/world"
	"github.com/Justyer/JekoServer/tcp/service/base"
	"github.com/Justyer/jie"
)

type readyService struct {
	base.BaseService
}

func NewReadyService(c *jie.Context) *readyService {
	s := &readyService{}
	s.Ctx = c
	return s
}

func (self *readyService) GetAllUserInRoom() []*prt.UserInfo {
	us_x := world.JekoWorld.RoomList[self.Ctx.Get("room_id").(int32)].UserList
	var us []*prt.UserInfo

	for _, u_x := range us_x {
		var u prt.UserInfo
		var w prt.Weapon
		w = u_x.Weapon.Weapon
		u.Weapon = &w
		us = append(us, &u)
	}

	return us
}

func (self *readyService) AddWeaponAddr(es []*prt.WeaponExtraAttr) {
	var es_x []*prtx.WeaponExtraAttr
	for _, e := range es {
		var e_x prtx.WeaponExtraAttr
		e_x.WeaponExtraAttr = *e
		es_x = append(es_x, &e_x)
	}
	user := self.Ctx.Get("user").(prtx.UserInfo)
	user.Weapon.WeaponExtraAttrList = es_x
}
