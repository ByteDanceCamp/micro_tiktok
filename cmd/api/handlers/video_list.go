package handlers

import (
	"context"
	"micro_tiktok/cmd/api/rpc"
	"micro_tiktok/kitex_gen/video"
	"micro_tiktok/pkg/constants"
	"micro_tiktok/pkg/errno"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type videoListResponse struct {
	Code   int64    `json:"status_code"`
	Msg    string   `json:"status_msg"`
	Videos []*Video `json:"video_list"`
}

func VList(c *gin.Context) {
	var videoListParam CommonGETParam
	if err := c.ShouldBind(&videoListParam); err != nil {
		e := errno.ConvertErr(err)
		c.JSON(http.StatusOK, BaseResponse{
			Code: e.ErrCode,
			Msg:  e.ErrMsg,
		})
	}
	claims := jwt.ExtractClaims(c)
	uid := int64(claims[constants.IdentityKey].(float64))
	videos, err := rpc.List(context.Background(), &video.ListRequest{
		UserId:    uid,
		TargetUid: videoListParam.Uid,
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
	c.JSON(http.StatusOK, videoListResponse{
		Code:   e.ErrCode,
		Msg:    e.ErrMsg,
		Videos: VideosRPC2Gin(videos),
	})
}
