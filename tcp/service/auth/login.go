package auth

import (
	"fmt"

	"github.com/Justyer/jie"

	"github.com/Justyer/JekoServer/model"
	prt "github.com/Justyer/JekoServer/tcp/model/proto"
	prtx "github.com/Justyer/JekoServer/tcp/model/protox"
	"github.com/Justyer/JekoServer/tcp/service/base"
)

// LoginService : 登录服务
type LoginService struct {
	base.BaseService
}

// NewLoginService : 实例化
func NewLoginService(c *jie.Context) *LoginService {
	s := &LoginService{}
	s.Ctx = c
	return s
}

// LoginByUser : 账号密码登录
func (svc *LoginService) LoginByUser(u, p string) (prt.UserInfo, error) {
	dao := prtx.NewUserDao()
	user, err := dao.QueryUser(u, p)
	user.Conn = svc.Ctx.Link.Conn
	user.IconURL = fmt.Sprintf("http://%s:%s%s", model.IP, model.HTTPPort, user.IconURL)

	var w prtx.Weapon
	w.ID = 1
	w.SN = 2
	user.Weapon = &w

	svc.Ctx.Put("user", user)

	return user.UserInfo, err
}

// func (self *loginService) LoginByMAC(mac string) (user.User, error) {
// 	var user user.User
// 	user.UserID = 9527
// 	user.UserName = "zxy"
// 	return user, nil
// }
