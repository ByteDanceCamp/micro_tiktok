package service

import (
	"context"
	"micro_tiktok/cmd/favorite/dal/redis"
	"micro_tiktok/kitex_gen/favorite"
	"micro_tiktok/pkg/constants"
	"strconv"
)

type CountService struct {
	ctx context.Context
}

func NewCountService(ctx context.Context) *CountService {
	return &CountService{ctx: ctx}
}

func (c *CountService) CountInfo(req *favorite.VideoFavoriteCountRequest) (count int64, err error) {
	targetVid := strconv.Itoa(int(req.VideoId))
	likeCount, err := redis.GetLikeCount(c.ctx, constants.FavoriteVideoPre, targetVid)
	if err != nil {
		return 0, err
	}
	return likeCount, nil
}
