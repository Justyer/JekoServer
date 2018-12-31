package user

import (
	mysql "Groot/src/golibs/db/mysql/gorm"
	"testing"

	"github.com/Justyer/JekoServer/tcp/model"
)

func TestUserDao(t *testing.T) {
	model.JekoDB = mysql.Conn(mysql.DBInfo{
		User:   "root",
		Pass:   "495495",
		Host:   "localhost",
		Port:   3306,
		DbName: "db_jeko",
	})
	model.JekoDB.LogMode(true)
	dao := NewUserDao()
	dao.QueryUser("dxc", "1121")
}
