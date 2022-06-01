package pack

import (
	"micro_tiktok/kitex_gen/relation"
	"micro_tiktok/kitex_gen/user"
)

// User user.User to relation.User
func User(u *user.User) *relation.User {
	if u == nil {
		return nil
	}
	return &relation.User{
		Id:            u.Id,
		Name:          u.Name,
		FollowCount:   u.FollowCount,
		FollowerCount: u.FollowerCount,
		IsFollow:      u.IsFollow,
	}
}

// Users pack list of user info
func Users(users []*user.User) []*relation.User {
	us := make([]*relation.User, 0)
	for _, u := range users {
		if user2 := User(u); user2 != nil {
			us = append(us, user2)
		}
	}
	return us
}
