package service

import (
	"context"
	"micro_tiktok/cmd/video/dal/db"
	"micro_tiktok/kitex_gen/video"
)

type PublishService struct {
	ctx context.Context
}

func NewPublishService(ctx context.Context) *PublishService {
	return &PublishService{
		ctx: ctx,
	}
}

func (l *PublishService) Publish(req *video.PublishRequest) error {
	if err := db.Create(l.ctx, &db.Video{
		Uid:      req.Uid,
		PlayUrl:  req.PlayUrl,
		CoverUrl: req.CoverUrl,
		Title:    req.Title,
	}); err != nil {
		return err
	}
	return nil
}
