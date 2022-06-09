package service

import (
	"context"
	"micro_tiktok/cmd/user/dal/db"
	"micro_tiktok/kitex_gen/user"
)

type IsExistService struct {
	ctx context.Context
}

func NewIsExistService(ctx context.Context) *IsExistService {
	return &IsExistService{ctx: ctx}
}

func (i *IsExistService) IsExist(req *user.IsExistByIdRequest) (bool, error) {
	return db.IsExist(i.ctx, req.UserId)
}
