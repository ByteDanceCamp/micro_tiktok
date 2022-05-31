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

type followListResp struct {
	Code     int64   `json:"status_code"`
	Msg      string  `json:"status_msg"`
	UserList []*User `json:"user_list"`
}

func FollowList(c *gin.Context) {
	var followListVar CommonGETParam
	if err := c.ShouldBind(&followListVar); err != nil {
		e := errno.ParamsErr.WithMsg(err.Error())
		c.JSON(http.StatusOK, BaseResponse{
			Code: e.ErrCode,
			Msg:  e.ErrMsg,
		})
		return
	}
	claims := jwt.ExtractClaims(c)
	uid := int64(claims[constants.IdentityKey].(float64))
	followList, err := rpc.RelationList(context.Background(), &relation.ListRequest{
		UserId:       uid,
		TargetUserId: followListVar.Uid,
		ActionType:   1,
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
	c.JSON(http.StatusOK, followListResp{
		Code:     e.ErrCode,
		Msg:      e.ErrMsg,
		UserList: RelationUsersRPC2Gin(followList),
	})
}
