package rpc

import (
	"context"
	"micro_tiktok/kitex_gen/video"
	"micro_tiktok/kitex_gen/video/videoservice"
	"micro_tiktok/pkg/constants"
	"micro_tiktok/pkg/errno"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var videoClient videoservice.Client

func initVideoRPC() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	c, err := videoservice.NewClient(
		constants.VideoServiceName,
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	videoClient = c
}

func List(ctx context.Context, req *video.ListRequest) ([]*video.Video, error) {
	resp, err := videoClient.List(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(int64(resp.BaseResp.StatusCode), resp.BaseResp.StatusMsg)
	}
	return resp.VideoList, nil
}

func Publish(ctx context.Context, req *video.PublishRequest) error {
	resp, err := videoClient.Publish(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(int64(resp.BaseResp.StatusCode), resp.BaseResp.StatusMsg)
	}
	return nil
}

func Feed(ctx context.Context, req *video.FeedRequest) ([]*video.Video, int64, error) {
	resp, err := videoClient.Feed(ctx, req)
	if err != nil {
		return nil, 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, 0, errno.NewErrNo(int64(resp.BaseResp.StatusCode), resp.BaseResp.StatusMsg)
	}
	return resp.VideoList, resp.NextTime, nil
}
