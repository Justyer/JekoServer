package load_model

import (
	mysql "Groot/src/golibs/db/mysql/gorm"
	"fmt"
	"log"

	"github.com/Justyer/JekoServer/tcp/model"
	"github.com/Justyer/lingo/util"
	"github.com/spf13/viper"
)

func InitGlobalConfig() {
	initIP()
	initDB()
}

func initIP() {
	IP, err := util.LocalIP()
	if err != nil {
		log.Fatal(err)
	}
	model.TCPIPort = fmt.Sprintf("%s:%d", IP, viper.GetInt("config.tcp_serve_port"))
	model.HTTPIPort = fmt.Sprintf("%s:%d", IP, viper.GetInt("config.http_serve_port"))
}

func initDB() {
	model.JekoDB = mysql.Conn(mysql.DBInfo{
		User:   viper.GetString("mysql.user"),
		Pass:   viper.GetString("mysql.pass"),
		Host:   viper.GetString("mysql.host"),
		Port:   viper.GetInt("mysql.port"),
		DbName: viper.GetString("mysql.db"),
	})
	model.JekoDB.LogMode(true)
}
