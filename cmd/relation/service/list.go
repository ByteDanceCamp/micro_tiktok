package service

import (
	"context"
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
	return &ListService{
		ctx: ctx,
	}
}

func (l *ListService) RelationList(req *relation.ListRequest) (res []*relation.User, err error) {
	uid, toUid := strconv.Itoa(int(req.UserId)), strconv.Itoa(int(req.TargetUserId))
	var targetKey string
	followKey := constants.RelationFollowPre + uid
	switch req.ActionType {
	case 1:
		targetKey = constants.RelationFollowPre + toUid
	case 2:
		targetKey = constants.RelationFansPre + toUid
	default:
		return nil, errno.ParamsErr.WithMsg("invalid action type")
	}
	var listStrSlice []string
	var listIntSlice []int64
	// 待查询列表
	listStrSlice, err = redis.GetListStrSlice(l.ctx, targetKey, 0, -1)
	if err != nil {
		return nil, err
	}
	listIntSlice, err = pack.StrSlice2IntSlice(listStrSlice)
	if err != nil {
		return nil, err
	}
	urs, err := rpc.MGetUser(l.ctx, &user.MGetUserRequest{UserIds: listIntSlice})
	res = pack.Users(urs)
	if err != nil {
		return nil, err
	}
	// 处理每个用户的关注、粉丝数及是否被当前用户关注
	for i, v := range res {
		res[i].FollowCount, err = redis.GetCount(l.ctx, constants.RelationFollowPre, strconv.Itoa(int(v.Id)))
		if err != nil {
			return nil, err
		}
		res[i].FollowerCount, err = redis.GetCount(l.ctx, constants.RelationFansPre, strconv.Itoa(int(v.Id)))
		if err != nil {
			return nil, err
		}
		res[i].IsFollow = redis.IsFollow(l.ctx, followKey, strconv.Itoa(int(v.Id)))
	}
	return res, nil
}
