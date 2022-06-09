package pack

import (
	"micro_tiktok/cmd/user/dal/db"
	"micro_tiktok/kitex_gen/user"
)

// User db to idl
func User(u *db.User) *user.User {
	if u == nil {
		return nil
	}
	return &user.User{
		Id:   int64(u.ID),
		Name: u.UserName,
	}
}

// Users pack list of user info
func Users(users []*db.User) []*user.User {
	us := make([]*user.User, 0)
	for _, u := range users {
		if user2 := User(u); user2 != nil {
			us = append(us, user2)
		}
	}
	return us
}
