package service

import (
	"context"
	"micro_tiktok/cmd/video/dal"
	"micro_tiktok/cmd/video/rpc"
	"micro_tiktok/kitex_gen/video"
	"testing"
)

func TestFeedService_Feed(t *testing.T) {
	dal.Init()
	rpc.InitRPC()
	res, _, err := NewFeedService(context.Background()).Feed(&video.FeedRequest{
		Uid:        2,
		LatestTime: 0,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(res)

}
