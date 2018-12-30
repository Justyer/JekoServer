package ver

import (
	"github.com/Justyer/JekoServer/http/controller/v1/user"
	"github.com/gin-gonic/gin"
)

func API_v1(r *gin.RouterGroup) {
	r.GET("/u/info", user.GetUserInfo)
}
