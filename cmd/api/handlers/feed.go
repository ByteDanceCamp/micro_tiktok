package handlers

import (
	"context"
	"micro_tiktok/cmd/api/rpc"
	"micro_tiktok/kitex_gen/comment"
	"micro_tiktok/kitex_gen/favorite"
	"micro_tiktok/kitex_gen/user"
	"micro_tiktok/kitex_gen/video"
	"micro_tiktok/pkg/constants"
	"micro_tiktok/pkg/errno"
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type feedResponse struct {
	Code   int64    `json:"status_code"`
	Msg    string   `json:"status_msg"`
	Videos []*Video `json:"video_list"`
}

type FeedParam struct {
	latest_time int64  `form:"latest_time" binding:"required,number"`
	Token       string `form:"token" binding:"required,jwt"`
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
	claims := jwt.ExtractClaims(c)
	uid := int64(claims[constants.IdentityKey].(float64))
	videos, err := rpc.Feed(context.Background(), &video.FeedRequest{
		LatestTime: param.latest_time,
	})

	if err != nil {
		e := errno.ConvertErr(err)
		c.JSON(http.StatusOK, BaseResponse{
			Code: e.ErrCode,
			Msg:  e.ErrMsg,
		})
		return
	}
	for _, vi := range videos {
		var tmp = []int64{}
		tmp = append(tmp, vi.Author.Id)
		rep, _ := rpc.MGetUsers(context.Background(), &user.MGetUserRequest{
			UserId:        uid,
			TargetUserIds: tmp,
		})
		u := rep[0]
		vi.Author.Id = u.Id
		vi.Author.Name = u.Name
		vi.Author.FollowCount = u.FollowCount
		vi.Author.FollowerCount = u.FollowerCount
		vi.Author.IsFollow = u.IsFollow
		vi.CommentCount, _ = rpc.CountRes(context.Background(), &comment.CommentCountRequest{VideoId: vi.Id})
		vi.FavoriteCount, _ = rpc.GetFavoriteCount(context.Background(), &favorite.VideoFavoriteCountRequest{VideoId: vi.Id})
	}

	if err != nil {
		e := errno.ConvertErr(err)
		c.JSON(http.StatusOK, BaseResponse{
			Code: e.ErrCode,
			Msg:  e.ErrMsg,
		})
		return
	}

	e := errno.Success
	c.JSON(http.StatusOK, videolistResponse{
		Code:   e.ErrCode,
		Msg:    e.ErrMsg,
		Videos: VideosRPC2Gin(videos),
	})
}
