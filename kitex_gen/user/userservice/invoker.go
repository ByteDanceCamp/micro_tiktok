// Code generated by Kitex v0.3.1. DO NOT EDIT.

package userservice

import (
	"github.com/cloudwego/kitex/server"
	"micro_tiktok/kitex_gen/user"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler user.UserService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
