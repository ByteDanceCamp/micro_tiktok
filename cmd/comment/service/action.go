package service

import (
	"context"
	"gorm.io/gorm"
	"micro_tiktok/cmd/comment/dal/db"
	"micro_tiktok/cmd/comment/pack"
	"micro_tiktok/cmd/comment/rpc"
	"micro_tiktok/kitex_gen/comment"
	"micro_tiktok/kitex_gen/user"
	"micro_tiktok/pkg/errno"
)

type ActionService struct {
	ctx context.Context
}

func NewActionService(ctx context.Context) *ActionService {
	return &ActionService{ctx: ctx}
}

func (a *ActionService) Action(req *comment.CommentRequest) (res []*comment.Comment, err error) {
	switch req.ActionType {
	case 1:
		// publish comment
		res, err = a.Publish(req)
		if err != nil {
			return nil, err
		}
		return res, nil
	case 2:
		// del comment
		err = a.Del(req)
		if err != nil {
			return nil, err
		}
		return nil, nil
	default:
		return nil, errno.ParamsErr.WithMsg("action type is invalid")
	}
}

func (a *ActionService) Publish(req *comment.CommentRequest) (res []*comment.Comment, err error) {
	commentDB, err := db.Publish(a.ctx, &db.Comment{
		VideoId: req.VideoId,
		Uid:     req.UserId,
		Content: req.CommentText,
	})
	if err != nil {
		return nil, err
	}
	ct := pack.Comment(commentDB)
	users, err := rpc.MGetUser(a.ctx, &user.MGetUserRequest{
		TargetUserIds: []int64{ct.User.Id},
		UserId:        req.UserId,
	})
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errno.UserErr.WithMsg("author don't exist")
	}
	ct.User = pack.User(users[0])
	res = append(res, ct)
	return res, nil
}

func (a *ActionService) Del(req *comment.CommentRequest) (err error) {
	err = db.Delete(a.ctx, &db.Comment{
		Model: gorm.Model{
			ID: uint(req.CommentId),
		},
	})
	if err != nil {
		return err
	}
	return nil
}
