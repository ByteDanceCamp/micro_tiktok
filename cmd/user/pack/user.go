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
func Users(us []*db.User) []*user.User {
	users := make([]*user.User, 0)
	for _, u := range us {
		if user2 := User(u); user2 != nil {
			users = append(users, user2)
		}
	}
	return users
}
