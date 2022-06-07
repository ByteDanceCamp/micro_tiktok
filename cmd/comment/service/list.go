package service

import (
	"context"
	"micro_tiktok/cmd/comment/dal/db"
	"micro_tiktok/cmd/comment/pack"
	"micro_tiktok/cmd/comment/rpc"
	"micro_tiktok/kitex_gen/comment"
	"micro_tiktok/kitex_gen/user"
)

type ListService struct {
	ctx context.Context
}

func NewListService(ctx context.Context) *ListService {
	return &ListService{ctx: ctx}
}

func (l *ListService) CommentList(req *comment.CommentListRequest) (res []*comment.Comment, err error) {
	listDB, err := db.GetList(l.ctx, &db.Comment{
		VideoId: req.VideoId,
	})
	if err != nil {
		return nil, err
	}
	for _, v := range listDB {
		c := pack.Comment(v)
		users, err := rpc.MGetUser(l.ctx, &user.MGetUserRequest{
			TargetUserIds: []int64{c.User.Id},
			UserId:        req.UserId,
		})
		if err != nil {
			return nil, err
		}
		if len(users) == 1 {
			c.User = pack.User(users[0])
		}
		res = append(res, c)
	}
	return res, nil
}
