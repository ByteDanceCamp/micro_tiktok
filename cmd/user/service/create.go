package service

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"io"
	"micro_tiktok/cmd/user/dal/db"
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

func (c *CreateService) Create(req *user.CreateUserRequest) error {
	_, err := db.QueryUser(c.ctx, req.Username)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errno.UserErr.WithMsg("user already exist")
	}
	if err != nil {
		return errno.ConvertErr(err)
	}

	h := sha256.New()
	if _, err = io.WriteString(h, req.Password+constants.UserSalt); err != nil {
		return err
	}
	password := fmt.Sprintf("%x", h.Sum(nil))

	return db.Create(c.ctx, &db.User{
		UserName: req.Username,
		PassWord: password,
	})
}
