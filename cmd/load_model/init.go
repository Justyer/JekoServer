package load_model

import (
	"github.com/Justyer/JekoServer/model"
	"github.com/Justyer/lingo/db"
	"github.com/Justyer/lingo/ip"
	"github.com/spf13/viper"
)

func InitGlobalConfig() {
	initIP()
	initDB()
}

func initIP() {
	model.IP = ip.MustInnerIP()
	model.HTTPPort = viper.GetString("config.http_port")
	model.TCPPort = viper.GetString("config.tcp_port")
}

func initDB() {
	model.JekoDB = db.Conn(db.DBInfo{
		User:   viper.GetString("mysql.user"),
		Pass:   viper.GetString("mysql.pass"),
		Host:   viper.GetString("mysql.host"),
		Port:   viper.GetInt("mysql.port"),
		DbName: viper.GetString("mysql.db"),
	})
	model.JekoDB.LogMode(true)
}
