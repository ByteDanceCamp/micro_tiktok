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

type InfoService struct {
	ctx context.Context
}

func NewInfoService(ctx context.Context) *InfoService {
	return &InfoService{ctx: ctx}
}

func (i *InfoService) Info(req *relation.InfoRequest) (res *relation.User, err error) {
	users, err := rpc.MGetUser(i.ctx, &user.MGetUserRequest{UserIds: []int64{req.TargetUserId}})
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errno.ParamsErr.WithMsg("usr is not exits")
	}
	usr := pack.User(users[0])
	usr.FollowCount, err = redis.GetCount(i.ctx, constants.RelationFollowPre, strconv.Itoa(int(usr.Id)))
	usr.FollowerCount, err = redis.GetCount(i.ctx, constants.RelationFansPre, strconv.Itoa(int(usr.Id)))
	usr.IsFollow = redis.IsFollow(i.ctx, constants.RelationFollowPre+strconv.Itoa(int(req.UserId)), strconv.Itoa(int(usr.Id)))
	return usr, nil
}
