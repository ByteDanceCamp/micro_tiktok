package router

import (
	"github.com/gin-gonic/gin"
	"micro_tiktok/cmd/api/handlers"
	"micro_tiktok/cmd/api/middleware/auth"
)

func Router(r *gin.Engine) {
	authMiddleware := auth.NewMiddleware(handlers.AuthConfig)

	api := r.Group("/douyin")
	user := api.Group("/user/")
	user.POST("/login/", handlers.Login)
	user.POST("/register/", handlers.Register)
	user.Use(authMiddleware.JWT.MiddlewareFunc())
	{
		user.GET("", handlers.UserInfo)
	}
	relation := api.Group("/relation/")
	relation.Use(authMiddleware.JWT.MiddlewareFunc())
	{
		relation.POST("/action/", handlers.RelationAction)
		relation.GET("/follow/list/", handlers.FollowList)
		relation.GET("/follower/list/", handlers.FollowerList)
	}
	comment := api.Group("/comment/")
	comment.Use(authMiddleware.JWT.MiddlewareFunc())
	{
		comment.POST("/action/", handlers.CommentAction)
		comment.GET("/list/", handlers.CommentList)
	}
}
