package service

import (
	"context"
	"micro_tiktok/cmd/user/dal/db"
	"micro_tiktok/cmd/user/pack"
	"micro_tiktok/kitex_gen/user"
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
	users, err := db.MGet(m.ctx, req.UserIds)
	if err != nil {
		return nil, err
	}
	return pack.Users(users), nil
}
