package prtx

import (
	"net"

	"github.com/Justyer/JekoServer/tcp/model"
	prt "github.com/Justyer/JekoServer/tcp/model/proto"
)

type UserInfo struct {
	prt.UserInfo

	Weapon *Weapon
	Conn   net.Conn
}

func NewUserDao() *UserInfo {
	return &UserInfo{}
}

func (self *UserInfo) QueryUser(u, p string) (UserInfo, error) {
	var user UserInfo
	err := model.JekoDB.
		Table("jk_user").
		Where("user_name = ? and user_pass = ?", u, p).
		Find(&user).Error
	return user, err
}
