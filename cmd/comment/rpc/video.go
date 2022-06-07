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

func initVideoRpc() {
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
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	videoClient = c
}

func QueryByVId(ctx context.Context, req *video.QueryByVidRequest) bool {
	resp, err := videoClient.QueryByVid(ctx, req)
	if err != nil {
		return false
	}
	if resp.BaseResp.StatusCode != errno.SuccessCode {
		return false
	}
	return resp.IsExist
}
