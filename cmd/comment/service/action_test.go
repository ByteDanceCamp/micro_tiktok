package service

import (
	"context"
	"micro_tiktok/cmd/comment/dal"
	"micro_tiktok/cmd/comment/rpc"
	"micro_tiktok/kitex_gen/comment"
	"testing"
)

func TestActionService_Action(t *testing.T) {
	dal.Init()
	rpc.Init()
	res, err := NewActionService(context.Background()).Action(&comment.CommentRequest{
		UserId:      2,
		VideoId:     1,
		ActionType:  2,
		CommentText: "123",
		CommentId:   1,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
