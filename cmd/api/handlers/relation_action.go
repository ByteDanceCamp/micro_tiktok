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

type relationActionParams struct {
	ToUserId   int64 `form:"to_user_id" binding:"required,number"`
	ActionType int32 `form:"action_type" binding:"required,number"`
}

func RelationAction(c *gin.Context) {
	var actionVar relationActionParams
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
	err := rpc.RelationAction(context.Background(), &relation.ActionRequest{
		UserId:     uid,
		ToUserId:   actionVar.ToUserId,
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
