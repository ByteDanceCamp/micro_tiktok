package handlers

import (
	"context"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"micro_tiktok/cmd/api/rpc"
	"micro_tiktok/kitex_gen/comment"
	"micro_tiktok/pkg/constants"
	"micro_tiktok/pkg/errno"
	"net/http"
)

type CommentActionParams struct {
	VideId      int64  `form:"video_id"`
	ActionType  int32  `form:"action_type"`
	CommentText string `form:"comment_text"`
	CommentId   int64  `form:"comment_id"`
}

type CommentActionResp struct {
	StatusCode int64    `json:"status_code"`
	StatusMsg  string   `json:"status_msg"`
	Comment    *Comment `json:"comment"`
}

func CommentAction(c *gin.Context) {
	var actionParams CommentActionParams
	if err := c.ShouldBind(&actionParams); err != nil {
		e := errno.ParamsErr.WithMsg(err.Error())
		c.JSON(http.StatusOK, BaseResponse{
			Code: e.ErrCode,
			Msg:  e.ErrMsg,
		})
		return
	}
	claims := jwt.ExtractClaims(c)
	uid := int64(claims[constants.IdentityKey].(float64))
	res, err := rpc.CommentAction(context.Background(), &comment.CommentRequest{
		UserId:      uid,
		VideoId:     actionParams.VideId,
		ActionType:  actionParams.ActionType,
		CommentText: actionParams.CommentText,
		CommentId:   actionParams.CommentId,
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
	c.JSON(http.StatusOK, CommentActionResp{
		StatusCode: e.ErrCode,
		StatusMsg:  e.ErrMsg,
		Comment:    CommentRPC2Gin(res),
	})
}
