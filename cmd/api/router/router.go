package router

import (
	"micro_tiktok/cmd/api/handlers"
	"micro_tiktok/cmd/api/middleware/auth"

	"github.com/gin-gonic/gin"
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

	favorite := api.Group("/favorite/")
	favorite.Use(authMiddleware.JWT.MiddlewareFunc())
	{
		favorite.POST("/action/", handlers.FavoriteAction)
		favorite.GET("/list/", handlers.FavoriteVideoList)
	}
	api.Use(authMiddleware.JWT.MiddlewareFunc())
	{
		api.GET("/feed", handlers.Feed)
	}
	video := api.Group("/publish/")
	video.Use(authMiddleware.JWT.MiddlewareFunc())
	{
		video.POST("/action/", handlers.VideoPublish)
		video.GET("/list/", handlers.VideoList)
	}
}
