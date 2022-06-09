package service

import (
	"context"
	"micro_tiktok/cmd/user/dal/db"
	"micro_tiktok/cmd/user/pack"
	"micro_tiktok/cmd/user/rpc"
	"micro_tiktok/kitex_gen/relation"
	"micro_tiktok/kitex_gen/user"
	"micro_tiktok/pkg/errno"
)

type MGetService struct {
	ctx context.Context
}

func NewMGet(ctx context.Context) *MGetService {
	return &MGetService{
		ctx: ctx,
	}
}

func (m *MGetService) MGet(req *user.MGetUserRequest) ([]*user.User, error) {
	if len(req.TargetUserIds) == 0 {
		return make([]*user.User, 0), nil
	}
	urs, err := db.MGet(m.ctx, req.TargetUserIds)
	if err != nil {
		return nil, err
	}
	if len(urs) == 0 {
		return nil, errno.UserErr.WithMsg("user isn't exist")
	}
	users := pack.Users(urs)
	//if isExist, err := db.IsExist(m.ctx, req.UserId); err != nil || !isExist {
	//	return users, nil
	//}
	for i, u := range users {
		countInfo, err := rpc.RelationInfo(m.ctx, &relation.InfoRequest{
			UserId:       req.UserId,
			TargetUserId: u.Id,
		})
		if err != nil {
			return users, err
		}
		users[i].FollowCount = countInfo.FollowCount
		users[i].FollowerCount = countInfo.FollowerCount
		users[i].IsFollow = countInfo.IsFollow
	}
	return users, nil
}
