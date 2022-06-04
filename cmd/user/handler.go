package main

import (
	"context"
	"micro_tiktok/cmd/user/pack"
	"micro_tiktok/cmd/user/service"
	"micro_tiktok/kitex_gen/user"
	"micro_tiktok/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (resp *user.CreateUserResponse, err error) {
	resp = new(user.CreateUserResponse)
	// TODO: validate the params...

	usr, err := service.NewCreateService(ctx).Create(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.UserId = usr.Id
	return resp, nil
}

// MGetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) MGetUser(ctx context.Context, req *user.MGetUserRequest) (resp *user.MGetUserResponse, err error) {
	resp = new(user.MGetUserResponse)

	// TODO: Validate the params...

	users, err := service.NewMGet(ctx).MGet(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Users = users
	return resp, nil
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *user.CheckUserRequest) (resp *user.CheckUserResponse, err error) {
	resp = new(user.CheckUserResponse)
	// TODO: validate the params...

	uid, err := service.NewCheckService(ctx).Check(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.UserId = uid
	return resp, nil
}

// IsExist implements the UserServiceImpl interface.
func (s *UserServiceImpl) IsExist(ctx context.Context, req *user.IsExistByIdRequest) (resp *user.IsExistByIdResponse, err error) {
	resp = new(user.IsExistByIdResponse)
	resp.IsExist, err = service.NewIsExistService(ctx).IsExist(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(errno.UserErr.WithMsg(err.Error()))
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
