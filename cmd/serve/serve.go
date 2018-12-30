package serve

import (
	"github.com/Justyer/JekoServer/http"
	"github.com/Justyer/JekoServer/tcp"
	"github.com/Justyer/JekoServer/tcp/model"
)

func Run() {

	switch model.ServeType {
	case "tcp":
		tcp.TCPServe()
	case "http":
		http.HTTPServe()
	default:
		go http.HTTPServe()
		tcp.TCPServe()
	}
}
