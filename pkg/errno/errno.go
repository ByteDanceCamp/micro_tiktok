package errno

import (
	"errors"
	"fmt"
)

const (
	SuccessCode     = 0
	ParamsErrCode   = 10000
	ServiceErrCode  = 20000
	UserErrCode     = 30000
	VideoErrCode    = 40000
	FavoriteErrCode = 50000
	CommentErrCode  = 60000
	RelationErrCode = 70000
)

type ErrNo struct {
	ErrCode int64
	ErrMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("err_code=%d, err_msg=%s", e.ErrCode, e.ErrMsg)
}

func NewErrNo(errCode int64, errMsg string) ErrNo {
	return ErrNo{errCode, errMsg}
}

var (
	Success     = NewErrNo(SuccessCode, "Success")
	ParamsErr   = NewErrNo(ParamsErrCode, "params have some problem")
	ServiceErr  = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	UserErr     = NewErrNo(UserErrCode, "User server has some problem")
	VideoErr    = NewErrNo(VideoErrCode, "Video server has some problem")
	FavoriteErr = NewErrNo(FavoriteErrCode, "Favorite server has some problem")
	CommentErr  = NewErrNo(CommentErrCode, "Comment server has some problem")
	RelationErr = NewErrNo(RelationErrCode, "Relation server has some problem")
)

// WithMsg 为错误编码自定义错误信息
func (e ErrNo) WithMsg(errMsg string) ErrNo {
	e.ErrMsg = errMsg
	return e
}

// ConvertErr convert error to Errno
func ConvertErr(err error) ErrNo {
	Err := ErrNo{}
	if errors.As(err, &Err) {
		return Err
	}

	s := ServiceErr
	s.ErrMsg = err.Error()
	return s
}
