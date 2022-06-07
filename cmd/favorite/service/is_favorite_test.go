package service

import (
	"context"
	"micro_tiktok/cmd/favorite/dal"
	"micro_tiktok/cmd/favorite/rpc"
	"micro_tiktok/kitex_gen/favorite"
	"testing"
)

func TestIsFavoriteService_IsFavorite(t *testing.T) {
	dal.Init()
	rpc.INitRPC()
	res, err := NewIsFavoriteService(context.Background()).IsFavorite(&favorite.IsFavoriteVideoRequest{
		VideoId: 1,
		UserId:  3,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
