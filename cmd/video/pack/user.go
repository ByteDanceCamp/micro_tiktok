package pack

import (
	"micro_tiktok/kitex_gen/user"
	"micro_tiktok/kitex_gen/video"
)

// User rpc user.User to video.Author
func User(u *user.User) *video.User {
	return &video.User{
		Id:            u.Id,
		Name:          u.Name,
		FollowCount:   u.FollowCount,
		FollowerCount: u.FollowerCount,
		IsFollow:      u.IsFollow,
	}
}
