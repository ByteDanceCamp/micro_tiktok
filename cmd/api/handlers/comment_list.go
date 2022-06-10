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

type CommentListParams struct {
	VideoId int64 `form:"video_id" binding:"required,number"`
}

type CommentListResp struct {
	StatusCode int64      `json:"status_code"`
	StatusMsg  string     `json:"status_msg"`
	Comment    []*Comment `json:"comment_list"`
}

func CommentList(c *gin.Context) {
	var params CommentListParams
	if err := c.ShouldBind(&params); err != nil {
		e := errno.ParamsErr.WithMsg(err.Error())
		c.JSON(http.StatusOK, BaseResponse{
			Code: e.ErrCode,
			Msg:  e.ErrMsg,
		})
		return
	}
	claims := jwt.ExtractClaims(c)
	uid := int64(claims[constants.IdentityKey].(float64))
	res, err := rpc.CommentList(context.Background(), &comment.CommentListRequest{
		VideoId: params.VideoId,
		UserId:  uid,
	})
	if err != nil {
		e := errno.CommentErr.WithMsg(err.Error())
		c.JSON(http.StatusOK, BaseResponse{
			Code: e.ErrCode,
			Msg:  e.ErrMsg,
		})
		return
	}
	e := errno.Success
	c.JSON(http.StatusOK, &CommentListResp{
		StatusCode: e.ErrCode,
		StatusMsg:  e.ErrMsg,
		Comment:    CommentsRPC2Gin(res),
	})
}
