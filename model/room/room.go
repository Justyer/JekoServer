package room

import "github.com/Justyer/JekoServer/model/user"

type Room struct {
	ID    int
	Users []user.User
}
