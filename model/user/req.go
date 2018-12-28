package user

type ReqUserInfo struct {
	ID       int    `form:"id" binding:"-"`
	UserName string `form:"username"  binding:"-"`
}
