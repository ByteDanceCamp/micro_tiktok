// Code generated by Kitex v0.3.1. DO NOT EDIT.

package favoritevideoservice

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
	"google.golang.org/protobuf/proto"
	"micro_tiktok/kitex_gen/favorite"
)

func serviceInfo() *kitex.ServiceInfo {
	return favoriteVideoServiceServiceInfo
}

var favoriteVideoServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "FavoriteVideoService"
	handlerType := (*favorite.FavoriteVideoService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Favorite":         kitex.NewMethodInfo(favoriteHandler, newFavoriteArgs, newFavoriteResult, false),
		"FavoriteList":     kitex.NewMethodInfo(favoriteListHandler, newFavoriteListArgs, newFavoriteListResult, false),
		"GetFavoriteCount": kitex.NewMethodInfo(getFavoriteCountHandler, newGetFavoriteCountArgs, newGetFavoriteCountResult, false),
		"IsFavoriteVideo":  kitex.NewMethodInfo(isFavoriteVideoHandler, newIsFavoriteVideoArgs, newIsFavoriteVideoResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "favorite.video",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.3.1",
		Extra:           extra,
	}
	return svcInfo
}

func favoriteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(favorite.FavoriteRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(favorite.FavoriteVideoService).Favorite(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FavoriteArgs:
		success, err := handler.(favorite.FavoriteVideoService).Favorite(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FavoriteResult)
		realResult.Success = success
	}
	return nil
}
func newFavoriteArgs() interface{} {
	return &FavoriteArgs{}
}

func newFavoriteResult() interface{} {
	return &FavoriteResult{}
}

type FavoriteArgs struct {
	Req *favorite.FavoriteRequest
}

