package main

import (
	"context"
	"micro_tiktok/cmd/comment/pack"
	"micro_tiktok/cmd/comment/service"
	"micro_tiktok/kitex_gen/comment"
	"micro_tiktok/pkg/errno"
)

// CommentVideoServiceImpl implements the last service interface defined in the IDL.
type CommentVideoServiceImpl struct{}

// Comment implements the CommentVideoServerImpl interface.
func (s *CommentVideoServiceImpl) Comment(ctx context.Context, req *comment.CommentRequest) (resp *comment.CommentResponse, err error) {
	resp = new(comment.CommentResponse)
	res, err := service.NewActionService(ctx).Action(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.Comment = res
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// List implements the CommentVideoServerImpl interface.
func (s *CommentVideoServiceImpl) List(ctx context.Context, req *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {
	resp = new(comment.CommentListResponse)
	res, err := service.NewListService(ctx).CommentList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.CommentList = res
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// CountRes implements the CommentVideoServerImpl interface.
func (s *CommentVideoServiceImpl) CountRes(ctx context.Context, req *comment.CommentCountRequest) (resp *comment.CommentCountResponse, err error) {
	resp = new(comment.CommentCountResponse)
	res, err := service.NewCountService(ctx).GetCount(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Count = res
	return resp, nil
}
