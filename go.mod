module micro_tiktok

go 1.16

require (
	github.com/appleboy/gin-jwt/v2 v2.8.0
	github.com/cloudwego/kitex v0.3.1
	github.com/gin-gonic/gin v1.7.7
	github.com/go-playground/validator/v10 v10.11.0 // indirect
	github.com/go-redis/redis/v8 v8.11.5
	github.com/kitex-contrib/registry-etcd v0.0.0-20220110034026-b1c94979cea3
	github.com/kitex-contrib/tracer-opentracing v0.0.2
	github.com/opentracing/opentracing-go v1.2.0
	github.com/uber/jaeger-client-go v2.30.0+incompatible
	go.uber.org/atomic v1.9.0 // indirect
	google.golang.org/protobuf v1.28.0
	gorm.io/driver/mysql v1.3.3
	gorm.io/gorm v1.23.5
	gorm.io/plugin/opentracing v0.0.0-20211220013347-7d2b2af23560
)
