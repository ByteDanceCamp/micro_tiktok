package handlers

import (
	"context"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"micro_tiktok/cmd/api/rpc"
	"micro_tiktok/kitex_gen/relation"
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
	user, err := rpc.UserInfo(context.Background(), &relation.InfoRequest{
		UserId:       uid,
		TargetUserId: userInfoVar.Uid,
	})
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
		User: RelationUserRPC2Gin(user),
	})

}
