package service

import (
	"context"
	"micro_tiktok/cmd/favorite/dal/redis"
	"micro_tiktok/pkg/constants"
	"strconv"

	"micro_tiktok/kitex_gen/favorite"
)

type IsFavoriteService struct {
	ctx context.Context
}

func NewIsFavoriteService(ctx context.Context) *IsFavoriteService {
	return &IsFavoriteService{ctx: ctx}
}

func (i *IsFavoriteService) IsFavorite(req *favorite.IsFavoriteVideoRequest) (bool, error) {
	uid, vid := strconv.Itoa(int(req.UserId)), strconv.Itoa(int(req.VideoId))
	likeKeyPre := constants.FavoriteLikePre
	like, err := redis.IsLike(i.ctx, likeKeyPre, uid, vid)
	if err != nil {
		return false, err
	}
	return like, nil
}
