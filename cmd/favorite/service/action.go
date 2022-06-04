package service

import (
	"context"
	"errors"
	"micro_tiktok/cmd/favorite/dal/redis"
	"micro_tiktok/kitex_gen/favorite"
	"strconv"
)

type ActionService struct {
	ctx context.Context
}

func NewActionService(ctx context.Context) *ActionService {
	return &ActionService{ctx: ctx}
}

func (a *ActionService) DoAction(req *favorite.FavoriteRequest) error {
	var err error

	switch req.ActionType {
	case 1:
		err = redis.Like(a.ctx, &redis.ActionParams{
			Uid: strconv.Itoa(int(req.UserId)),
			Vid: strconv.Itoa(int(req.VideoId)),
		})
	case 2:
		err = redis.UnLike(a.ctx, &redis.ActionParams{
			Uid: strconv.Itoa(int(req.UserId)),
			Vid: strconv.Itoa(int(req.VideoId)),
		})
	default:
		err = errors.New("invalid action params")
	}
	return err
}
