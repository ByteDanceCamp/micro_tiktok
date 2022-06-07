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
)

type FeedService struct {
	ctx context.Context
}

func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{
		ctx: ctx,
	}
}

func (l *FeedService) Feed(req *video.FeedRequest) (res []*video.Video, nextTime int64, err error) {
	vs, nextTime, err := db.GetNewestList(l.ctx, req.LatestTime)
	if err != nil {
		return nil, nextTime, err
	}
	res = pack.Videos(vs)
	for i, v := range res {
		// 获取作者信息
		u, err := rpc.MGetUser(l.ctx, &user.MGetUserRequest{
			TargetUserIds: []int64{v.Author.Id},
			UserId:        req.Uid,
		})
		if err != nil {
			return nil, nextTime, err
		}
		res[i].Author = pack.User(u[0])
		// 获取评论总数
		countRes, err := rpc.CountRes(l.ctx, &comment.CommentCountRequest{VideoId: v.Id})
		if err != nil {
			res[i].CommentCount = 0
		} else {
			res[i].CommentCount = countRes
		}

		res[i].FavoriteCount, err = rpc.GetFavoriteCount(l.ctx, &favorite.VideoFavoriteCountRequest{VideoId: v.Id})
		if err != nil {
			res[i].FavoriteCount = 0
		}
		res[i].IsFavorite, err = rpc.IsFavoriteVideo(l.ctx, &favorite.IsFavoriteVideoRequest{
			VideoId: v.Id,
			UserId:  req.Uid,
		})
		if err != nil {
			res[i].IsFavorite = false
		}
	}
	return res, nextTime, nil
}
