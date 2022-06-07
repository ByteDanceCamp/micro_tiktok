package service

import (
	"context"
	"micro_tiktok/cmd/favorite/dal/redis"
	"micro_tiktok/cmd/favorite/pack"
	"micro_tiktok/cmd/favorite/rpc"
	"micro_tiktok/kitex_gen/favorite"
	"micro_tiktok/kitex_gen/video"
	"micro_tiktok/pkg/constants"
	"strconv"
)

type ListService struct {
	ctx context.Context
}

func NewListService(ctx context.Context) *ListService {
	return &ListService{ctx: ctx}
}

func (l *ListService) FavoriteList(req *favorite.FavoriteListRequest) (resp []*favorite.Video, err error) {
	targetUid := strconv.Itoa(int(req.TargetUid))
	targetKey := constants.FavoriteLikePre + targetUid

	likestStrSlice, err := redis.GetLikeList(l.ctx, targetKey, 0, -1)
	if err != nil {
		return nil, err
	}

	listIntSlice, err := pack.StrSlice2IntSlice(likestStrSlice)
	if err != nil {
		return nil, err
	}
	videos, err := rpc.MGetVideos(l.ctx, &video.MGetRequest{
		Vids:   listIntSlice,
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}

	return pack.Videos(videos), nil
}
