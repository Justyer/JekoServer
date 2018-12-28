package router

import (
	"github.com/Justyer/JekoServer/router/ver"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	ver.API_v1(r.Group("v1"))
}
