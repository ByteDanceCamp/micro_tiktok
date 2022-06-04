package main

import (
	"context"
	"micro_tiktok/kitex_gen/comment"
)

// CommentVideoServerImpl implements the last service interface defined in the IDL.
type CommentVideoServerImpl struct{}

// Comment implements the CommentVideoServerImpl interface.
func (s *CommentVideoServerImpl) Comment(ctx context.Context, req *comment.CommentRequest) (resp *comment.CommentResponse, err error) {
	// TODO: Your code here...
	return
}

// List implements the CommentVideoServerImpl interface.
func (s *CommentVideoServerImpl) List(ctx context.Context, req *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {
	// TODO: Your code here...
	return
}

// CountRes implements the CommentVideoServerImpl interface.
func (s *CommentVideoServerImpl) CountRes(ctx context.Context, req *comment.CommentCountRequest) (resp *comment.CommentCountResponse, err error) {
	// TODO: Your code here...
	return
}
