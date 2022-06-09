package router

import (
	"micro_tiktok/cmd/api/handlers"
	"micro_tiktok/cmd/api/middleware/auth"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	authMiddleware := auth.NewMiddleware(auth.Config)

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
	video := api.Group("/publish")
	//api.GET("/feed", handlers.Feed)
	api.GET("/feed", auth.SelectMiddleWare(), handlers.Feed)
	video.POST("/action/", auth.FormMiddleWare(), handlers.VideoPublish)

	video.Use(authMiddleware.JWT.MiddlewareFunc())
	{
		video.GET("/list/", handlers.VList)
	}

	favorite := api.Group("/favorite/")
	favorite.Use(authMiddleware.JWT.MiddlewareFunc())
	{
		favorite.POST("/action/", handlers.FavoriteAction)
		favorite.GET("/list/", handlers.FavoriteVideoList)
	}
}
