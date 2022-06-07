package service

import (
	"context"
	"micro_tiktok/cmd/favorite/dal"
	"micro_tiktok/cmd/favorite/rpc"
	"micro_tiktok/kitex_gen/favorite"
	"testing"
)

func TestCountService_CountInfo(t *testing.T) {
	dal.Init()
	rpc.INitRPC()
	res, err := NewCountService(context.Background()).CountInfo(&favorite.VideoFavoriteCountRequest{VideoId: 2})
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
