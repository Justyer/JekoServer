package load_model

import (
	"fmt"
	"log"

	"github.com/Justyer/JekoServer/model"
	"github.com/Justyer/lingo/util"
	"github.com/spf13/viper"
)

func InitGlobalConfig() {
	initIP()
}

func initIP() {
	IP, err := util.LocalIP()
	if err != nil {
		log.Fatal(err)
	}
	model.TCPIPort = fmt.Sprintf("%s:%d", IP, viper.GetInt("config.tcp_serve_port"))
	model.HTTPIPort = fmt.Sprintf("%s:%d", IP, viper.GetInt("config.http_serve_port"))
}
