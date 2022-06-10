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

type CountService struct {
	ctx context.Context
}

func NewCountService(ctx context.Context) *CountService {
	return &CountService{ctx: ctx}
}

func (c *CountService) CountInfo(req *relation.InfoRequest) (res *relation.CountInfo, err error) {
	// 验证目标用户是否存在
	isExist, err := rpc.IsExist(c.ctx, &user.IsExistByIdRequest{UserId: req.TargetUserId})
	if err != nil {
		return nil, err
	}
	if !isExist {
		return nil, errno.UserErr.WithMsg("user is not exist")
	}
	res = &relation.CountInfo{
		FollowCount:   0,
		FollowerCount: 0,
		IsFollow:      false,
	}

	// 查缓存
	count, e := redis.GetCount(c.ctx, req.TargetUserId)
	// 查询出错
	if e.ErrCode == errno.ServiceErrCode {
		return nil, err
	}
	// 总数表缓存命中
	if e.ErrCode == errno.SuccessCode {
		res.FollowCount = count.FollowCount
		res.FollowerCount = count.FollowerCount
		isFollow, err := c.IsFollow(req)
		if err != nil {
			return nil, err
		}
		res.IsFollow = isFollow
		return res, nil
	}

	// 总数表缓存不命中
	result, err := db.MGetCount(c.ctx, []int64{req.TargetUserId})
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		// 数据库没有记录
		record, err := db.AddRecord(c.ctx, req.TargetUserId)
		if err != nil {
			return nil, err
		}
		return &relation.CountInfo{
			FollowCount:   record.FollowCount,
			FollowerCount: record.FollowerCount,
			IsFollow:      false,
		}, nil
	}
	for _, v := range result {
		err = redis.UpdateCount(c.ctx, &redis.CountInfo{
			Uid:           v.Uid,
			FollowCount:   v.FollowCount,
			FollowerCount: v.FollowerCount,
		})
		if err != nil {
			return nil, err
		}
	}
	isFollow, err := c.IsFollow(req)
	if err != nil {
		return nil, err
	}
	res.IsFollow = isFollow
	res.FollowCount = result[0].FollowerCount
	res.FollowerCount = result[0].FollowerCount
	return res, nil

}

func (c *CountService) IsFollow(req *relation.InfoRequest) (res bool, err error) {
	if req.UserId <= 0 {
		return false, nil
	}
	isFollow, e := redis.IsFollow(c.ctx, &redis.ActionParams{
		Uid:   req.UserId,
		ToUid: req.TargetUserId,
	})
	// 关注表缓存命中
	if e.ErrCode == errno.SuccessCode {
		res = isFollow
		return res, nil
	}
	// 关注表缓存不命中
	res, err = db.IsFollow(c.ctx, req.UserId, req.TargetUserId)
	if err != nil {
		return false, err
	}
	// 将数据库中的数据刷新到 redis 中
	result, _ := db.MGetFollowList(c.ctx, req.UserId)
	if result == nil || len(result) == 0 {
		return res, nil
	}
	for _, v := range result {
		err = redis.AddFollow(c.ctx, v)
		if err != nil {
			return false, err
		}
	}
	return res, nil
}
