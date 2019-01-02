package user

import (
	"net"

	"github.com/Justyer/JekoServer/tcp/model/weapon"
)

type User struct {
	UserID   int32         `gorm:"user_id"`
	UserName string        `gorm:"user_name"`
	UserPass string        `gorm:"user_pass"`
	IconURL  string        `gorm:"icon_url"`
	CurRoom  int32         `gorm:"-"`
	Conn     net.Conn      `gorm:"-"`
	Weapon   weapon.Weapon `gorm:"-"`
	AttrNum  int32         `gorm:"-"`
}
