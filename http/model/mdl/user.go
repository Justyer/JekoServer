package mdl

// User : 用户模型
type User struct {
	UserName string `gorm:"user_name" json:"user_name"`
	UserPass string `gorm:"user_pass" json:"user_pass"`
	UserHead string `gorm:"icon_url" json:"user_head"`
}
