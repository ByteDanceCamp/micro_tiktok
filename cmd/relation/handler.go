package main

import (
	"context"
	"micro_tiktok/cmd/relation/pack"
	"micro_tiktok/cmd/relation/service"
	"micro_tiktok/kitex_gen/relation"
	"micro_tiktok/pkg/errno"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// Action implements the RelationServerImpl interface.
func (s *RelationServiceImpl) Action(ctx context.Context, req *relation.ActionRequest) (resp *relation.ActionResponse, err error) {
	resp = new(relation.ActionResponse)
	err = service.NewActionService(ctx).Action(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// List implements the RelationServerImpl interface.
func (s *RelationServiceImpl) List(ctx context.Context, req *relation.ListRequest) (resp *relation.ListResponse, err error) {
	resp = new(relation.ListResponse)
	resp.UserList, err = service.NewListService(ctx).RelationList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// Info implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) Info(ctx context.Context, req *relation.InfoRequest) (resp *relation.InfoResponse, err error) {
	resp = new(relation.InfoResponse)
	resp.User, err = service.NewInfoService(ctx).Info(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
