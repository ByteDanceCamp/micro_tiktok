package service

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"io"
	"micro_tiktok/cmd/user/dal/db"
	"micro_tiktok/cmd/user/pack"
	"micro_tiktok/kitex_gen/user"
	"micro_tiktok/pkg/constants"
	"micro_tiktok/pkg/errno"
)

type CreateService struct {
	ctx context.Context
}

func NewCreateService(ctx context.Context) *CreateService {
	return &CreateService{
		ctx: ctx,
	}
}

func (c *CreateService) Create(req *user.CreateUserRequest) (*user.User, error) {
	_, err := db.QueryUser(c.ctx, req.Username)
	if err == nil {
		return nil, errno.UserErr.WithMsg("user is exist")
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	h := sha256.New()
	if _, err = io.WriteString(h, req.Password+constants.UserSalt); err != nil {
		return nil, err
	}
	password := fmt.Sprintf("%x", h.Sum(nil))

	usr, err := db.Create(c.ctx, &db.User{
		UserName: req.Username,
		PassWord: password,
	})
	if err != nil {
		return nil, err
	}
	return pack.User(usr), nil
}
