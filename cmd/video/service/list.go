package service

import (
	"context"
	"micro_tiktok/cmd/video/dal/db"
	"micro_tiktok/cmd/video/pack"
	"micro_tiktok/cmd/video/rpc"
	"micro_tiktok/kitex_gen/comment"
	"micro_tiktok/kitex_gen/favorite"
	"micro_tiktok/kitex_gen/user"
	"micro_tiktok/kitex_gen/video"
	"micro_tiktok/pkg/errno"
)

type ListService struct {
	ctx context.Context
}

func NewListService(ctx context.Context) *ListService {
	return &ListService{
		ctx: ctx,
	}
}

func (l *ListService) List(req *video.ListRequest) ([]*video.Video, error) {
	users, err := rpc.MGetUser(l.ctx, &user.MGetUserRequest{
		TargetUserIds: []int64{req.TargetUid},
		UserId:        req.UserId,
	})
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errno.UserErr.WithMsg("target user isn't exist")
	}
	videosDB, err := db.GetMyList(l.ctx, req.TargetUid)
	if err != nil {
		return nil, err
	}
	videos := pack.Videos(videosDB)
	for i, v := range videos {
		videos[i].Author = pack.User(users[0])
		videos[i].CommentCount, err = rpc.CountRes(l.ctx, &comment.CommentCountRequest{VideoId: v.Id})
		if err != nil {
			videos[i].CommentCount = 0
		}
		videos[i].FavoriteCount, err = rpc.GetFavoriteCount(l.ctx, &favorite.VideoFavoriteCountRequest{VideoId: v.Id})
		if err != nil {
			videos[i].FavoriteCount = 0
		}
		videos[i].IsFavorite, err = rpc.IsFavoriteVideo(l.ctx, &favorite.IsFavoriteVideoRequest{
			VideoId: v.Id,
			UserId:  req.UserId,
		})
		if err != nil {
			videos[i].IsFavorite = false
		}
	}
	return videos, nil
}
