package user

type User struct {
	UserID   int32  `gorm:"user_id"`
	UserName string `gorm:"user_name"`
}
