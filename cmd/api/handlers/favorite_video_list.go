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

type videoListResp struct {
	Code      int64    `json:"status_code"`
	Msg       string   `json:"status_msg"`
	VideoList []*Video `json:"video_list"`
}

func FavoriteVideoList(c *gin.Context) {
	var videoListVar CommonGETParam
	if err := c.ShouldBindQuery(&videoListVar); err != nil {
		c.JSON(200, gin.H{
			"status_code": 400,
			"status_msg":  "参数错误",
		})
		return
	}
	claims := jwt.ExtractClaims(c)
	uid := int64(claims[constants.IdentityKey].(float64))
	videoList, err := rpc.FavoriteVideosList(context.Background(), &favorite.FavoriteListRequest{
		UserId:    uid,
		TargetUid: videoListVar.Uid,
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
	c.JSON(http.StatusOK, videoListResp{
		Code:      e.ErrCode,
		Msg:       e.ErrMsg,
		VideoList: FavoriteVideosRPC2Gin(videoList),
	})
}