func (p *FavoriteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FavoriteArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FavoriteArgs) Unmarshal(in []byte) error {
	msg := new(favorite.FavoriteRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FavoriteArgs_Req_DEFAULT *favorite.FavoriteRequest

func (p *FavoriteArgs) GetReq() *favorite.FavoriteRequest {
	if !p.IsSetReq() {
		return FavoriteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FavoriteArgs) IsSetReq() bool {
	return p.Req != nil
}

type FavoriteResult struct {
	Success *favorite.FavoriteResponse
}

var FavoriteResult_Success_DEFAULT *favorite.FavoriteResponse

func (p *FavoriteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FavoriteResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FavoriteResult) Unmarshal(in []byte) error {
	msg := new(favorite.FavoriteResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FavoriteResult) GetSuccess() *favorite.FavoriteResponse {
	if !p.IsSetSuccess() {
		return FavoriteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FavoriteResult) SetSuccess(x interface{}) {
	p.Success = x.(*favorite.FavoriteResponse)
}

func (p *FavoriteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func favoriteListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(favorite.FavoriteListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(favorite.FavoriteVideoService).FavoriteList(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FavoriteListArgs:
		success, err := handler.(favorite.FavoriteVideoService).FavoriteList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FavoriteListResult)
		realResult.Success = success
	}
	return nil
}
func newFavoriteListArgs() interface{} {
	return &FavoriteListArgs{}
}

func newFavoriteListResult() interface{} {
	return &FavoriteListResult{}
}

type FavoriteListArgs struct {
	Req *favorite.FavoriteListRequest
}

func (p *FavoriteListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FavoriteListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FavoriteListArgs) Unmarshal(in []byte) error {
	msg := new(favorite.FavoriteListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FavoriteListArgs_Req_DEFAULT *favorite.FavoriteListRequest

func (p *FavoriteListArgs) GetReq() *favorite.FavoriteListRequest {
	if !p.IsSetReq() {
		return FavoriteListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FavoriteListArgs) IsSetReq() bool {
	return p.Req != nil
}

type FavoriteListResult struct {
	Success *favorite.FavoriteListResponse
}

var FavoriteListResult_Success_DEFAULT *favorite.FavoriteListResponse

func (p *FavoriteListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FavoriteListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FavoriteListResult) Unmarshal(in []byte) error {
	msg := new(favorite.FavoriteListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FavoriteListResult) GetSuccess() *favorite.FavoriteListResponse {
	if !p.IsSetSuccess() {
		return FavoriteListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FavoriteListResult) SetSuccess(x interface{}) {
	p.Success = x.(*favorite.FavoriteListResponse)
}

func (p *FavoriteListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getFavoriteCountHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(favorite.VideoFavoriteCountRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(favorite.FavoriteVideoService).GetFavoriteCount(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetFavoriteCountArgs:
		success, err := handler.(favorite.FavoriteVideoService).GetFavoriteCount(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetFavoriteCountResult)
		realResult.Success = success
	}
	return nil
}
func newGetFavoriteCountArgs() interface{} {
	return &GetFavoriteCountArgs{}
}

func newGetFavoriteCountResult() interface{} {
	return &GetFavoriteCountResult{}
}

type GetFavoriteCountArgs struct {
	Req *favorite.VideoFavoriteCountRequest
}

func (p *GetFavoriteCountArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetFavoriteCountArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetFavoriteCountArgs) Unmarshal(in []byte) error {
	msg := new(favorite.VideoFavoriteCountRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetFavoriteCountArgs_Req_DEFAULT *favorite.VideoFavoriteCountRequest

func (p *GetFavoriteCountArgs) GetReq() *favorite.VideoFavoriteCountRequest {
	if !p.IsSetReq() {
		return GetFavoriteCountArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetFavoriteCountArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetFavoriteCountResult struct {
	Success *favorite.VideoFavoriteCountResponse
}

var GetFavoriteCountResult_Success_DEFAULT *favorite.VideoFavoriteCountResponse

func (p *GetFavoriteCountResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetFavoriteCountResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetFavoriteCountResult) Unmarshal(in []byte) error {
	msg := new(favorite.VideoFavoriteCountResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetFavoriteCountResult) GetSuccess() *favorite.VideoFavoriteCountResponse {
	if !p.IsSetSuccess() {
		return GetFavoriteCountResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetFavoriteCountResult) SetSuccess(x interface{}) {
	p.Success = x.(*favorite.VideoFavoriteCountResponse)
}

func (p *GetFavoriteCountResult) IsSetSuccess() bool {
	return p.Success != nil
}

func isFavoriteVideoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(favorite.IsFavoriteVideoRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(favorite.FavoriteVideoService).IsFavoriteVideo(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *IsFavoriteVideoArgs:
		success, err := handler.(favorite.FavoriteVideoService).IsFavoriteVideo(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*IsFavoriteVideoResult)
		realResult.Success = success
	}
	return nil
}
func newIsFavoriteVideoArgs() interface{} {
	return &IsFavoriteVideoArgs{}
}

func newIsFavoriteVideoResult() interface{} {
	return &IsFavoriteVideoResult{}
}

type IsFavoriteVideoArgs struct {
	Req *favorite.IsFavoriteVideoRequest
}

func (p *IsFavoriteVideoArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in IsFavoriteVideoArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *IsFavoriteVideoArgs) Unmarshal(in []byte) error {
	msg := new(favorite.IsFavoriteVideoRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var IsFavoriteVideoArgs_Req_DEFAULT *favorite.IsFavoriteVideoRequest

func (p *IsFavoriteVideoArgs) GetReq() *favorite.IsFavoriteVideoRequest {
	if !p.IsSetReq() {
		return IsFavoriteVideoArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *IsFavoriteVideoArgs) IsSetReq() bool {
	return p.Req != nil
}

type IsFavoriteVideoResult struct {
	Success *favorite.IsFavoriteVideoResponse
}

var IsFavoriteVideoResult_Success_DEFAULT *favorite.IsFavoriteVideoResponse

func (p *IsFavoriteVideoResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in IsFavoriteVideoResult")
	}
	return proto.Marshal(p.Success)
}

func (p *IsFavoriteVideoResult) Unmarshal(in []byte) error {
	msg := new(favorite.IsFavoriteVideoResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *IsFavoriteVideoResult) GetSuccess() *favorite.IsFavoriteVideoResponse {
	if !p.IsSetSuccess() {
		return IsFavoriteVideoResult_Success_DEFAULT
	}
	return p.Success
}

func (p *IsFavoriteVideoResult) SetSuccess(x interface{}) {
	p.Success = x.(*favorite.IsFavoriteVideoResponse)
}

func (p *IsFavoriteVideoResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Favorite(ctx context.Context, Req *favorite.FavoriteRequest) (r *favorite.FavoriteResponse, err error) {
	var _args FavoriteArgs
	_args.Req = Req
	var _result FavoriteResult
	if err = p.c.Call(ctx, "Favorite", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FavoriteList(ctx context.Context, Req *favorite.FavoriteListRequest) (r *favorite.FavoriteListResponse, err error) {
	var _args FavoriteListArgs
	_args.Req = Req
	var _result FavoriteListResult
	if err = p.c.Call(ctx, "FavoriteList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetFavoriteCount(ctx context.Context, Req *favorite.VideoFavoriteCountRequest) (r *favorite.VideoFavoriteCountResponse, err error) {
	var _args GetFavoriteCountArgs
	_args.Req = Req
	var _result GetFavoriteCountResult
	if err = p.c.Call(ctx, "GetFavoriteCount", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) IsFavoriteVideo(ctx context.Context, Req *favorite.IsFavoriteVideoRequest) (r *favorite.IsFavoriteVideoResponse, err error) {
	var _args IsFavoriteVideoArgs
	_args.Req = Req
	var _result IsFavoriteVideoResult
	if err = p.c.Call(ctx, "IsFavoriteVideo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
