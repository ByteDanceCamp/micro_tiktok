package main

import (
	"context"
	"micro_tiktok/cmd/favorite/pack"
	"micro_tiktok/cmd/favorite/service"
	"micro_tiktok/kitex_gen/favorite"
	"micro_tiktok/pkg/errno"
)

// FavoriteVideoServiceImpl implements the last service interface defined in the IDL.
type FavoriteVideoServiceImpl struct{}

// Favorite implements the FavoriteVideoServiceImpl interface.
func (s *FavoriteVideoServiceImpl) Favorite(ctx context.Context, req *favorite.FavoriteRequest) (resp *favorite.FavoriteResponse, err error) {
	resp = new(favorite.FavoriteResponse)
	err = service.NewActionService(ctx).DoAction(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// FavoriteList implements the FavoriteVideoServiceImpl interface.
func (s *FavoriteVideoServiceImpl) FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (resp *favorite.FavoriteListResponse, err error) {
	resp = new(favorite.FavoriteListResponse)
	resp.VideoList, err = service.NewListService(ctx).FavoriteList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetFavoriteCount implements the FavoriteVideoServiceImpl interface.
func (s *FavoriteVideoServiceImpl) GetFavoriteCount(ctx context.Context, req *favorite.VideoFavoriteCountRequest) (resp *favorite.VideoFavoriteCountResponse, err error) {
	resp = new(favorite.VideoFavoriteCountResponse)
	resp.Count, err = service.NewCountService(ctx).CountInfo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// IsFavoriteVideo implements the FavoriteVideoServiceImpl interface.
func (s *FavoriteVideoServiceImpl) IsFavoriteVideo(ctx context.Context, req *favorite.IsFavoriteVideoRequest) (resp *favorite.IsFavoriteVideoResponse, err error) {
	resp = new(favorite.IsFavoriteVideoResponse)
	resp.IsFavorite, err = service.NewIsFavoriteService(ctx).IsFavorite(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
