package auth

import (
	"github.com/Justyer/JekoServer/plugin/log"
	prt "github.com/Justyer/JekoServer/tcp/model/proto"
	"github.com/Justyer/JekoServer/tcp/service/auth"

	"github.com/Justyer/jie"
)

// Login : 登录
func Login(c *jie.Context) {
	var req prt.LoginReq
	if err := c.BindProtoReq(&req); err != nil {
		log.Err("[Login bind_req_err]: %s", err.Error())
		return
	}
	log.Tx("[Login req]: %s", req.String())

	svc := auth.NewLoginService(c)
	user, err := svc.LoginByUser(req.UserName, req.PassWord)

	var resp prt.LoginResp
	if err != nil {
		resp.Code = 1
	}
	resp.UserID = user.UserID
	resp.UserName = user.UserName
	dataB, err := c.PackProtoResp(&resp)
	if err != nil {
		log.Err("[Login pack_resp_err]: %s", err.Error())
		return
	}

	log.Tx("[Login resp]: %s", resp.String())

	finalB := svc.PackProtocol(uint16(prt.MsgCmd_value["Login_LoginResp"]), dataB)

	c.Send(finalB)
}
