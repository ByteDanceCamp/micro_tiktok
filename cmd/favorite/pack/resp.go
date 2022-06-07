package pack

import (
	"errors"
	"micro_tiktok/kitex_gen/favorite"
	"micro_tiktok/pkg/errno"
)

// BuildBaseResp build base resp from error
func BuildBaseResp(err error) *favorite.BaseResponse {
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
func baseResp(err errno.ErrNo) *favorite.BaseResponse {
	return &favorite.BaseResponse{
		StatusCode: err.ErrCode,
		StatusMsg:  err.ErrMsg,
	}
}
