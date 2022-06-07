package service

import (
	"context"
	"micro_tiktok/cmd/video/dal/db"
	"micro_tiktok/cmd/video/pack"
	"micro_tiktok/kitex_gen/video"
)

type QueryByIdService struct {
	ctx context.Context
}

func NewQueryByIdService(ctx context.Context) *QueryByIdService {
	return &QueryByIdService{ctx: ctx}
}

func (q *QueryByIdService) QueryById(req *video.QueryByVidRequest) (res *video.Video, err error) {
	result, err := db.QueryById(q.ctx, req.VideoId)
	if err != nil {
		return nil, err
	}
	return pack.Video(result), nil
}
