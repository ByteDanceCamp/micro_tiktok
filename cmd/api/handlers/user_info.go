package handlers

import (
	"context"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"micro_tiktok/cmd/api/rpc"
	user "micro_tiktok/kitex_gen/user"
	"micro_tiktok/pkg/constants"
	"micro_tiktok/pkg/errno"
	"net/http"
)

type userInfoResponse struct {
	Code int64  `json:"status_code"`
	Msg  string `json:"status_msg"`
	User *User  `json:"user"`
}

func UserInfo(c *gin.Context) {
	var userInfoVar CommonGETParam
	if err := c.ShouldBind(&userInfoVar); err != nil {
		e := errno.ConvertErr(err)
		c.JSON(http.StatusOK, BaseResponse{
			Code: e.ErrCode,
			Msg:  e.ErrMsg,
		})
	}
	claims := jwt.ExtractClaims(c)
	uid := int64(claims[constants.IdentityKey].(float64))
	users, err := rpc.MGetUsers(context.Background(), &user.MGetUserRequest{
		UserId:        uid,
		TargetUserIds: []int64{userInfoVar.Uid},
	})
	if err != nil {
		e := errno.ConvertErr(err)
		c.JSON(http.StatusOK, BaseResponse{
			Code: e.ErrCode,
			Msg:  e.ErrMsg,
		})
		return
	}
	if len(users) == 0 {
		e := errno.UserErr.WithMsg("user isn't exist")
		c.JSON(http.StatusOK, BaseResponse{
			Code: e.ErrCode,
			Msg:  e.ErrMsg,
		})
	}

	e := errno.Success
	c.JSON(http.StatusOK, userInfoResponse{
		Code: e.ErrCode,
		Msg:  e.ErrMsg,
		User: UserRPC2Gin(users[0]),
	})

}
