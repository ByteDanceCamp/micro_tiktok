package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"micro_tiktok/kitex_gen/comment"
	"micro_tiktok/kitex_gen/comment/commentvideoservice"
	"micro_tiktok/pkg/constants"
	"micro_tiktok/pkg/errno"
	"time"
)

var commentClient commentvideoservice.Client

func initCommentRPC() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	c, err := commentvideoservice.NewClient(
		constants.CommentServiceName,
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
	commentClient = c
}

func CommentAction(ctx context.Context, req *comment.CommentRequest) (*comment.Comment, error) {
	resp, err := commentClient.Comment(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.Comment, nil
}

func CommentList(ctx context.Context, req *comment.CommentListRequest) ([]*comment.Comment, error) {
	resp, err := commentClient.List(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != errno.SuccessCode {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMsg)
	}
	return resp.CommentList, nil
}
