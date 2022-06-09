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
	resp = new(video.FeedResponse)
	var nextTime int64
	resp.VideoList, nextTime, err = service.NewFeedService(ctx).Feed(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.VideoErr)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.NextTime = nextTime
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
		resp.BaseResp = pack.BuildBaseResp(errno.VideoErr)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetFavoriteVideoList implements the VideoServiceImpl interface.
//func (s *VideoServiceImpl) GetFavoriteVideoList(ctx context.Context, req *video.FavoriteListRequest) (resp *video.FavoriteListResponse, err error) {
//	// TODO: Your code here...
//	resp = new(video.FavoriteListResponse)
//	resp.VideoList, err = service.NewListService(ctx).GetFavoriteVideoList(req)
//	if err != nil {
//		resp.BaseResp = pack.BuildBaseResp(errno.Success)
//		return resp, nil
//	}
//	resp.BaseResp = pack.BuildBaseResp(errno.Success)
//	return resp, nil
//}

// QueryByVid implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) QueryByVid(ctx context.Context, req *video.QueryByVidRequest) (resp *video.QueryByVidResponse, err error) {
	resp = new(video.QueryByVidResponse)
	_, err = service.NewQueryByIdService(ctx).QueryById(req)
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	if err != nil {

		resp.IsExist = false
		return resp, nil
	}
	resp.IsExist = true
	return resp, nil
}

// MGet implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) MGet(ctx context.Context, req *video.MGetRequest) (resp *video.MGetResponse, err error) {
	resp = new(video.MGetResponse)
	videos, err := service.NewMGetVideosService(ctx).MGet(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Videos = videos
	return resp, nil
}
