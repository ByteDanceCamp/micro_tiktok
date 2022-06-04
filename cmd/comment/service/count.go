package service

import (
	"context"
	"micro_tiktok/cmd/comment/dal/db"
	"micro_tiktok/kitex_gen/comment"
)

type CountService struct {
	ctx context.Context
}

func NewCountService(ctx context.Context) *CountService {
	return &CountService{ctx: ctx}
}

func (c *CountService) GetCount(req *comment.CommentCountRequest) (res int64, err error) {
	count, err := db.GetCount(c.ctx, &db.CommentCount{
		VideoId: req.VideoId,
	})
	if err != nil {
		return 0, err
	}
	return count, nil
}
