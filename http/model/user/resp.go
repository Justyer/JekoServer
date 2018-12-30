package user

type RespUserInfo struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	User User   `json:"user"`
}
