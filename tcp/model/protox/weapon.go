package prtx

import (
	prt "github.com/Justyer/JekoServer/tcp/model/proto"
)

type Weapon struct {
	prt.Weapon

	WeaponExtraAttrList []*WeaponExtraAttr
}

func NewWeaponDao() *Weapon {
	return &Weapon{}
}
