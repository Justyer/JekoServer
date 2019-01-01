package room

import "github.com/Justyer/JekoServer/tcp/model/user"

type RoomInfo struct {
	ID    int32
	Users []*user.User
}
