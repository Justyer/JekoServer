package serve

import (
	"github.com/Justyer/JekoServer/model"
	"github.com/Justyer/JekoServer/serve/http"
	"github.com/Justyer/JekoServer/serve/tcp"
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
