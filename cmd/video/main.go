package main

import (
	"micro_tiktok/cmd/video/dal"
	"micro_tiktok/cmd/video/rpc"
	video "micro_tiktok/kitex_gen/video/videoservice"
	"micro_tiktok/pkg/constants"
	tracer2 "micro_tiktok/pkg/tracer"
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
<<<<<<< HEAD:cmd/video/main.go
=======
	"micro_tiktok/cmd/user/dal"
	"micro_tiktok/cmd/user/rpc"
	user "micro_tiktok/kitex_gen/user/userservice"
	"micro_tiktok/pkg/constants"
	tracer2 "micro_tiktok/pkg/tracer"
	"net"
>>>>>>> e0ff5720c3c794e62dfbe091dfd9abb1c011fd34:cmd/user/main.go
)

func Init() {
	tracer2.InitJaeger(constants.UserServiceName)
	dal.Init()
	rpc.InitRPC()
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:18080")
	if err != nil {
		panic(err)
	}

	Init()

	svr := video.NewServer(new(VideoServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.UserServiceName}), // server name
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		server.WithSuite(trace.NewDefaultServerSuite()),                    // tracer
		server.WithRegistry(r),                                             // registry
	)

	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
