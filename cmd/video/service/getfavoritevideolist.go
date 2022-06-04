package service

import (
	"context"
	"micro_tiktok/cmd/video/dal/db"
	"micro_tiktok/cmd/video/pack"
	"micro_tiktok/kitex_gen/video"
)

type GetFavoriteVideoListService struct {
	ctx context.Context
}

func NewGetFavoriteVideoListService(ctx context.Context) *GetFavoriteVideoListService {
	return &GetFavoriteVideoListService{
		ctx: ctx,
	}
}

func (l *ListService) GetFavoriteVideoList(req *video.FavoriteListRequest) ([]*video.Video, error) {
	videos, err := db.QuertyVideos(l.ctx, req.VidList)
	if err != nil {
		return nil, err
	}
	return pack.Videos(videos), nil
}
