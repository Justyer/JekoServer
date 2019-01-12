package load_model

import (
	mysql "Groot/src/golibs/db/mysql/gorm"

	"github.com/Justyer/lingo/ip"

	"github.com/Justyer/JekoServer/tcp/model"
	"github.com/spf13/viper"
)

type TCPConfig struct {
}

func NewTCPConfig() *TCPConfig {
	return &TCPConfig{}
}

func (self *TCPConfig) LoadModel() {
	self.initIP()
	self.initDB()
}

func (self *TCPConfig) initIP() {
	model.IP = ip.MustInnerIP()
	model.HTTPPort = viper.GetString("config.http_port")
	model.TCPPort = viper.GetString("config.tcp_port")
}

func (self *TCPConfig) initDB() {
	model.JekoDB = mysql.Conn(mysql.DBInfo{
		User:   viper.GetString("mysql.user"),
		Pass:   viper.GetString("mysql.pass"),
		Host:   viper.GetString("mysql.host"),
		Port:   viper.GetInt("mysql.port"),
		DbName: viper.GetString("mysql.db"),
	})
	model.JekoDB.LogMode(true)
}
