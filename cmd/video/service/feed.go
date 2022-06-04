package service

import (
	"context"
	"micro_tiktok/cmd/video/dal/db"
	"micro_tiktok/cmd/video/pack"
	"micro_tiktok/kitex_gen/video"
)

type FeedService struct {
	ctx context.Context
}

func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{
		ctx: ctx,
	}
}

func (l *FeedService) Feed(req *video.FeedRequest) ([]*video.Video, error) {
	videos, err := db.GetNewestList(l.ctx, req.LatestTime)
	if err != nil {
		return nil, err
	}
	return pack.Videos(videos), nil
}
