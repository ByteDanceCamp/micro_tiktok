package auth

import (
	"github.com/gin-gonic/gin"
)

func SelectMiddleWare() gin.HandlerFunc {

	return func(c *gin.Context) {
		token := c.Query("token")
		if token != "" {
			formMiddlewareImpl(c)
		}
	}
}
