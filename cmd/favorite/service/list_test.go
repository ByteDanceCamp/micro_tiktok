package service

import (
	"context"
	"micro_tiktok/cmd/favorite/dal"
	"micro_tiktok/cmd/favorite/rpc"
	"micro_tiktok/kitex_gen/favorite"
	"testing"
)

func TestListService_FavoriteList(t *testing.T) {
	dal.Init()
	rpc.INitRPC()
	res, err := NewListService(context.Background()).FavoriteList(&favorite.FavoriteListRequest{
		UserId:    2,
		TargetUid: 2,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
