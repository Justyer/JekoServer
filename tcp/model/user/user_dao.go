package user

import (
	"github.com/Justyer/JekoServer/tcp/model"
)

type userDao struct {
}

func NewUserDao() *userDao {
	return &userDao{}
}

func (self *userDao) QueryUser(u, p string) (User, error) {
	var user User
	err := model.JekoDB.
		Table("jk_user").
		Where("user_name = ? and user_pass = ?", u, p).
		Find(&user).Error
	return user, err
}
