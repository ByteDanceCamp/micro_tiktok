package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"micro_tiktok/cmd/comment/dal"
	"micro_tiktok/cmd/comment/rpc"
	comment2 "micro_tiktok/kitex_gen/comment/commentvideoservice"
	"micro_tiktok/pkg/constants"
	tracer2 "micro_tiktok/pkg/tracer"
	"net"
)

func Init() {
	tracer2.InitJaeger(constants.CommentServiceName)
	dal.Init()
	rpc.Init()
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:38080")
	if err != nil {
		panic(err)
	}
	Init()
	svr := comment2.NewServer(new(CommentVideoServerImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.CommentServiceName}), // server name
		server.WithServiceAddr(addr),                                                                      // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),                                // limit
		server.WithMuxTransport(),                                                                         // Multiplex
		server.WithSuite(trace.NewDefaultServerSuite()),                                                   // tracer
		server.WithRegistry(r),
	)

	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
