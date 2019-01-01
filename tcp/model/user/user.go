package user

import "net"

type User struct {
	UserID   int32    `gorm:"user_id"`
	UserName string   `gorm:"user_name"`
	UserPass string   `gorm:"user_pass"`
	IconURL  string   `gorm:"icon_url"`
	Conn     net.Conn `gorm:"-"`
}
