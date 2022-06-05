// Package handlers common.go: 各 handlers 需要的公共代码
package handlers

import (
	"micro_tiktok/kitex_gen/comment"
	"micro_tiktok/kitex_gen/relation"
	"micro_tiktok/kitex_gen/user"
)

// ========= 请求参数相关 ============

// UserParam 登录/注册时需要获取的参数信息
type UserParam struct {
	UserName string `form:"username" binding:"required,min=2,max=32,alphanumunicode"`
	PassWord string `form:"password" binding:"required,min=5,max=32,alphanumunicode"`
}

// CommonGETParam 大部分需要鉴权的 GET 请求的参数信息
type CommonGETParam struct {
	Uid   int64  `form:"user_id" binding:"required,number"`
	Token string `form:"token" binding:"required,jwt"`
}

// ========= 返回相关 ============

// User Gin 返回的用户信息
type User struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type Comment struct {
	ID         int64 `json:"id"`
	User       User
	Content    string
	CreateDate string
}

// BaseResponse Gin 返回非预期（错误）结果时使用
type BaseResponse struct {
	Code int64  `json:"status_code"`
	Msg  string `json:"status_msg"`
}

// ========= 其他公共部分 ============

func UserRPC2Gin(user *user.User) *User {
	return &User{
		ID:            user.Id,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      user.IsFollow,
	}
}

func UsersRPC2Gin(users []*user.User) []*User {
	us := make([]*User, 0)
	for _, v := range users {
		if u2 := UserRPC2Gin(v); u2 != nil {
			us = append(us, u2)
		}
	}
	return us
}

func CommentUserRPC2Gin(user *comment.User) *User {
	return &User{
		ID:            user.Id,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      user.IsFollow,
	}
}

func RelationUserRPC2Gin(user *relation.User) *User {
	return &User{
		ID:            user.Id,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      user.IsFollow,
	}
}

func RelationUsersRPC2Gin(users []*relation.User) []*User {
	us := make([]*User, 0)
	for _, v := range users {
		if u2 := RelationUserRPC2Gin(v); u2 != nil {
			us = append(us, u2)
		}
	}
	return us
}

func CommentRPC2Gin(c *comment.Comment) *Comment {
	return &Comment{
		ID:         c.Id,
		User:       *CommentUserRPC2Gin(c.User),
		Content:    c.Content,
		CreateDate: c.CreateDate,
	}
}

func CommentsRPC2Gin(cs []*comment.Comment) []*Comment {
	counts := make([]*Comment, 0)
	for _, v := range cs {
		counts = append(counts, CommentRPC2Gin(v))
	}
	return counts
}
