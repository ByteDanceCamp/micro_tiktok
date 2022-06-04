package service

import (
	"context"
	"micro_tiktok/cmd/relation/dal/db"
	"micro_tiktok/cmd/relation/dal/redis"
	"micro_tiktok/cmd/relation/rpc"
	"micro_tiktok/kitex_gen/relation"
	"micro_tiktok/kitex_gen/user"
	"micro_tiktok/pkg/errno"
)

type ActionService struct {
	ctx context.Context
}

func NewActionService(ctx context.Context) *ActionService {
	return &ActionService{ctx: ctx}
}

func (a *ActionService) Action(req *relation.ActionRequest) (err error) {
	// 验证目标用户是否存在
	isExist, _ := rpc.IsExist(a.ctx, &user.IsExistByIdRequest{UserId: req.ToUserId})
	if !isExist {
		return errno.UserErr.WithMsg("user is not exist")
	}
	// 更改数据库
	switch req.ActionType {
	case 1:
		// 关注操作
		err = db.FollowAction(a.ctx, req.UserId, req.ToUserId)
		if err != nil {
			return err
		}
	case 2:
		// 取关操作
		err = db.UnFollowAction(a.ctx, req.UserId, req.ToUserId)
		if err != nil {
			return err
		}
	default:
		return errno.ParamsErr.WithMsg("action type is invalid")
	}
	// 删除 redis 缓存
	err = redis.Action(a.ctx, &redis.ActionParams{
		Uid:   req.UserId,
		ToUid: req.ToUserId,
	})
	if err != nil {
		return err
	}
	return nil
}
