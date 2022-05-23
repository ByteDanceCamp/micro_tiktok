// Package handlers common.go: 各 handlers 需要的公共代码
package handlers

import (
	"context"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"micro_tiktok/cmd/api/rpc"
	"micro_tiktok/kitex_gen/user"
	"micro_tiktok/pkg/constants"
	"net/http"
	"time"
)

// ========= 请求参数相关 ============

// UserParam 登录/注册时需要获取的参数信息
type UserParam struct {
	UserName string `form:"username" binding:"required,min=2,max=32,alphanumunicode"`
	PassWord string `form:"password" binding:"required,min=6,max=32,alphanumunicode"`
}

// CommonGETParam 大部分需要鉴权的 GET 请求的参数信息
type CommonGETParam struct {
	Uid   int64  `form:"user_id" binding:"required, number"`
	Token string `form:"token" binding:"required, jwt"`
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
		ID:   user.Id,
		Name: user.Name,
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
