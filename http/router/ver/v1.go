package ver

import (
	"github.com/Justyer/JekoServer/http/controller/v1/auth"
	"github.com/gin-gonic/gin"
)

// APIV1 : v1版本接口
func APIV1(r *gin.RouterGroup) {
	r.GET("/u/reg", auth.Register)
}
