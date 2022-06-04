package rpc

import (
	"context"
	"micro_tiktok/kitex_gen/favorite"
	"micro_tiktok/kitex_gen/favorite/favoritevideoservice"
	"micro_tiktok/pkg/constants"
	"micro_tiktok/pkg/errno"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var favoriteClient favoritevideoservice.Client

func initFavoriteRPC() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	c, err := favoritevideoservice.NewClient(
		constants.FavoriteServiceName,
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
	favoriteClient = c
}

func GetFavoriteCount(ctx context.Context, req *favorite.VideoFavoriteCountRequest) (int64, error) {
	resp, err := favoriteClient.GetFavoriteCount(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.Count, nil
}
