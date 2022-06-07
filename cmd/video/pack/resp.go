package pack

import (
	"errors"
	"micro_tiktok/kitex_gen/video"
	"micro_tiktok/pkg/errno"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *video.BaseResponse {
	if err == nil {
		return baseResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}

	s := errno.ServiceErr.WithMsg(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *video.BaseResponse {
	return &video.BaseResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}
