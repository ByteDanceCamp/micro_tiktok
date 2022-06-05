package auth

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"micro_tiktok/pkg/constants"
	"net/http"
	"time"
)

// Config 生成 JWT 中间件对象时需要的配置信息
var Config = jwt.GinJWTMiddleware{
	Key:              []byte(constants.JWTSecretKey),
	SigningAlgorithm: "HS256",
	Timeout:          time.Hour,
	MaxRefresh:       time.Hour,
	Authenticator:    nil,
	Authorizator:     nil,
	PayloadFunc: func(data interface{}) jwt.MapClaims {
		if v, ok := data.(int64); ok {
			return jwt.MapClaims{
				constants.IdentityKey: v,
			}
		}
		return jwt.MapClaims{}
	},
	Unauthorized: func(c *gin.Context, code int, message string) {
		c.JSON(http.StatusOK, gin.H{
			"status_code": code,
			"status_msg":  message,
		})
	},
	TokenLookup: "query: token, param: token",
}
