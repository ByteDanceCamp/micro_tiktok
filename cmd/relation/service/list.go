package service

import (
	"context"
	"micro_tiktok/cmd/relation/dal/db"
	"micro_tiktok/cmd/relation/dal/redis"
	"micro_tiktok/cmd/relation/pack"
	"micro_tiktok/cmd/relation/rpc"
	"micro_tiktok/kitex_gen/relation"
	"micro_tiktok/kitex_gen/user"
	"micro_tiktok/pkg/constants"
	"micro_tiktok/pkg/errno"
	"strconv"
)

type ListService struct {
	ctx context.Context
}

func NewListService(ctx context.Context) *ListService {
	return &ListService{ctx: ctx}
}

func (l *ListService) RelationList(req *relation.ListRequest) (res []*relation.User, err error) {
	// 验证目标用户是否存在
	isExist, err := rpc.IsExist(l.ctx, &user.IsExistByIdRequest{UserId: req.TargetUserId})
	if err != nil {
		return nil, err
	}
	if !isExist {
		return nil, errno.UserErr.WithMsg("user is not exist")
	}

	var key string
	toUid := strconv.Itoa(int(req.TargetUserId))
	switch req.ActionType {
	case 1:
		// 关注列表
		key = constants.RelationFollowPre + toUid
	case 2:
		// 粉丝列表
		key = constants.RelationFansPre + toUid
	default:
		return nil, errno.ParamsErr.WithMsg("action type is invalid")
	}
	// 从缓存获取目标列表
	strSlice, e := redis.GetListStrSlice(l.ctx, key, 0, -1)
	if e.ErrCode == errno.ServiceErrCode {
		return nil, e
	}
	ids, err := pack.StrSlice2IntSlice(strSlice)
	if err != nil {
		return nil, err
	}
	// 缓存命中
	if e.ErrCode == errno.SuccessCode {
		urs, err := rpc.MGetUser(l.ctx, &user.MGetUserRequest{
			TargetUserIds: ids,
			UserId:        req.UserId,
		})
		if err != nil {
			return nil, err
		}
		return pack.Users(urs), nil
	}
	// 缓存不命中
	ids = []int64{}
	switch req.ActionType {
	case 1:
		f1List, err := db.MGetFollowList(l.ctx, req.TargetUserId)
		if err != nil {
			return nil, err
		}
		for _, v := range f1List {
			ids = append(ids, v.FollowedId)
			err = redis.AddFollow(l.ctx, v)
			if err != nil {
				return nil, err
			}
		}
	case 2:
		f2List, err := db.MGetFollowerList(l.ctx, req.TargetUserId)
		if err != nil {
			return nil, err
		}
		for _, v := range f2List {
			ids = append(ids, v.FollowerId)
			err = redis.AddFollower(l.ctx, v)
			if err != nil {
				return nil, err
			}
		}
	}
	urs, err := rpc.MGetUser(l.ctx, &user.MGetUserRequest{
		TargetUserIds: ids,
		UserId:        req.UserId,
	})
	if err != nil {
		return nil, err
	}
	return pack.Users(urs), nil
}
