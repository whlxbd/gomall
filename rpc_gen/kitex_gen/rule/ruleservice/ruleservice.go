// Code generated by Kitex v0.9.1. DO NOT EDIT.

package ruleservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	rule "github.com/whlxbd/gomall/rpc_gen/kitex_gen/rule"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Create": kitex.NewMethodInfo(
		createHandler,
		newCreateArgs,
		newCreateResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"List": kitex.NewMethodInfo(
		listHandler,
		newListArgs,
		newListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"Delete": kitex.NewMethodInfo(
		deleteHandler,
		newDeleteArgs,
		newDeleteResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"Get": kitex.NewMethodInfo(
		getHandler,
		newGetArgs,
		newGetResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"Update": kitex.NewMethodInfo(
		updateHandler,
		newUpdateArgs,
		newUpdateResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"AddWhiteRouter": kitex.NewMethodInfo(
		addWhiteRouterHandler,
		newAddWhiteRouterArgs,
		newAddWhiteRouterResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"GetWhiteList": kitex.NewMethodInfo(
		getWhiteListHandler,
		newGetWhiteListArgs,
		newGetWhiteListResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"DeleteWhiteRouter": kitex.NewMethodInfo(
		deleteWhiteRouterHandler,
		newDeleteWhiteRouterArgs,
		newDeleteWhiteRouterResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	ruleServiceServiceInfo                = NewServiceInfo()
	ruleServiceServiceInfoForClient       = NewServiceInfoForClient()
	ruleServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return ruleServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return ruleServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return ruleServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "RuleService"
	handlerType := (*rule.RuleService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "rule",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.9.1",
		Extra:           extra,
	}
	return svcInfo
}

func createHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(rule.CreateReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(rule.RuleService).Create(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *CreateArgs:
		success, err := handler.(rule.RuleService).Create(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newCreateArgs() interface{} {
	return &CreateArgs{}
}

func newCreateResult() interface{} {
	return &CreateResult{}
}

type CreateArgs struct {
	Req *rule.CreateReq
}

func (p *CreateArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(rule.CreateReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *CreateArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *CreateArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *CreateArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *CreateArgs) Unmarshal(in []byte) error {
	msg := new(rule.CreateReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateArgs_Req_DEFAULT *rule.CreateReq

func (p *CreateArgs) GetReq() *rule.CreateReq {
	if !p.IsSetReq() {
		return CreateArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *CreateArgs) GetFirstArgument() interface{} {
	return p.Req
}

type CreateResult struct {
	Success *rule.CreateResp
}

var CreateResult_Success_DEFAULT *rule.CreateResp

func (p *CreateResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(rule.CreateResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *CreateResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *CreateResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *CreateResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *CreateResult) Unmarshal(in []byte) error {
	msg := new(rule.CreateResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateResult) GetSuccess() *rule.CreateResp {
	if !p.IsSetSuccess() {
		return CreateResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateResult) SetSuccess(x interface{}) {
	p.Success = x.(*rule.CreateResp)
}

func (p *CreateResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CreateResult) GetResult() interface{} {
	return p.Success
}

func listHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(rule.ListReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(rule.RuleService).List(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ListArgs:
		success, err := handler.(rule.RuleService).List(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ListResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newListArgs() interface{} {
	return &ListArgs{}
}

func newListResult() interface{} {
	return &ListResult{}
}

type ListArgs struct {
	Req *rule.ListReq
}

func (p *ListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(rule.ListReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ListArgs) Unmarshal(in []byte) error {
	msg := new(rule.ListReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ListArgs_Req_DEFAULT *rule.ListReq

func (p *ListArgs) GetReq() *rule.ListReq {
	if !p.IsSetReq() {
		return ListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ListArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ListArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ListResult struct {
	Success *rule.ListResp
}

var ListResult_Success_DEFAULT *rule.ListResp

func (p *ListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(rule.ListResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ListResult) Unmarshal(in []byte) error {
	msg := new(rule.ListResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ListResult) GetSuccess() *rule.ListResp {
	if !p.IsSetSuccess() {
		return ListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ListResult) SetSuccess(x interface{}) {
	p.Success = x.(*rule.ListResp)
}

func (p *ListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ListResult) GetResult() interface{} {
	return p.Success
}

func deleteHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(rule.DeleteReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(rule.RuleService).Delete(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *DeleteArgs:
		success, err := handler.(rule.RuleService).Delete(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeleteResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newDeleteArgs() interface{} {
	return &DeleteArgs{}
}

func newDeleteResult() interface{} {
	return &DeleteResult{}
}

type DeleteArgs struct {
	Req *rule.DeleteReq
}

func (p *DeleteArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(rule.DeleteReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DeleteArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DeleteArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DeleteArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *DeleteArgs) Unmarshal(in []byte) error {
	msg := new(rule.DeleteReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeleteArgs_Req_DEFAULT *rule.DeleteReq

func (p *DeleteArgs) GetReq() *rule.DeleteReq {
	if !p.IsSetReq() {
		return DeleteArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *DeleteArgs) GetFirstArgument() interface{} {
	return p.Req
}

type DeleteResult struct {
	Success *rule.DeleteResp
}

var DeleteResult_Success_DEFAULT *rule.DeleteResp

func (p *DeleteResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(rule.DeleteResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DeleteResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DeleteResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DeleteResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *DeleteResult) Unmarshal(in []byte) error {
	msg := new(rule.DeleteResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteResult) GetSuccess() *rule.DeleteResp {
	if !p.IsSetSuccess() {
		return DeleteResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteResult) SetSuccess(x interface{}) {
	p.Success = x.(*rule.DeleteResp)
}

func (p *DeleteResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteResult) GetResult() interface{} {
	return p.Success
}

func getHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(rule.GetReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(rule.RuleService).Get(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetArgs:
		success, err := handler.(rule.RuleService).Get(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetArgs() interface{} {
	return &GetArgs{}
}

func newGetResult() interface{} {
	return &GetResult{}
}

type GetArgs struct {
	Req *rule.GetReq
}

func (p *GetArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(rule.GetReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetArgs) Unmarshal(in []byte) error {
	msg := new(rule.GetReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetArgs_Req_DEFAULT *rule.GetReq

func (p *GetArgs) GetReq() *rule.GetReq {
	if !p.IsSetReq() {
		return GetArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetResult struct {
	Success *rule.GetResp
}

var GetResult_Success_DEFAULT *rule.GetResp

func (p *GetResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(rule.GetResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetResult) Unmarshal(in []byte) error {
	msg := new(rule.GetResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetResult) GetSuccess() *rule.GetResp {
	if !p.IsSetSuccess() {
		return GetResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetResult) SetSuccess(x interface{}) {
	p.Success = x.(*rule.GetResp)
}

func (p *GetResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetResult) GetResult() interface{} {
	return p.Success
}

func updateHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(rule.UpdateReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(rule.RuleService).Update(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *UpdateArgs:
		success, err := handler.(rule.RuleService).Update(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdateResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newUpdateArgs() interface{} {
	return &UpdateArgs{}
}

func newUpdateResult() interface{} {
	return &UpdateResult{}
}

type UpdateArgs struct {
	Req *rule.UpdateReq
}

func (p *UpdateArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(rule.UpdateReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UpdateArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UpdateArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UpdateArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *UpdateArgs) Unmarshal(in []byte) error {
	msg := new(rule.UpdateReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdateArgs_Req_DEFAULT *rule.UpdateReq

func (p *UpdateArgs) GetReq() *rule.UpdateReq {
	if !p.IsSetReq() {
		return UpdateArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UpdateArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UpdateResult struct {
	Success *rule.UpdateResp
}

var UpdateResult_Success_DEFAULT *rule.UpdateResp

func (p *UpdateResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(rule.UpdateResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UpdateResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UpdateResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UpdateResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *UpdateResult) Unmarshal(in []byte) error {
	msg := new(rule.UpdateResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateResult) GetSuccess() *rule.UpdateResp {
	if !p.IsSetSuccess() {
		return UpdateResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateResult) SetSuccess(x interface{}) {
	p.Success = x.(*rule.UpdateResp)
}

func (p *UpdateResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UpdateResult) GetResult() interface{} {
	return p.Success
}

func addWhiteRouterHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(rule.AddWhiteRouterReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(rule.RuleService).AddWhiteRouter(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *AddWhiteRouterArgs:
		success, err := handler.(rule.RuleService).AddWhiteRouter(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*AddWhiteRouterResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newAddWhiteRouterArgs() interface{} {
	return &AddWhiteRouterArgs{}
}

func newAddWhiteRouterResult() interface{} {
	return &AddWhiteRouterResult{}
}

type AddWhiteRouterArgs struct {
	Req *rule.AddWhiteRouterReq
}

func (p *AddWhiteRouterArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(rule.AddWhiteRouterReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *AddWhiteRouterArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *AddWhiteRouterArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *AddWhiteRouterArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *AddWhiteRouterArgs) Unmarshal(in []byte) error {
	msg := new(rule.AddWhiteRouterReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var AddWhiteRouterArgs_Req_DEFAULT *rule.AddWhiteRouterReq

func (p *AddWhiteRouterArgs) GetReq() *rule.AddWhiteRouterReq {
	if !p.IsSetReq() {
		return AddWhiteRouterArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *AddWhiteRouterArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *AddWhiteRouterArgs) GetFirstArgument() interface{} {
	return p.Req
}

type AddWhiteRouterResult struct {
	Success *rule.AddWhiteRouterResp
}

var AddWhiteRouterResult_Success_DEFAULT *rule.AddWhiteRouterResp

func (p *AddWhiteRouterResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(rule.AddWhiteRouterResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *AddWhiteRouterResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *AddWhiteRouterResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *AddWhiteRouterResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *AddWhiteRouterResult) Unmarshal(in []byte) error {
	msg := new(rule.AddWhiteRouterResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *AddWhiteRouterResult) GetSuccess() *rule.AddWhiteRouterResp {
	if !p.IsSetSuccess() {
		return AddWhiteRouterResult_Success_DEFAULT
	}
	return p.Success
}

func (p *AddWhiteRouterResult) SetSuccess(x interface{}) {
	p.Success = x.(*rule.AddWhiteRouterResp)
}

func (p *AddWhiteRouterResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *AddWhiteRouterResult) GetResult() interface{} {
	return p.Success
}

func getWhiteListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(rule.GetWhiteListReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(rule.RuleService).GetWhiteList(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *GetWhiteListArgs:
		success, err := handler.(rule.RuleService).GetWhiteList(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetWhiteListResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newGetWhiteListArgs() interface{} {
	return &GetWhiteListArgs{}
}

func newGetWhiteListResult() interface{} {
	return &GetWhiteListResult{}
}

type GetWhiteListArgs struct {
	Req *rule.GetWhiteListReq
}

func (p *GetWhiteListArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(rule.GetWhiteListReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *GetWhiteListArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *GetWhiteListArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *GetWhiteListArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *GetWhiteListArgs) Unmarshal(in []byte) error {
	msg := new(rule.GetWhiteListReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetWhiteListArgs_Req_DEFAULT *rule.GetWhiteListReq

func (p *GetWhiteListArgs) GetReq() *rule.GetWhiteListReq {
	if !p.IsSetReq() {
		return GetWhiteListArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetWhiteListArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *GetWhiteListArgs) GetFirstArgument() interface{} {
	return p.Req
}

type GetWhiteListResult struct {
	Success *rule.GetWhiteListResp
}

var GetWhiteListResult_Success_DEFAULT *rule.GetWhiteListResp

func (p *GetWhiteListResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(rule.GetWhiteListResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *GetWhiteListResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *GetWhiteListResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *GetWhiteListResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *GetWhiteListResult) Unmarshal(in []byte) error {
	msg := new(rule.GetWhiteListResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetWhiteListResult) GetSuccess() *rule.GetWhiteListResp {
	if !p.IsSetSuccess() {
		return GetWhiteListResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetWhiteListResult) SetSuccess(x interface{}) {
	p.Success = x.(*rule.GetWhiteListResp)
}

func (p *GetWhiteListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetWhiteListResult) GetResult() interface{} {
	return p.Success
}

func deleteWhiteRouterHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(rule.DeleteWhiteRouterReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(rule.RuleService).DeleteWhiteRouter(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *DeleteWhiteRouterArgs:
		success, err := handler.(rule.RuleService).DeleteWhiteRouter(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*DeleteWhiteRouterResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newDeleteWhiteRouterArgs() interface{} {
	return &DeleteWhiteRouterArgs{}
}

func newDeleteWhiteRouterResult() interface{} {
	return &DeleteWhiteRouterResult{}
}

type DeleteWhiteRouterArgs struct {
	Req *rule.DeleteWhiteRouterReq
}

func (p *DeleteWhiteRouterArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(rule.DeleteWhiteRouterReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *DeleteWhiteRouterArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *DeleteWhiteRouterArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *DeleteWhiteRouterArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *DeleteWhiteRouterArgs) Unmarshal(in []byte) error {
	msg := new(rule.DeleteWhiteRouterReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeleteWhiteRouterArgs_Req_DEFAULT *rule.DeleteWhiteRouterReq

func (p *DeleteWhiteRouterArgs) GetReq() *rule.DeleteWhiteRouterReq {
	if !p.IsSetReq() {
		return DeleteWhiteRouterArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeleteWhiteRouterArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *DeleteWhiteRouterArgs) GetFirstArgument() interface{} {
	return p.Req
}

type DeleteWhiteRouterResult struct {
	Success *rule.DeleteWhiteRouterResp
}

var DeleteWhiteRouterResult_Success_DEFAULT *rule.DeleteWhiteRouterResp

func (p *DeleteWhiteRouterResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(rule.DeleteWhiteRouterResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *DeleteWhiteRouterResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *DeleteWhiteRouterResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *DeleteWhiteRouterResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *DeleteWhiteRouterResult) Unmarshal(in []byte) error {
	msg := new(rule.DeleteWhiteRouterResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeleteWhiteRouterResult) GetSuccess() *rule.DeleteWhiteRouterResp {
	if !p.IsSetSuccess() {
		return DeleteWhiteRouterResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeleteWhiteRouterResult) SetSuccess(x interface{}) {
	p.Success = x.(*rule.DeleteWhiteRouterResp)
}

func (p *DeleteWhiteRouterResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *DeleteWhiteRouterResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Create(ctx context.Context, Req *rule.CreateReq) (r *rule.CreateResp, err error) {
	var _args CreateArgs
	_args.Req = Req
	var _result CreateResult
	if err = p.c.Call(ctx, "Create", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) List(ctx context.Context, Req *rule.ListReq) (r *rule.ListResp, err error) {
	var _args ListArgs
	_args.Req = Req
	var _result ListResult
	if err = p.c.Call(ctx, "List", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Delete(ctx context.Context, Req *rule.DeleteReq) (r *rule.DeleteResp, err error) {
	var _args DeleteArgs
	_args.Req = Req
	var _result DeleteResult
	if err = p.c.Call(ctx, "Delete", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Get(ctx context.Context, Req *rule.GetReq) (r *rule.GetResp, err error) {
	var _args GetArgs
	_args.Req = Req
	var _result GetResult
	if err = p.c.Call(ctx, "Get", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Update(ctx context.Context, Req *rule.UpdateReq) (r *rule.UpdateResp, err error) {
	var _args UpdateArgs
	_args.Req = Req
	var _result UpdateResult
	if err = p.c.Call(ctx, "Update", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AddWhiteRouter(ctx context.Context, Req *rule.AddWhiteRouterReq) (r *rule.AddWhiteRouterResp, err error) {
	var _args AddWhiteRouterArgs
	_args.Req = Req
	var _result AddWhiteRouterResult
	if err = p.c.Call(ctx, "AddWhiteRouter", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetWhiteList(ctx context.Context, Req *rule.GetWhiteListReq) (r *rule.GetWhiteListResp, err error) {
	var _args GetWhiteListArgs
	_args.Req = Req
	var _result GetWhiteListResult
	if err = p.c.Call(ctx, "GetWhiteList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) DeleteWhiteRouter(ctx context.Context, Req *rule.DeleteWhiteRouterReq) (r *rule.DeleteWhiteRouterResp, err error) {
	var _args DeleteWhiteRouterArgs
	_args.Req = Req
	var _result DeleteWhiteRouterResult
	if err = p.c.Call(ctx, "DeleteWhiteRouter", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
