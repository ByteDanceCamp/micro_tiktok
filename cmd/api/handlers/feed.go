package handlers

import (
	"context"
	jwt "github.com/appleboy/gin-jwt/v2"
	"micro_tiktok/cmd/api/rpc"
	"micro_tiktok/kitex_gen/video"
	"micro_tiktok/pkg/constants"
	"micro_tiktok/pkg/errno"
	"net/http"

	"github.com/gin-gonic/gin"
)

type feedResponse struct {
	Code     int64    `json:"status_code"`
	Msg      string   `json:"status_msg"`
	NextTime int64    `json:"next_time"`
	Videos   []*Video `json:"video_list"`
}

type FeedParam struct {
	latestTime int64 `form:"latest_time" binding:"number"`
}

func Feed(c *gin.Context) {
	var param FeedParam
	if err := c.ShouldBind(&param); err != nil {
		e := errno.ConvertErr(err)
		c.JSON(http.StatusOK, BaseResponse{
			Code: e.ErrCode,
			Msg:  e.ErrMsg,
		})
	}
	var uid int64 = 0
	if c.Query("token") != "" {
		claims := jwt.ExtractClaims(c)
		uid = int64(claims[constants.IdentityKey].(float64))
	}

	videos, nextTime, err := rpc.Feed(context.Background(), &video.FeedRequest{
		LatestTime: param.latestTime,
		Uid:        uid,
	})
	if err != nil {
		e := errno.ConvertErr(err)
		c.JSON(http.StatusOK, BaseResponse{
			Code: e.ErrCode,
			Msg:  e.ErrMsg,
		})
		return
	}
	vs := VideosRPC2Gin(videos)
	e := errno.Success
	c.JSON(http.StatusOK, feedResponse{
		Code:     e.ErrCode,
		Msg:      e.ErrMsg,
		NextTime: nextTime,
		Videos:   vs,
	})
}
