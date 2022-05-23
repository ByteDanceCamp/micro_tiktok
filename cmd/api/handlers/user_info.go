package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"micro_tiktok/cmd/api/rpc"
	"micro_tiktok/kitex_gen/user"
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

	users, err := rpc.MGetUser(context.Background(), &user.MGetUserRequest{UserIds: []int64{userInfoVar.Uid}})
	if err != nil {
		e := errno.ConvertErr(err)
		c.JSON(http.StatusOK, BaseResponse{
			Code: e.ErrCode,
			Msg:  e.ErrMsg,
		})
		return
	}

	e := errno.Success
	c.JSON(http.StatusOK, userInfoResponse{
		Code: e.ErrCode,
		Msg:  e.ErrMsg,
		User: UsersRPC2Gin(users)[0],
	})

}
