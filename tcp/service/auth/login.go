package auth

import (
	"fmt"

	"github.com/Justyer/jie"

	"github.com/Justyer/JekoServer/tcp/model"
	prt "github.com/Justyer/JekoServer/tcp/model/proto"
	prtx "github.com/Justyer/JekoServer/tcp/model/protox"
	"github.com/Justyer/JekoServer/tcp/service/base"
)

type loginService struct {
	base.BaseService
}

func NewLoginService(c *jie.Context) *loginService {
	s := &loginService{}
	s.Ctx = c
	return s
}

func (self *loginService) LoginByUser(u, p string) (prt.UserInfo, error) {
	dao := prtx.NewUserDao()
	user, err := dao.QueryUser(u, p)
	user.Conn = self.Ctx.Link.Conn
	user.IconURL = fmt.Sprintf("http://%s:%s%s", model.IP, model.HTTPPort, user.IconURL)

	self.Ctx.Put("user", user)

	var w prtx.Weapon
	w.ID = 1
	w.SN = 2
	user.Weapon = &w
	return user.UserInfo, err
}

// func (self *loginService) LoginByMAC(mac string) (user.User, error) {
// 	var user user.User
// 	user.UserID = 9527
// 	user.UserName = "zxy"
// 	return user, nil
// }
