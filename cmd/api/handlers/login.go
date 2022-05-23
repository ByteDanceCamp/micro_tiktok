package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"micro_tiktok/cmd/api/middleware/auth"
	"micro_tiktok/cmd/api/rpc"
	"micro_tiktok/kitex_gen/user"
	"micro_tiktok/pkg/errno"
	"net/http"
	"strconv"
)

type loginResponse struct {
	Code  int64  `json:"status_code"`
	Msg   string `json:"status_msg"`
	Uid   int64  `json:"user_id"`
	Token string `json:"token"`
}

func Login(c *gin.Context) {
	var loginVar UserParam
	if err := c.ShouldBind(&loginVar); err != nil {
		e := errno.ParamsErr.WithMsg(err.Error())
		c.JSON(http.StatusOK, BaseResponse{
			Code: e.ErrCode,
			Msg:  e.ErrMsg,
		})
		return
	}

	uid, err := rpc.CheckUser(context.Background(), &user.CheckUserRequest{
		Username: loginVar.UserName,
		Password: loginVar.PassWord,
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
	token, _, err := author.JWT.TokenGenerator(strconv.Itoa(int(uid)))
	if err != nil {
		e := errno.ConvertErr(err)
		c.JSON(http.StatusOK, BaseResponse{
			Code: e.ErrCode,
			Msg:  e.ErrMsg,
		})
		return
	}
	e := errno.Success
	c.JSON(http.StatusOK, loginResponse{
		Code:  e.ErrCode,
		Msg:   e.ErrMsg,
		Uid:   uid,
		Token: token,
	})
}
