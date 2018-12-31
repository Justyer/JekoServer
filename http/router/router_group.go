package router

import (
	"net/http"

	"github.com/Justyer/JekoServer/http/router/ver"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	r.Static("/static", "http/static")
	r.LoadHTMLGlob("http/template/*")

	r.GET("/canvas", func(c *gin.Context) {
		c.HTML(http.StatusOK, "canvas.html", gin.H{})
	})
	ver.API_v1(r.Group("v1"))
}
