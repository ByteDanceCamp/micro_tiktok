package service

import (
	"context"
	"micro_tiktok/cmd/video/dal/db"
	"micro_tiktok/cmd/video/pack"
	"micro_tiktok/kitex_gen/video"
)

type ListService struct {
	ctx context.Context
}

func NewListService(ctx context.Context) *ListService {
	return &ListService{
		ctx: ctx,
	}
}

func (l *ListService) List(req *video.ListRequest) ([]*video.Video, error) {
	videos, err := db.GetMyList(l.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return pack.Videos(videos), nil
}
