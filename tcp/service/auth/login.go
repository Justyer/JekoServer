package auth

import (
	"github.com/Justyer/JekoServer/tcp/model/user"
)

type LoginServe struct {
}

func NewLoginService() *LoginServe {
	return &LoginServe{}
}

func (self *LoginServe) MacLogin(mac string) (user.User, error) {
	var user user.User
	user.UserID = 9527
	user.UserName = "zxy"
	return user, nil
}

func (self *LoginServe) UserLogin(u, p string) (user.User, error) {
	dao := user.NewUserDao()
	user, err := dao.QueryUser(u, p)
	return user, err
}
