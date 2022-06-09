package service

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"micro_tiktok/cmd/user/dal/db"
	"micro_tiktok/kitex_gen/user"
	"micro_tiktok/pkg/constants"
	"micro_tiktok/pkg/errno"
)

type CheckService struct {
	ctx context.Context
}

func NewCheckService(ctx context.Context) *CheckService {
	return &CheckService{
		ctx: ctx,
	}
}

func (c *CheckService) Check(req *user.CheckUserRequest) (uid int64, err error) {
	h := sha256.New()
	if _, err = io.WriteString(h, req.Password+constants.UserSalt); err != nil {
		return 0, err
	}
	password := fmt.Sprintf("%x", h.Sum(nil))
	username := req.Username
	u, err := db.QueryUser(c.ctx, username)
	if err != nil {
		return 0, err
	}

	if u.PassWord != password {
		return 0, errno.UserErr.WithMsg("username or password is wrong")
	}
	return int64(u.ID), nil
}
