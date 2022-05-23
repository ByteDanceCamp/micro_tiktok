package service

import (
	"context"
	"micro_tiktok/cmd/user/dal/db"
	"micro_tiktok/cmd/user/pack"
	"micro_tiktok/kitex_gen/user"
	"micro_tiktok/pkg/errno"
)

type MGetService struct {
	ctx context.Context
}

func NewMGet(ctx context.Context) *MGetService {
	return &MGetService{
		ctx: ctx,
	}
}

func (m *MGetService) MGet(req *user.MGetUserRequest) ([]*user.User, error) {
	if len(req.UserIds) == 0 {
		return nil, errno.UserErr.WithMsg("params is invalid.")
	}
	users, err := db.MGet(m.ctx, req.UserIds)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errno.UserErr.WithMsg("user don't exist")
	}
	return pack.Users(users), nil
}
