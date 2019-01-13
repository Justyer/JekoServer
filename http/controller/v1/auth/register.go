package auth

import (
	"net/http"

	"github.com/Justyer/JekoServer/http/service/auth"
	"github.com/juju/errors"

	"github.com/Justyer/JekoServer/http/model/qsp"
	"github.com/Justyer/JekoServer/plugin/log"
	"github.com/gin-gonic/gin"
)

// Register : 正常注册
func Register(c *gin.Context) {
	var req qsp.RegisterReq
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Err("[Regisger req_err]: %s", err.Error())
		return
	}
	log.Tx("[Regisger req]: %#v", req)

	svc := auth.NewRegisterService(c)

	var resp qsp.RegisterResp

	_, err := svc.QueryUser(req.UserName, req.UserPass)
	if err == nil {
		resp.Code = 1
		resp.Msg = "user exist"
		c.JSON(http.StatusOK, &resp)
		return
	}
	err = svc.Insert(req.UserName, req.UserPass)
	if err != nil {
		resp.Code = 1
		resp.Msg = errors.Annotate(err, "insert user fail").Error()
	}

	c.JSON(http.StatusOK, &resp)
}
