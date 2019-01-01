package auth

import (
	"github.com/Justyer/JekoServer/tcp/model/user"
)

type loginServe struct {
}

func NewLoginService() *loginServe {
	return &loginServe{}
}

func (self *loginServe) MacLogin(mac string) (user.User, error) {
	var user user.User
	user.UserID = 9527
	user.UserName = "zxy"
	return user, nil
}

func (self *loginServe) UserLogin(u, p string) (user.User, error) {
	dao := user.NewUserDao()
	user, err := dao.QueryUser(u, p)
	return user, err
}
