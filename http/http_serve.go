package http

import (
	"fmt"
	"log"
	"syscall"
	"time"

	"github.com/Justyer/JekoServer/http/router"
	"github.com/Justyer/JekoServer/model"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func HTTPServe() {
	engine := gin.Default()

	// 加载所有版本的路由
	router.InitRouter(engine)

	endless.DefaultReadTimeOut = 5 * time.Second
	endless.DefaultWriteTimeOut = 5 * time.Second
	endless.DefaultMaxHeaderBytes = 1 << 20

	server := endless.NewServer(fmt.Sprintf("%s:%s", model.IP, model.HTTPPort), engine)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}

}
