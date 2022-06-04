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
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// FavoriteList implements the FavoriteVideoServiceImpl interface.
func (s *FavoriteVideoServiceImpl) FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (resp *favorite.FavoriteListResponse, err error) {
	resp = new(favorite.FavoriteListResponse)
	resp.VideoList, err = service.NewListService(ctx).FavoriteList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// GetFavoriteCount implements the FavoriteVideoServiceImpl interface.
func (s *FavoriteVideoServiceImpl) GetFavoriteCount(ctx context.Context, req *favorite.VideoFavoriteCountRequest) (resp *favorite.VideoFavoriteCountResponse, err error) {
	resp = new(favorite.VideoFavoriteCountResponse)
	resp, err = service.NewCountService(ctx).CountInfo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}
