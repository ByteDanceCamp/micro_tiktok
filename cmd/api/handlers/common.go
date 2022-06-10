// Package handlers common.go: 各 handlers 需要的公共代码
package handlers

import (
	"micro_tiktok/kitex_gen/comment"
	"micro_tiktok/kitex_gen/relation"
	"micro_tiktok/kitex_gen/user"

	"micro_tiktok/kitex_gen/favorite"
	"micro_tiktok/kitex_gen/video"
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
	ID         int64  `json:"id"`
	User       *User  `json:"user"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}

// BaseResponse Gin 返回非预期（错误）结果时使用
type BaseResponse struct {
	Code int64  `json:"status_code"`
	Msg  string `json:"status_msg"`
}

// Video Gin 返回的视频信息
type Video struct {
	Id            int64  `json:"id"`             // 视频 ID
	Author        User   `json:"author"`         // 视频作者信息
	PlayUrl       string `json:"play_url"`       // 视频播放地址
	CoverUrl      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"` // 视频点赞总数
	CommentCount  int64  `json:"comment_count"`  // 视频评论总数
	IsFavorite    bool   `json:"is_favorite"`    // 是否点赞，true：已点赞；false：未点赞
	Title         string `json:"title"`          // 视频标题
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
	if user == nil {
		return nil
	}
	return &User{
		ID:            user.Id,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      user.IsFollow,
	}
}

func FavoriteVideoRPC2Gin(fv *favorite.Video) *Video {
	return &Video{
		Id: fv.Id,
		Author: User{
			ID:            fv.Author.Id,
			Name:          fv.Author.Name,
			FollowCount:   fv.Author.FollowCount,
			FollowerCount: fv.Author.FollowerCount,
			IsFollow:      fv.Author.IsFollow,
		},
		PlayUrl:       fv.PlayUrl,
		CoverUrl:      fv.CoverUrl,
		FavoriteCount: fv.FavoriteCount,
		CommentCount:  fv.CommentCount,
		IsFavorite:    fv.IsFavorite,
		Title:         fv.Title,
	}
}

func FavoriteVideosRPC2Gin(fvs []*favorite.Video) []*Video {
	vs := make([]*Video, 0)
	for _, v := range fvs {
		if v2 := FavoriteVideoRPC2Gin(v); v2 != nil {
			vs = append(vs, v2)
		}
	}
	return vs
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
	if c == nil {
		return nil
	}
	return &Comment{
		ID:         c.Id,
		User:       CommentUserRPC2Gin(c.User),
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

func VideoRPC2Gin(vi *video.Video) *Video {
	return &Video{
		Id: vi.Id,
		Author: User{
			ID:            vi.Author.Id,
			Name:          vi.Author.Name,
			FollowCount:   vi.Author.FollowCount,
			FollowerCount: vi.Author.FollowerCount,
			IsFollow:      vi.Author.IsFollow,
		},
		PlayUrl:       vi.PlayUrl,
		CoverUrl:      vi.CoverUrl,
		FavoriteCount: vi.FavoriteCount,
		CommentCount:  vi.CommentCount,
		IsFavorite:    vi.IsFavorite,
		Title:         vi.Title,
	}
}

func VideosRPC2Gin(videos []*video.Video) []*Video {
	vs := make([]*Video, 0)
	for _, v := range videos {
		if v2 := VideoRPC2Gin(v); v2 != nil {
			vs = append(vs, v2)
		}
	}
	return vs
}
