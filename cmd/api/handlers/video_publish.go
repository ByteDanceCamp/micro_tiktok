package handlers

import (
	"context"
	"fmt"
	"micro_tiktok/cmd/api/rpc"
	"micro_tiktok/kitex_gen/video"
	"micro_tiktok/pkg/constants"
	"micro_tiktok/pkg/errno"
	"net/http"
	"path/filepath"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type videopublisgParam struct {
	Data  int64  `json:"data"`
	Token string `json:"token"`
	Title string `json:"title"`
}

func VideoPublish(c *gin.Context) {
	// token := c.PostForm("token")
	title := c.PostForm("title")

	data, err := c.FormFile("data")
	if err != nil {
		e := errno.ConvertErr(err)
		c.JSON(http.StatusOK, BaseResponse{
			Code: e.ErrCode,
			Msg:  e.ErrMsg,
		})
		return
	}

	claims := jwt.ExtractClaims(c)
	uid := int64(claims[constants.IdentityKey].(float64))

	ntime := time.Now()

	finalName := fmt.Sprintf("%d_%s_%d_%d_%d_%d", uid, title, ntime.Month(), ntime.Day(), ntime.Minute(), ntime.Second())
	finalName = finalName + ".mp4"
	saveFile := filepath.Join("../../../resource/video/", finalName) // 目录位置需要确定一下
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		e := errno.ConvertErr(err)
		c.JSON(http.StatusOK, BaseResponse{
			Code: e.ErrCode,
			Msg:  e.ErrMsg,
		})
		return
	}

	// 截取封面
	// filename := "test.mp4"
	// width := 640
	// height := 360
	// cmd := exec.Command("ffmpeg", "-i", filename, "-vframes", "1", "-s", fmt.Sprintf("%dx%d", width, height), "-f", "singlejpeg", "-")
	// var buffer bytes.Buffer
	// cmd.Stdout = &buffer
	// if cmd.Run() != nil {
	//     panic("could not generate frame")
	// }
	// // Do something with buffer, which contains a JPEG image

	err = rpc.Publish(context.Background(), &video.PublishRequest{
		Title:   title,
		PlayUrl: saveFile,
		// CoverUrl:
		Author: string(uid),
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
