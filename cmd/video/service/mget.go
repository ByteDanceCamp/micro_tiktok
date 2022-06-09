package service

import (
	"context"
	"errors"
	"micro_tiktok/cmd/video/dal/db"
	"micro_tiktok/cmd/video/pack"
	"micro_tiktok/cmd/video/rpc"
	"micro_tiktok/kitex_gen/comment"
	"micro_tiktok/kitex_gen/favorite"
	"micro_tiktok/kitex_gen/user"
	"micro_tiktok/kitex_gen/video"
)

type MGetVideosService struct {
	ctx context.Context
}

func NewMGetVideosService(ctx context.Context) *MGetVideosService {
	return &MGetVideosService{ctx: ctx}
}

func (m *MGetVideosService) MGet(req *video.MGetRequest) (res []*video.Video, err error) {
	videosDB, err := db.MGet(m.ctx, req.Vids)
	if err != nil {
		return nil, err
	}
	res = pack.Videos(videosDB)
	for i, v := range res {
		users, err := rpc.MGetUser(m.ctx, &user.MGetUserRequest{
			TargetUserIds: []int64{v.Author.Id},
			UserId:        req.UserId,
		})
		if err != nil {
			return nil, err
		}
		if len(users) == 0 {
			return nil, errors.New("author is not exist")
		}
		res[i].Author = pack.User(users[0])
		countRes, err := rpc.CountRes(m.ctx, &comment.CommentCountRequest{VideoId: v.Id})
		if err != nil {
			res[i].CommentCount = 0
		} else {
			res[i].CommentCount = countRes
		}
		res[i].FavoriteCount, err = rpc.GetFavoriteCount(m.ctx, &favorite.VideoFavoriteCountRequest{VideoId: v.Id})
		if err != nil {
			res[i].FavoriteCount = 0
		}
		res[i].IsFavorite, err = rpc.IsFavoriteVideo(m.ctx, &favorite.IsFavoriteVideoRequest{
			VideoId: v.Id,
			UserId:  req.UserId,
		})
		if err != nil {
			res[i].IsFavorite = false
		}
	}
	return res, nil
}
