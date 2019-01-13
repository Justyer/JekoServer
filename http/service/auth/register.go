package auth

import (
	"github.com/Justyer/JekoServer/http/model/mdl"
	"github.com/Justyer/JekoServer/http/service/base"
	"github.com/Justyer/JekoServer/model"
	"github.com/gin-gonic/gin"
	"github.com/juju/errors"
)

type registerService struct {
	base.BaseService
}

func NewRegisterService(c *gin.Context) *registerService {
	s := &registerService{}
	s.Ctx = c
	return s
}

func (svc *registerService) QueryUser(u, p string) (mdl.User, error) {
	var user mdl.User
	if err := model.JekoDB.Table("jk_user").
		Where("user_name=? and user_pass=?", u, p).
		Error; err != nil {
		return user, errors.Annotate(err, "query user")
	}

	return user, nil
}

func (svc *registerService) Insert(u, p string) error {
	var user mdl.User
	user.UserName = u
	user.UserPass = p
	if err := model.JekoDB.Table("jk_user").
		Create(&user).Error; err != nil {
		return errors.Annotate(err, "insert user")
	}
	return nil
}
