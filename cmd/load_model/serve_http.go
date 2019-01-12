package load_model

import (
	mysql "Groot/src/golibs/db/mysql/gorm"
	"fmt"
	"log"

	"github.com/Justyer/JekoServer/http/model"
	"github.com/Justyer/lingo/ip"
	"github.com/spf13/viper"
)

type HTTPConfig struct {
}

func NewHTTPConfig() *HTTPConfig {
	return &HTTPConfig{}
}

func (self *HTTPConfig) LoadModel() {
	self.initIP()
	self.initDB()
}

func (self *HTTPConfig) initIP() {
	IP, err := ip.InnerIP()
	if err != nil {
		log.Fatal(err)
	}
	model.HTTPIPort = fmt.Sprintf("%s:%d", IP, viper.GetInt("config.http_serve_port"))
}

func (self *HTTPConfig) initDB() {
	model.JekoDB = mysql.Conn(mysql.DBInfo{
		User:   viper.GetString("mysql.user"),
		Pass:   viper.GetString("mysql.pass"),
		Host:   viper.GetString("mysql.host"),
		Port:   viper.GetInt("mysql.port"),
		DbName: viper.GetString("mysql.db"),
	})
	model.JekoDB.LogMode(true)
}
