package user

import (
	"log"
	"net/http"

	"github.com/Justyer/JekoServer/http/model"
	"github.com/Justyer/JekoServer/http/model/user"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	var err error
	var req user.ReqUserInfo

	if err = c.ShouldBindQuery(&req); err != nil {
		log.Println(err)
	}

	aci := model.GlobalAcountInfo[req.UserName]
	var resp user.RespUserInfo
	resp.Code = 0
	resp.Msg = "succ"
	resp.User = aci.User

	c.JSON(http.StatusOK, resp)
}
