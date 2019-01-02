package load_model

import (
	"github.com/Justyer/JekoServer/model"
)

func InitGlobalConfig() {
	switch model.TaskType {
	case "serve":
		switch model.Flag_T_Type {
		case "tcp":
			cfg := NewTCPConfig()
			cfg.LoadModel()
		case "http":
			cfg := NewHTTPConfig()
			cfg.LoadModel()
		case "all":
			cfg_tcp := NewTCPConfig()
			cfg_tcp.LoadModel()
			cfg_http := NewHTTPConfig()
			cfg_http.LoadModel()
		}
	}
}
