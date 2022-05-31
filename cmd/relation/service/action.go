package service

import (
	"context"
	"errors"
	"micro_tiktok/cmd/relation/dal/redis"
	"micro_tiktok/cmd/relation/rpc"
	"micro_tiktok/kitex_gen/relation"
	"micro_tiktok/kitex_gen/user"
	"strconv"
)

type ActionService struct {
	ctx context.Context
}

func NewActionService(ctx context.Context) *ActionService {
	return &ActionService{
		ctx: ctx,
	}
}

func (a *ActionService) Action(req *relation.ActionRequest) error {
	var err error
	// 验证目标用户是否存在
	_, err = rpc.MGetUser(a.ctx, &user.MGetUserRequest{UserIds: []int64{req.ToUserId}})
	if err != nil {
		return err
	}
	switch req.ActionType {
	case 1:
		err = redis.Follow(a.ctx, &redis.ActionParams{
			Uid:   strconv.Itoa(int(req.UserId)),
			ToUid: strconv.Itoa(int(req.ToUserId)),
		})
	case 2:
		err = redis.UnFollow(a.ctx, &redis.ActionParams{
			Uid:   strconv.Itoa(int(req.UserId)),
			ToUid: strconv.Itoa(int(req.ToUserId)),
		})
	default:
		err = errors.New("invalid action params")
	}
	return err
}
