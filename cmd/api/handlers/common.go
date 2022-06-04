// Package handlers common.go: 各 handlers 需要的公共代码
package handlers

import (
	"context"
	"micro_tiktok/cmd/api/rpc"
	"micro_tiktok/kitex_gen/favorite"
	"micro_tiktok/kitex_gen/relation"
	"micro_tiktok/kitex_gen/user"
	"micro_tiktok/kitex_gen/video"
	"micro_tiktok/pkg/constants"
	"net/http"
	"time"
	"unsafe"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
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

// BaseResponse Gin 返回非预期（错误）结果时使用
type BaseResponse struct {
	Code int64  `json:"status_code"`
	Msg  string `json:"status_msg"`
}

// Video Gin 返回的视频信息
type Video struct {
	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                            // 视频 ID
	Author        *User  `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`                                     // 视频作者信息
	PlayUrl       string `protobuf:"bytes,3,opt,name=play_url,json=playUrl,proto3" json:"play_url,omitempty"`                    // 视频播放地址
	CoverUrl      string `protobuf:"bytes,4,opt,name=cover_url,json=coverUrl,proto3" json:"cover_url,omitempty"`                 // 视频封面地址
	FavoriteCount int64  `protobuf:"varint,5,opt,name=favorite_count,json=favoriteCount,proto3" json:"favorite_count,omitempty"` // 视频点赞总数
	CommentCount  int64  `protobuf:"varint,6,opt,name=comment_count,json=commentCount,proto3" json:"comment_count,omitempty"`    // 视频评论总数
	IsFavorite    bool   `protobuf:"varint,7,opt,name=is_favorite,json=isFavorite,proto3" json:"is_favorite,omitempty"`          // 是否点赞，true：已点赞；false：未点赞
	Title         string `protobuf:"bytes,8,opt,name=title,proto3" json:"title,omitempty"`                                       // 视频标题
}

// ========= 其他公共部分 ============

// AuthConfig 生成 JWT 中间件对象时需要的配置信息
var AuthConfig = jwt.GinJWTMiddleware{
	Key:        []byte(constants.JWTSecretKey),
	Timeout:    time.Hour,
	MaxRefresh: time.Hour,
	Authenticator: func(c *gin.Context) (interface{}, error) {
		var loginVar UserParam
		if err := c.ShouldBind(&loginVar); err != nil {
			return "", jwt.ErrMissingLoginValues
		}

		if len(loginVar.UserName) == 0 || len(loginVar.PassWord) == 0 {
			return "", jwt.ErrMissingLoginValues
		}

		return rpc.CheckUser(context.Background(), &user.CheckUserRequest{Username: loginVar.UserName, Password: loginVar.PassWord})
	},
	Authorizator: nil,
	PayloadFunc: func(data interface{}) jwt.MapClaims {
		if v, ok := data.(int64); ok {
			return jwt.MapClaims{
				constants.IdentityKey: v,
			}
		}
		return jwt.MapClaims{}
	},
	Unauthorized: func(c *gin.Context, code int, message string) {
		c.JSON(http.StatusOK, gin.H{
			"status_code": code,
			"status_msg":  message,
		})
	},
	TokenLookup: "query: token",
}

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

func FavoriteVideoRPC2Gin(fv *favorite.Video) *Video {
	return &Video{
		Id:            fv.Id,
		Author:        (*User)(unsafe.Pointer(fv.Author)),
		PlayUrl:       fv.PlayUrl,
		CoverUrl:      fv.CoverUrl,
		FavoriteCount: fv.FavoriteCount,
		CommentCount:  fv.CommentCount,
		IsFavorite:    fv.IsFavorite,
		Title:         fv.Title,
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
func FavoriteVideosRPC2Gin(fvs []*favorite.Video) []*Video {
	vs := make([]*Video, 0)
	for _, v := range fvs {
		if v2 := FavoriteVideoRPC2Gin(v); v2 != nil {
			vs = append(vs, v2)
		}
	}
	return vs
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
