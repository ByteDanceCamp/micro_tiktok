package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
	"micro_tiktok/cmd/api/router"
	"micro_tiktok/cmd/api/rpc"
	"micro_tiktok/pkg/constants"
	"micro_tiktok/pkg/tracer"
	"net/http"
)

func Init() {
	tracer.InitJaeger(constants.APIServiceName)
	rpc.InitRPC()
}

func main() {
	r := gin.New()
	Init()
	router.Router(r)
	if err := http.ListenAndServe(":8080", r); err != nil {
		klog.Fatal(err)
	}
}
