package pack

import (
	"micro_tiktok/kitex_gen/comment"
	"micro_tiktok/kitex_gen/user"
)

func User(u *user.User) *comment.User {
	if u == nil {
		return nil
	}
	return &comment.User{
		Id:            u.Id,
		Name:          u.Name,
		FollowCount:   u.FollowCount,
		FollowerCount: u.FollowerCount,
		IsFollow:      u.IsFollow,
	}
}

func Users(users []*user.User) []*comment.User {
	us := make([]*comment.User, 0)
	for _, u := range users {
		if user2 := User(u); user2 != nil {
			us = append(us, user2)
		}
	}
	return us
}
