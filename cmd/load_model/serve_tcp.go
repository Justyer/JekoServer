package load_model

import (
	mysql "Groot/src/golibs/db/mysql/gorm"
	"fmt"
	"log"

	"github.com/Justyer/JekoServer/tcp/model"
	"github.com/Justyer/lingo/util"
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
	IP, err := util.LocalIP()
	if err != nil {
		log.Fatal(err)
	}
	model.HTTPIPort = fmt.Sprintf("%s:%d", IP, viper.GetInt("config.http_serve_port"))
	model.TCPIPort = fmt.Sprintf("%s:%d", IP, viper.GetInt("config.tcp_serve_port"))
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
