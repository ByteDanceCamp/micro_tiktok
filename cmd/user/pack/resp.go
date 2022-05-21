package pack

import (
	"errors"
	"micro_tiktok/kitex_gen/user"
	"micro_tiktok/pkg/errno"
)

// BuildBaseResp build baseResp from error
func BuildBaseResp(err error) *user.BaseResponse {
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

func baseResp(err errno.ErrNo) *user.BaseResponse {
	return &user.BaseResponse{StatusCode: err.ErrCode, StatusMsg: err.ErrMsg}
}
