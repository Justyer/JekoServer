package user

import "github.com/Justyer/JekoServer/model/weapon"

type User struct {
	ID       int           `json:"id"`
	UserName string        `json:"username"`
	Nick     string        `json:"nick"`
	Weapon   weapon.Weapon `json:"weapon"`
	CurBlood int           `json:"cur_blood"`
	MaxBlood int           `json:"max_blood"`
}

type AcountInfo struct {
	User User
}
