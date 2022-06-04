package service

import (
	"context"
	"micro_tiktok/cmd/favorite/dal/redis"
	"micro_tiktok/cmd/favorite/pack"
	"micro_tiktok/cmd/favorite/rpc"
	"micro_tiktok/kitex_gen/favorite"
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
	userId := strconv.Itoa(int(req.UserId))
	targetKey := constants.FavoriteLikePre + userId

	likstStrSlice, err := redis.GetLikeList(l.ctx, targetKey, 0, -1)
	if err != nil {
		return
	}

	listIntSlice, err := pack.StrSlice2IntSlice(likstStrSlice)
	if err != nil {
		return
	}
	// 此处需要video的rpc接口，用videoid集得到video集合
	videos, err := rpc.GetFavoriteVideoList(l.ctx, &video.GetVideos{
		VideoIds: listIntSlice,
		UserId:   req.UserId,
	})
	if err != nil {
		return
	}

	return pack.Videos(videos), nil
}
