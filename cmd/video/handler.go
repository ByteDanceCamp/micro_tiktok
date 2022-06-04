package main

import (
	"context"
	"micro_tiktok/cmd/video/pack"
	"micro_tiktok/cmd/video/service"
	"micro_tiktok/kitex_gen/video"
	"micro_tiktok/pkg/errno"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// Feed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Feed(ctx context.Context, req *video.FeedRequest) (resp *video.FeedResponse, err error) {
	// TODO: Your code here...
	resp = new(video.FeedResponse)
	resp.VideoList, err = service.NewFeedService(ctx).Feed(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.Success)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// Publish implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) Publish(ctx context.Context, req *video.PublishRequest) (resp *video.PublishResponse, err error) {
	resp = new(video.PublishResponse)

	err = service.NewPublishService(ctx).Publish(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// List implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) List(ctx context.Context, req *video.ListRequest) (resp *video.ListResponse, err error) {
	resp = new(video.ListResponse)
	resp.VideoList, err = service.NewListService(ctx).List(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.Success)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetFavoriteVideoList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetFavoriteVideoList(ctx context.Context, req *video.FavoriteListRequest) (resp *video.FavoriteListResponse, err error) {
	// TODO: Your code here...
	resp = new(video.FavoriteListResponse)
	resp.VideoList, err = service.NewListService(ctx).GetFavoriteVideoList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.Success)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
