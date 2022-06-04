package service

import (
	"context"
	"micro_tiktok/cmd/comment/dal"
	"micro_tiktok/cmd/comment/rpc"
	"micro_tiktok/kitex_gen/comment"
	"testing"
)

func TestListService_CommentList(t *testing.T) {
	dal.Init()
	rpc.Init()
	res, err := NewListService(context.Background()).CommentList(&comment.CommentListRequest{
		VideoId: 0,
		UserId:  0,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
