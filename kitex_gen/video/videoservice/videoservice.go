// Code generated by Kitex v0.3.1. DO NOT EDIT.

package videoservice

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
	"google.golang.org/protobuf/proto"
	"micro_tiktok/kitex_gen/video"
)

func serviceInfo() *kitex.ServiceInfo {
	return videoServiceServiceInfo
}

var videoServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "VideoService"
	handlerType := (*video.VideoService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Feed":       kitex.NewMethodInfo(feedHandler, newFeedArgs, newFeedResult, false),
		"Publish":    kitex.NewMethodInfo(publishHandler, newPublishArgs, newPublishResult, false),
		"List":       kitex.NewMethodInfo(listHandler, newListArgs, newListResult, false),
		"MGet":       kitex.NewMethodInfo(mGetHandler, newMGetArgs, newMGetResult, false),
		"QueryByVid": kitex.NewMethodInfo(queryByVidHandler, newQueryByVidArgs, newQueryByVidResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "video.core",
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

func feedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.FeedRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.VideoService).Feed(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *FeedArgs:
		success, err := handler.(video.VideoService).Feed(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*FeedResult)
		realResult.Success = success
	}
	return nil
}
func newFeedArgs() interface{} {
	return &FeedArgs{}
}

func newFeedResult() interface{} {
	return &FeedResult{}
}

type FeedArgs struct {
	Req *video.FeedRequest
}

func (p *FeedArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in FeedArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *FeedArgs) Unmarshal(in []byte) error {
	msg := new(video.FeedRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var FeedArgs_Req_DEFAULT *video.FeedRequest

func (p *FeedArgs) GetReq() *video.FeedRequest {
	if !p.IsSetReq() {
		return FeedArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *FeedArgs) IsSetReq() bool {
	return p.Req != nil
}

type FeedResult struct {
	Success *video.FeedResponse
}

var FeedResult_Success_DEFAULT *video.FeedResponse

func (p *FeedResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in FeedResult")
	}
	return proto.Marshal(p.Success)
}

func (p *FeedResult) Unmarshal(in []byte) error {
	msg := new(video.FeedResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *FeedResult) GetSuccess() *video.FeedResponse {
	if !p.IsSetSuccess() {
		return FeedResult_Success_DEFAULT
	}
	return p.Success
}

func (p *FeedResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.FeedResponse)
}

func (p *FeedResult) IsSetSuccess() bool {
	return p.Success != nil
}

func publishHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.PublishRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.VideoService).Publish(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *PublishArgs:
		success, err := handler.(video.VideoService).Publish(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PublishResult)
		realResult.Success = success
	}
	return nil
}
func newPublishArgs() interface{} {
	return &PublishArgs{}
}

func newPublishResult() interface{} {
	return &PublishResult{}
}

type PublishArgs struct {
	Req *video.PublishRequest
}

