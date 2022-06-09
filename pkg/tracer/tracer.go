package tracer

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegerCfg "github.com/uber/jaeger-client-go/config"
)

func InitJaeger(service string) {
	cfg, _ := jaegerCfg.FromEnv()
	cfg.ServiceName = service
	tracer, _, err := cfg.NewTracer(jaegerCfg.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.InitGlobalTracer(tracer)
	return
}
