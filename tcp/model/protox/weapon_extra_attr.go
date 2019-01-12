package prtx

import (
	prt "github.com/Justyer/JekoServer/tcp/model/proto"
)

type WeaponExtraAttr struct {
	prt.WeaponExtraAttr
}

func NewWeaponExtraAttrDao() *WeaponExtraAttr {
	return &WeaponExtraAttr{}
}
