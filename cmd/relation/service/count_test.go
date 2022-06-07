package service

import (
	"context"
	"micro_tiktok/cmd/relation/dal"
	"micro_tiktok/cmd/relation/rpc"
	"micro_tiktok/kitex_gen/relation"
	"testing"
)

func TestCountService_CountInfo(t *testing.T) {
	dal.Init()
	rpc.InitRPC()
	res, err := NewCountService(context.Background()).CountInfo(&relation.InfoRequest{
		UserId:       3,
		TargetUserId: 3,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
