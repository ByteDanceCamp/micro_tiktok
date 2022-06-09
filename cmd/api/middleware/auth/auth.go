package auth

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/cloudwego/kitex/pkg/klog"
)

type Middleware struct {
	JWT *jwt.GinJWTMiddleware
}

func NewMiddleware(middleware jwt.GinJWTMiddleware) *Middleware {
	authMiddleware, err := jwt.New(&middleware)
	if err != nil {
		klog.Fatal(err)
	}
	return &Middleware{
		JWT: authMiddleware,
	}
}
