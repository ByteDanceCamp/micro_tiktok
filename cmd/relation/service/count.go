package service

import (
	"context"
	"micro_tiktok/cmd/relation/dal/redis"
	"micro_tiktok/kitex_gen/relation"
	"micro_tiktok/pkg/constants"
	"strconv"
)

type CountService struct {
	ctx context.Context
}

func NewCountService(ctx context.Context) *CountService {
	return &CountService{ctx: ctx}
}

func (c *CountService) CountInfo(req *relation.InfoRequest) (info *relation.CountInfo, err error) {
	targetUid := strconv.Itoa(int(req.TargetUserId))
	uid := strconv.Itoa(int(req.UserId))
	followCount, err := redis.GetCount(c.ctx, constants.RelationFollowPre, targetUid)
	if err != nil {
		return nil, err
	}
	followerCount, err := redis.GetCount(c.ctx, constants.RelationFansPre, targetUid)
	if err != nil {
		return nil, err
	}
	isFollow := redis.IsFollow(c.ctx, constants.RelationFollowPre, uid, targetUid)
	return &relation.CountInfo{
		FollowCount:   followCount,
		FollowerCount: followerCount,
		IsFollow:      isFollow,
	}, nil
}
