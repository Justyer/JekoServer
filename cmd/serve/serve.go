package serve

import (
	"fmt"

	"github.com/Justyer/JekoServer/http"
	"github.com/Justyer/JekoServer/model"
	"github.com/Justyer/JekoServer/tcp"
)

func Run() {

	switch model.Flag_T_Type {
	case "tcp":
		fmt.Println("-------tcp")
		tcp.TCPServe()
	case "http":
		fmt.Println("-------http")
		http.HTTPServe()
	default:
		fmt.Println("-------tcp/http")
		go http.HTTPServe()
		tcp.TCPServe()
	}
}
