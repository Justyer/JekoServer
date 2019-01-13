package qsp

// RegisterReq : 注册的请求参数
type RegisterReq struct {
	UserName string `json:"user_name"`
	UserPass string `json:"user_pass"`
}
