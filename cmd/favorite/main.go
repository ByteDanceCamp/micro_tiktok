package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"micro_tiktok/cmd/favorite/dal"
	"micro_tiktok/cmd/favorite/rpc"
	"micro_tiktok/kitex_gen/favorite/favoritevideoservice"
	"micro_tiktok/pkg/constants"
	tracer2 "micro_tiktok/pkg/tracer"
	"net"
)

func Init() {
	tracer2.InitJaeger(constants.FavoriteServiceName)
	dal.Init()
	rpc.INitRPC()
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:48080")
	if err != nil {
		panic(err)
	}
	Init()
	svr := favoritevideoservice.NewServer(new(FavoriteVideoServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.FavoriteServiceName}),
		server.WithServiceAddr(addr),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithMuxTransport(),
		server.WithSuite(trace.NewDefaultServerSuite()),
		server.WithRegistry(r),
	)

	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}

}
