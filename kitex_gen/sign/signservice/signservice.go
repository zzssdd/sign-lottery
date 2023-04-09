// Code generated by Kitex v0.5.1. DO NOT EDIT.

package signservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	sign "sign-lottery/kitex_gen/sign"
)

func serviceInfo() *kitex.ServiceInfo {
	return signServiceServiceInfo
}

var signServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "SignService"
	handlerType := (*sign.SignService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Sign":              kitex.NewMethodInfo(signHandler, newSignServiceSignArgs, newSignServiceSignResult, false),
		"AskLeave":          kitex.NewMethodInfo(askLeaveHandler, newSignServiceAskLeaveArgs, newSignServiceAskLeaveResult, false),
		"GetMonthSign":      kitex.NewMethodInfo(getMonthSignHandler, newSignServiceGetMonthSignArgs, newSignServiceGetMonthSignResult, false),
		"GetMonthSignByGid": kitex.NewMethodInfo(getMonthSignByGidHandler, newSignServiceGetMonthSignByGidArgs, newSignServiceGetMonthSignByGidResult, false),
		"GetAllRecord":      kitex.NewMethodInfo(getAllRecordHandler, newSignServiceGetAllRecordArgs, newSignServiceGetAllRecordResult, false),
		"GetUserRecord":     kitex.NewMethodInfo(getUserRecordHandler, newSignServiceGetUserRecordArgs, newSignServiceGetUserRecordResult, false),
		"SignPosAdd":        kitex.NewMethodInfo(signPosAddHandler, newSignServiceSignPosAddArgs, newSignServiceSignPosAddResult, false),
		"SignPosDel":        kitex.NewMethodInfo(signPosDelHandler, newSignServiceSignPosDelArgs, newSignServiceSignPosDelResult, false),
		"GetSignPos":        kitex.NewMethodInfo(getSignPosHandler, newSignServiceGetSignPosArgs, newSignServiceGetSignPosResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "sign",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.5.1",
		Extra:           extra,
	}
	return svcInfo
}

func signHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*sign.SignServiceSignArgs)
	realResult := result.(*sign.SignServiceSignResult)
	success, err := handler.(sign.SignService).Sign(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSignServiceSignArgs() interface{} {
	return sign.NewSignServiceSignArgs()
}

func newSignServiceSignResult() interface{} {
	return sign.NewSignServiceSignResult()
}

func askLeaveHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*sign.SignServiceAskLeaveArgs)
	realResult := result.(*sign.SignServiceAskLeaveResult)
	success, err := handler.(sign.SignService).AskLeave(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSignServiceAskLeaveArgs() interface{} {
	return sign.NewSignServiceAskLeaveArgs()
}

func newSignServiceAskLeaveResult() interface{} {
	return sign.NewSignServiceAskLeaveResult()
}

func getMonthSignHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*sign.SignServiceGetMonthSignArgs)
	realResult := result.(*sign.SignServiceGetMonthSignResult)
	success, err := handler.(sign.SignService).GetMonthSign(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSignServiceGetMonthSignArgs() interface{} {
	return sign.NewSignServiceGetMonthSignArgs()
}

func newSignServiceGetMonthSignResult() interface{} {
	return sign.NewSignServiceGetMonthSignResult()
}

func getMonthSignByGidHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*sign.SignServiceGetMonthSignByGidArgs)
	realResult := result.(*sign.SignServiceGetMonthSignByGidResult)
	success, err := handler.(sign.SignService).GetMonthSignByGid(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSignServiceGetMonthSignByGidArgs() interface{} {
	return sign.NewSignServiceGetMonthSignByGidArgs()
}

func newSignServiceGetMonthSignByGidResult() interface{} {
	return sign.NewSignServiceGetMonthSignByGidResult()
}

func getAllRecordHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*sign.SignServiceGetAllRecordArgs)
	realResult := result.(*sign.SignServiceGetAllRecordResult)
	success, err := handler.(sign.SignService).GetAllRecord(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSignServiceGetAllRecordArgs() interface{} {
	return sign.NewSignServiceGetAllRecordArgs()
}

func newSignServiceGetAllRecordResult() interface{} {
	return sign.NewSignServiceGetAllRecordResult()
}

func getUserRecordHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*sign.SignServiceGetUserRecordArgs)
	realResult := result.(*sign.SignServiceGetUserRecordResult)
	success, err := handler.(sign.SignService).GetUserRecord(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSignServiceGetUserRecordArgs() interface{} {
	return sign.NewSignServiceGetUserRecordArgs()
}

func newSignServiceGetUserRecordResult() interface{} {
	return sign.NewSignServiceGetUserRecordResult()
}

func signPosAddHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*sign.SignServiceSignPosAddArgs)
	realResult := result.(*sign.SignServiceSignPosAddResult)
	success, err := handler.(sign.SignService).SignPosAdd(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSignServiceSignPosAddArgs() interface{} {
	return sign.NewSignServiceSignPosAddArgs()
}

func newSignServiceSignPosAddResult() interface{} {
	return sign.NewSignServiceSignPosAddResult()
}

func signPosDelHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*sign.SignServiceSignPosDelArgs)
	realResult := result.(*sign.SignServiceSignPosDelResult)
	success, err := handler.(sign.SignService).SignPosDel(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSignServiceSignPosDelArgs() interface{} {
	return sign.NewSignServiceSignPosDelArgs()
}

func newSignServiceSignPosDelResult() interface{} {
	return sign.NewSignServiceSignPosDelResult()
}

func getSignPosHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*sign.SignServiceGetSignPosArgs)
	realResult := result.(*sign.SignServiceGetSignPosResult)
	success, err := handler.(sign.SignService).GetSignPos(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSignServiceGetSignPosArgs() interface{} {
	return sign.NewSignServiceGetSignPosArgs()
}

func newSignServiceGetSignPosResult() interface{} {
	return sign.NewSignServiceGetSignPosResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Sign(ctx context.Context, req *sign.SignRequest) (r *sign.BaseResponse, err error) {
	var _args sign.SignServiceSignArgs
	_args.Req = req
	var _result sign.SignServiceSignResult
	if err = p.c.Call(ctx, "Sign", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) AskLeave(ctx context.Context, req *sign.AskLeaveRequest) (r *sign.BaseResponse, err error) {
	var _args sign.SignServiceAskLeaveArgs
	_args.Req = req
	var _result sign.SignServiceAskLeaveResult
	if err = p.c.Call(ctx, "AskLeave", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetMonthSign(ctx context.Context, req *sign.GetMonthSignRequest) (r *sign.MonthSignResponse, err error) {
	var _args sign.SignServiceGetMonthSignArgs
	_args.Req = req
	var _result sign.SignServiceGetMonthSignResult
	if err = p.c.Call(ctx, "GetMonthSign", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetMonthSignByGid(ctx context.Context, req *sign.GetMonthSignsByGid) (r *sign.MonthSignsResponse, err error) {
	var _args sign.SignServiceGetMonthSignByGidArgs
	_args.Req = req
	var _result sign.SignServiceGetMonthSignByGidResult
	if err = p.c.Call(ctx, "GetMonthSignByGid", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetAllRecord(ctx context.Context, req *sign.GetAllRecordRequest) (r *sign.RecordsResponse, err error) {
	var _args sign.SignServiceGetAllRecordArgs
	_args.Req = req
	var _result sign.SignServiceGetAllRecordResult
	if err = p.c.Call(ctx, "GetAllRecord", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserRecord(ctx context.Context, req *sign.GetUserRecordRequest) (r *sign.RecordsResponse, err error) {
	var _args sign.SignServiceGetUserRecordArgs
	_args.Req = req
	var _result sign.SignServiceGetUserRecordResult
	if err = p.c.Call(ctx, "GetUserRecord", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SignPosAdd(ctx context.Context, req *sign.SignPosAddRequest) (r *sign.BaseResponse, err error) {
	var _args sign.SignServiceSignPosAddArgs
	_args.Req = req
	var _result sign.SignServiceSignPosAddResult
	if err = p.c.Call(ctx, "SignPosAdd", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SignPosDel(ctx context.Context, req *sign.SignPosDelRequest) (r *sign.BaseResponse, err error) {
	var _args sign.SignServiceSignPosDelArgs
	_args.Req = req
	var _result sign.SignServiceSignPosDelResult
	if err = p.c.Call(ctx, "SignPosDel", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetSignPos(ctx context.Context, req *sign.GetSignPosRequest) (r *sign.GetSignPosResponse, err error) {
	var _args sign.SignServiceGetSignPosArgs
	_args.Req = req
	var _result sign.SignServiceGetSignPosResult
	if err = p.c.Call(ctx, "GetSignPos", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