func (p *PublishArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PublishArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *PublishArgs) Unmarshal(in []byte) error {
	msg := new(video.PublishRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PublishArgs_Req_DEFAULT *video.PublishRequest

func (p *PublishArgs) GetReq() *video.PublishRequest {
	if !p.IsSetReq() {
		return PublishArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PublishArgs) IsSetReq() bool {
	return p.Req != nil
}

type PublishResult struct {
	Success *video.PublishResponse
}

var PublishResult_Success_DEFAULT *video.PublishResponse

func (p *PublishResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PublishResult")
	}
	return proto.Marshal(p.Success)
}

func (p *PublishResult) Unmarshal(in []byte) error {
	msg := new(video.PublishResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PublishResult) GetSuccess() *video.PublishResponse {
	if !p.IsSetSuccess() {
		return PublishResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PublishResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.PublishResponse)
}

func (p *PublishResult) IsSetSuccess() bool {
	return p.Success != nil
}

func listHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.ListRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.VideoService).List(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *ListArgs:
		success, err := handler.(video.VideoService).List(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ListResult)
		realResult.Success = success
	}
	return nil
}
func newListArgs() interface{} {
	return &ListArgs{}
}

func newListResult() interface{} {
	return &ListResult{}
}

type ListArgs struct {
	Req *video.ListRequest
}

func (p *ListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ListArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *ListArgs) Unmarshal(in []byte) error {
	msg := new(video.ListRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ListArgs_Req_DEFAULT *video.ListRequest

func (p *ListArgs) GetReq() *video.ListRequest {
	if !p.IsSetReq() {
		return ListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ListArgs) IsSetReq() bool {
	return p.Req != nil
}

type ListResult struct {
	Success *video.ListResponse
}

var ListResult_Success_DEFAULT *video.ListResponse

func (p *ListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ListResult")
	}
	return proto.Marshal(p.Success)
}

func (p *ListResult) Unmarshal(in []byte) error {
	msg := new(video.ListResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ListResult) GetSuccess() *video.ListResponse {
	if !p.IsSetSuccess() {
		return ListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ListResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.ListResponse)
}

func (p *ListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func mGetHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.MGetRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.VideoService).MGet(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *MGetArgs:
		success, err := handler.(video.VideoService).MGet(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*MGetResult)
		realResult.Success = success
	}
	return nil
}
func newMGetArgs() interface{} {
	return &MGetArgs{}
}

func newMGetResult() interface{} {
	return &MGetResult{}
}

type MGetArgs struct {
	Req *video.MGetRequest
}

func (p *MGetArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in MGetArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *MGetArgs) Unmarshal(in []byte) error {
	msg := new(video.MGetRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var MGetArgs_Req_DEFAULT *video.MGetRequest

func (p *MGetArgs) GetReq() *video.MGetRequest {
	if !p.IsSetReq() {
		return MGetArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MGetArgs) IsSetReq() bool {
	return p.Req != nil
}

type MGetResult struct {
	Success *video.MGetResponse
}

var MGetResult_Success_DEFAULT *video.MGetResponse

func (p *MGetResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in MGetResult")
	}
	return proto.Marshal(p.Success)
}

func (p *MGetResult) Unmarshal(in []byte) error {
	msg := new(video.MGetResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MGetResult) GetSuccess() *video.MGetResponse {
	if !p.IsSetSuccess() {
		return MGetResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MGetResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.MGetResponse)
}

func (p *MGetResult) IsSetSuccess() bool {
	return p.Success != nil
}

func queryByVidHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(video.QueryByVidRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(video.VideoService).QueryByVid(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *QueryByVidArgs:
		success, err := handler.(video.VideoService).QueryByVid(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*QueryByVidResult)
		realResult.Success = success
	}
	return nil
}
func newQueryByVidArgs() interface{} {
	return &QueryByVidArgs{}
}

func newQueryByVidResult() interface{} {
	return &QueryByVidResult{}
}

type QueryByVidArgs struct {
	Req *video.QueryByVidRequest
}

func (p *QueryByVidArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in QueryByVidArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *QueryByVidArgs) Unmarshal(in []byte) error {
	msg := new(video.QueryByVidRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var QueryByVidArgs_Req_DEFAULT *video.QueryByVidRequest

func (p *QueryByVidArgs) GetReq() *video.QueryByVidRequest {
	if !p.IsSetReq() {
		return QueryByVidArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *QueryByVidArgs) IsSetReq() bool {
	return p.Req != nil
}

type QueryByVidResult struct {
	Success *video.QueryByVidResponse
}

var QueryByVidResult_Success_DEFAULT *video.QueryByVidResponse

func (p *QueryByVidResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in QueryByVidResult")
	}
	return proto.Marshal(p.Success)
}

func (p *QueryByVidResult) Unmarshal(in []byte) error {
	msg := new(video.QueryByVidResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *QueryByVidResult) GetSuccess() *video.QueryByVidResponse {
	if !p.IsSetSuccess() {
		return QueryByVidResult_Success_DEFAULT
	}
	return p.Success
}

func (p *QueryByVidResult) SetSuccess(x interface{}) {
	p.Success = x.(*video.QueryByVidResponse)
}

func (p *QueryByVidResult) IsSetSuccess() bool {
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

func (p *kClient) Feed(ctx context.Context, Req *video.FeedRequest) (r *video.FeedResponse, err error) {
	var _args FeedArgs
	_args.Req = Req
	var _result FeedResult
	if err = p.c.Call(ctx, "Feed", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Publish(ctx context.Context, Req *video.PublishRequest) (r *video.PublishResponse, err error) {
	var _args PublishArgs
	_args.Req = Req
	var _result PublishResult
	if err = p.c.Call(ctx, "Publish", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) List(ctx context.Context, Req *video.ListRequest) (r *video.ListResponse, err error) {
	var _args ListArgs
	_args.Req = Req
	var _result ListResult
	if err = p.c.Call(ctx, "List", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MGet(ctx context.Context, Req *video.MGetRequest) (r *video.MGetResponse, err error) {
	var _args MGetArgs
	_args.Req = Req
	var _result MGetResult
	if err = p.c.Call(ctx, "MGet", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) QueryByVid(ctx context.Context, Req *video.QueryByVidRequest) (r *video.QueryByVidResponse, err error) {
	var _args QueryByVidArgs
	_args.Req = Req
	var _result QueryByVidResult
	if err = p.c.Call(ctx, "QueryByVid", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
