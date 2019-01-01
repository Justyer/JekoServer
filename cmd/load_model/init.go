package load_model

import (
	"fmt"

	"github.com/Justyer/JekoServer/model"
)

func InitGlobalConfig() {
	switch model.TaskType {
	case "serve":
		switch model.Flag_T_Type {
		case "tcp":
			fmt.Println("tcp-------")
			cfg := NewTCPConfig()
			cfg.LoadModel()
		case "http":
			fmt.Println("http-------")
			cfg := NewHTTPConfig()
			cfg.LoadModel()
		case "all":
			fmt.Println("tcp/http-------")
			cfg_tcp := NewTCPConfig()
			cfg_tcp.LoadModel()
			cfg_http := NewHTTPConfig()
			cfg_http.LoadModel()
		}
	}
}
