package weapon

type Weapon struct {
	ID         int32
	SN         int32
	ExtraAttrs []WeaponExtraAttr
}

type WeaponExtraAttr struct {
	AttrType int32
	Value    int32
}
