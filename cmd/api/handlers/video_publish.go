package handlers

import (
	"context"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"micro_tiktok/cmd/api/rpc"
	"micro_tiktok/cmd/api/utils"
	"micro_tiktok/kitex_gen/video"
	"micro_tiktok/pkg/constants"
	"micro_tiktok/pkg/errno"
	"net/http"
)

func VideoPublish(c *gin.Context) {
	// 接受数据（multipart/form-data）
	file, fileHeader, err := c.Request.FormFile("data")
	if err != nil {
		e := errno.VideoErr.WithMsg(err.Error())
		c.JSON(http.StatusOK, BaseResponse{
			Code: e.ErrCode,
			Msg:  e.ErrMsg,
		})
		return
	}
	content := c.Request.FormValue("title")
	claims := jwt.ExtractClaims(c)
	uid := int64(claims[constants.IdentityKey].(float64))

	fileSize := fileHeader.Size
	fileName, videoSuffix, err := utils.ValidateVideoInfo(fileHeader)
	if err != nil {
		e := errno.VideoErr.WithMsg(err.Error())
		c.JSON(http.StatusOK, BaseResponse{
			Code: e.ErrCode,
			Msg:  e.ErrMsg,
		})
		return
	}
	_, err = utils.UpLoadFile(file, fileName, videoSuffix, fileSize)
	if err != nil {
		e := errno.VideoErr.WithMsg(err.Error())
		c.JSON(http.StatusOK, BaseResponse{
			Code: e.ErrCode,
			Msg:  e.ErrMsg,
		})
		return
	}

	err = rpc.Publish(context.Background(), &video.PublishRequest{
		Title:    content,
		PlayUrl:  "video/" + fileName + videoSuffix,
		CoverUrl: "cover/" + fileName + ".jpg",
		Uid:      uid,
	})
	if err != nil {
		e := errno.VideoErr.WithMsg(err.Error())
		c.JSON(http.StatusOK, BaseResponse{
			Code: e.ErrCode,
			Msg:  e.ErrMsg,
		})
		return
	}
	c.JSON(http.StatusOK, BaseResponse{
		Code: errno.Success.ErrCode,
		Msg:  errno.Success.ErrMsg,
	})
}
