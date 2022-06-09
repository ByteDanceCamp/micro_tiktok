package pack

import (
	"errors"
	"micro_tiktok/kitex_gen/comment"
	"micro_tiktok/pkg/errno"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *comment.BaseResponse {
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

func baseResp(err errno.ErrNo) *comment.BaseResponse {
	return &comment.BaseResponse{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
