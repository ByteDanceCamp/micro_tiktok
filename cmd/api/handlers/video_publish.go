package handlers

import (
	"github.com/gin-gonic/gin"
	"micro_tiktok/cmd/api/utils"
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
	//content := c.Request.FormValue("title")
	//claims := jwt.ExtractClaims(c)
	//uid := int64(claims[constants.IdentityKey].(float64))
	fileSize := fileHeader.Size
	// ToDo: 校验文件格式和评论内容

	playUrl, err := utils.UpLoadFile(file, fileSize)
	if err != nil {
		e := errno.VideoErr.WithMsg(err.Error())
		c.JSON(http.StatusOK, BaseResponse{
			Code: e.ErrCode,
			Msg:  e.ErrMsg,
		})
		return
	}
	c.JSON(http.StatusOK, BaseResponse{
		Code: errno.SuccessCode,
		Msg:  playUrl,
	})
}
