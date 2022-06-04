package service

import (
	"context"
	"micro_tiktok/cmd/user/dal"
	"micro_tiktok/cmd/user/rpc"
	"micro_tiktok/kitex_gen/user"
	"testing"
)

func TestMGetService_MGet(t *testing.T) {
	dal.Init()
	rpc.InitRPC()
	res, err := NewMGet(context.Background()).MGet(&user.MGetUserRequest{
		TargetUserIds: []int64{2},
		UserId:        2,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
