package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"micro_tiktok/cmd/api/middleware/auth"
	"micro_tiktok/cmd/api/rpc"
	"micro_tiktok/kitex_gen/user"
	"micro_tiktok/pkg/errno"
	"net/http"
)

type registerResponse struct {
	Code  int64  `json:"status_code"`
	Msg   string `json:"status_msg"`
	Uid   int64  `json:"user_id"`
	Token string `json:"token"`
}

func Register(c *gin.Context) {
	var registerVar UserParam
	if err := c.ShouldBind(&registerVar); err != nil {
		e := errno.ParamsErr.WithMsg(err.Error())
		c.JSON(http.StatusOK, BaseResponse{
			Code: e.ErrCode,
			Msg:  e.ErrMsg,
		})
		return
	}
	usr, err := rpc.CreateUser(context.Background(), &user.CreateUserRequest{
		Username: registerVar.UserName,
		Password: registerVar.PassWord,
	})
	if err != nil {
		e := errno.ConvertErr(err)
		c.JSON(http.StatusOK, BaseResponse{
			Code: e.ErrCode,
			Msg:  e.ErrMsg,
		})
		return
	}
	author := auth.NewMiddleware(AuthConfig)
	token, _, err := author.JWT.TokenGenerator(usr.UserId)
	if err != nil {
		e := errno.ConvertErr(err)
		c.JSON(http.StatusOK, BaseResponse{
			Code: e.ErrCode,
			Msg:  e.ErrMsg,
		})
		return
	}
	e := errno.Success
	c.JSON(http.StatusOK, registerResponse{
		Code:  e.ErrCode,
		Msg:   e.ErrMsg,
		Uid:   usr.UserId,
		Token: token,
	})
}
