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

	// Id       string `json:"id" gorm:"unique; not null; type: varchar(80)"`
	// Author   string `json:"user_id" gorm:"not null; type: varchar(80)"`
	// PlayUrl  string `json:"play_url" gorm:"not null"`
	// CoverUrl string `json:"cover_url" gorm:"not null"`
	// Title    string `json:"title" gorm:"not null"`

	err := db.Create(l.ctx, &db.Video{
		Author:   req.Author,
		PlayUrl:  req.PlayUrl,
		CoverUrl: req.CoverUrl,
		Title:    req.Title,
	})
	if err != nil {
		return err
	}
	return nil
}
