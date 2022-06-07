package service

import (
	"context"
	"micro_tiktok/cmd/relation/dal"
	"micro_tiktok/cmd/relation/rpc"
	"micro_tiktok/kitex_gen/relation"
	"testing"
)

func TestListService_RelationList(t *testing.T) {
	dal.Init()
	rpc.InitRPC()
	res, err := NewListService(context.Background()).RelationList(&relation.ListRequest{
		UserId:       2,
		TargetUserId: 2,
		ActionType:   1,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
