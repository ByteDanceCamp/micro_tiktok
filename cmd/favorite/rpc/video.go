package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"micro_tiktok/kitex_gen/video"
	"micro_tiktok/kitex_gen/video/videoservice"
	"micro_tiktok/pkg/constants"
	"micro_tiktok/pkg/errno"
	"time"
)

var videoClient videoservice.Client

func initVideoRPC() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := videoservice.NewClient(
		constants.VideoServiceName,
		client.WithMuxConnection(1),
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(50*time.Millisecond),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithSuite(trace.NewDefaultClientSuite()),
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	videoClient = c
}

func GetFavoriteVideoList(ctx context.Context, req *video.ListRequest) ([]*video.Video, error) {
	resp, err := videoClient.List(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(int64(resp.BaseResp.StatusCode), resp.BaseResp.StatusMsg)
	}
	return resp.VideoList, nil
}
