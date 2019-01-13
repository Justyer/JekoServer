package ready

import (
	prt "github.com/Justyer/JekoServer/tcp/model/proto"
	prtx "github.com/Justyer/JekoServer/tcp/model/protox"
	"github.com/Justyer/JekoServer/tcp/model/world"
	"github.com/Justyer/JekoServer/tcp/service/base"
	"github.com/Justyer/jie"
)

// ReadyService : 准备阶段服务
type ReadyService struct {
	base.BaseService
}

// NewReadyService : 实例化
func NewReadyService(c *jie.Context) *ReadyService {
	s := &ReadyService{}
	s.Ctx = c
	return s
}

// GetAllUserInRoom : 获取房间内所有用户的信息
func (svc *ReadyService) GetAllUserInRoom() []*prt.UserInfo {
	usX := world.JekoWorld.RoomList[svc.Ctx.Get("room_id").(int32)].UserList
	var us []*prt.UserInfo

	for _, uX := range usX {
		var u prt.UserInfo
		var w prt.Weapon
		w = uX.Weapon.Weapon
		u = uX.UserInfo
		u.Weapon = &w
		us = append(us, &u)
	}

	return us
}

// AddWeaponAddr : 增加武器属性
func (svc *ReadyService) AddWeaponAddr(es []*prt.WeaponExtraAttr) {
	var esX []*prtx.WeaponExtraAttr
	for _, e := range es {
		var eX prtx.WeaponExtraAttr
		eX.WeaponExtraAttr = *e
		esX = append(esX, &eX)
	}
	user := svc.Ctx.Get("user").(prtx.UserInfo)
	user.Weapon.WeaponExtraAttrList = esX
}
