package serve

import (
	"github.com/Justyer/JekoServer/http"
	"github.com/Justyer/JekoServer/model"
	"github.com/Justyer/JekoServer/tcp"
)

func Run() {

	switch model.Flag_T_Type {
	case "tcp":
		tcp.TCPServe()
	case "http":
		http.HTTPServe()
	default:
		go http.HTTPServe()
		tcp.TCPServe()
	}
}
