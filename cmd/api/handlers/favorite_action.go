package handlers

import (
	"context"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"micro_tiktok/cmd/api/rpc"
	"micro_tiktok/kitex_gen/favorite"
	"micro_tiktok/pkg/constants"
	"micro_tiktok/pkg/errno"
	"net/http"
)

type FavoriteActionParams struct {
	UserID     int64 `json:"user_id" binding:"required,number"`
	VideoID    int64 `json:"video_id" binding:"required,number"`
	ActionType int32 `json:"action_type" binding:"required,number"`
}

func FavoriteAction(c *gin.Context) {
	var actionVar FavoriteActionParams
	if err := c.ShouldBind(&actionVar); err != nil {
		e := errno.ParamsErr.WithMsg(err.Error())
		c.JSON(http.StatusOK, BaseResponse{
			Code: e.ErrCode,
			Msg:  e.ErrMsg,
		})
		return
	}
	claims := jwt.ExtractClaims(c)
	uid := int64(claims[constants.IdentityKey].(float64))
	err := rpc.FavoriteAction(context.Background(), &favorite.FavoriteRequest{
		UserId:     uid,
		VideoId:    actionVar.VideoID,
		ActionType: actionVar.ActionType,
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
	c.JSON(http.StatusOK, BaseResponse{
		Code: e.ErrCode,
		Msg:  e.ErrMsg,
	})
}
