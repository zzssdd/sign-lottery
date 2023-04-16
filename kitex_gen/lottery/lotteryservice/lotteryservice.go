// Code generated by Kitex v0.5.1. DO NOT EDIT.

package lotteryservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	lottery "sign-lottery/kitex_gen/lottery"
)

func serviceInfo() *kitex.ServiceInfo {
	return lotteryServiceServiceInfo
}

var lotteryServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "LotteryService"
	handlerType := (*lottery.LotteryService)(nil)
	methods := map[string]kitex.MethodInfo{
		"ActivityAdd":      kitex.NewMethodInfo(activityAddHandler, newLotteryServiceActivityAddArgs, newLotteryServiceActivityAddResult, false),
		"ActivityDel":      kitex.NewMethodInfo(activityDelHandler, newLotteryServiceActivityDelArgs, newLotteryServiceActivityDelResult, false),
		"ActivityUpdate":   kitex.NewMethodInfo(activityUpdateHandler, newLotteryServiceActivityUpdateArgs, newLotteryServiceActivityUpdateResult, false),
		"GetActivityByGid": kitex.NewMethodInfo(getActivityByGidHandler, newLotteryServiceGetActivityByGidArgs, newLotteryServiceGetActivityByGidResult, false),
		"GetAllActivity":   kitex.NewMethodInfo(getAllActivityHandler, newLotteryServiceGetAllActivityArgs, newLotteryServiceGetAllActivityResult, false),
		"GetActivityById":  kitex.NewMethodInfo(getActivityByIdHandler, newLotteryServiceGetActivityByIdArgs, newLotteryServiceGetActivityByIdResult, false),
		"PrizeAdd":         kitex.NewMethodInfo(prizeAddHandler, newLotteryServicePrizeAddArgs, newLotteryServicePrizeAddResult, false),
		"PrizeDel":         kitex.NewMethodInfo(prizeDelHandler, newLotteryServicePrizeDelArgs, newLotteryServicePrizeDelResult, false),
		"PrizeUpdate":      kitex.NewMethodInfo(prizeUpdateHandler, newLotteryServicePrizeUpdateArgs, newLotteryServicePrizeUpdateResult, false),
		"GetPrizeByAid":    kitex.NewMethodInfo(getPrizeByAidHandler, newLotteryServiceGetPrizeByAidArgs, newLotteryServiceGetPrizeByAidResult, false),
		"GetPrizeById":     kitex.NewMethodInfo(getPrizeByIdHandler, newLotteryServiceGetPrizeByIdArgs, newLotteryServiceGetPrizeByIdResult, false),
		"Choose":           kitex.NewMethodInfo(chooseHandler, newLotteryServiceChooseArgs, newLotteryServiceChooseResult, false),
		"GetUserOrder":     kitex.NewMethodInfo(getUserOrderHandler, newLotteryServiceGetUserOrderArgs, newLotteryServiceGetUserOrderResult, false),
		"GetAllOrder":      kitex.NewMethodInfo(getAllOrderHandler, newLotteryServiceGetAllOrderArgs, newLotteryServiceGetAllOrderResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "lottery",
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

func activityAddHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*lottery.LotteryServiceActivityAddArgs)
	realResult := result.(*lottery.LotteryServiceActivityAddResult)
	success, err := handler.(lottery.LotteryService).ActivityAdd(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newLotteryServiceActivityAddArgs() interface{} {
	return lottery.NewLotteryServiceActivityAddArgs()
}

func newLotteryServiceActivityAddResult() interface{} {
	return lottery.NewLotteryServiceActivityAddResult()
}

func activityDelHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*lottery.LotteryServiceActivityDelArgs)
	realResult := result.(*lottery.LotteryServiceActivityDelResult)
	success, err := handler.(lottery.LotteryService).ActivityDel(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newLotteryServiceActivityDelArgs() interface{} {
	return lottery.NewLotteryServiceActivityDelArgs()
}

func newLotteryServiceActivityDelResult() interface{} {
	return lottery.NewLotteryServiceActivityDelResult()
}

func activityUpdateHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*lottery.LotteryServiceActivityUpdateArgs)
	realResult := result.(*lottery.LotteryServiceActivityUpdateResult)
	success, err := handler.(lottery.LotteryService).ActivityUpdate(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newLotteryServiceActivityUpdateArgs() interface{} {
	return lottery.NewLotteryServiceActivityUpdateArgs()
}

func newLotteryServiceActivityUpdateResult() interface{} {
	return lottery.NewLotteryServiceActivityUpdateResult()
}

func getActivityByGidHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*lottery.LotteryServiceGetActivityByGidArgs)
	realResult := result.(*lottery.LotteryServiceGetActivityByGidResult)
	success, err := handler.(lottery.LotteryService).GetActivityByGid(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newLotteryServiceGetActivityByGidArgs() interface{} {
	return lottery.NewLotteryServiceGetActivityByGidArgs()
}

func newLotteryServiceGetActivityByGidResult() interface{} {
	return lottery.NewLotteryServiceGetActivityByGidResult()
}

func getAllActivityHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*lottery.LotteryServiceGetAllActivityArgs)
	realResult := result.(*lottery.LotteryServiceGetAllActivityResult)
	success, err := handler.(lottery.LotteryService).GetAllActivity(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newLotteryServiceGetAllActivityArgs() interface{} {
	return lottery.NewLotteryServiceGetAllActivityArgs()
}

func newLotteryServiceGetAllActivityResult() interface{} {
	return lottery.NewLotteryServiceGetAllActivityResult()
}

func getActivityByIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*lottery.LotteryServiceGetActivityByIdArgs)
	realResult := result.(*lottery.LotteryServiceGetActivityByIdResult)
	success, err := handler.(lottery.LotteryService).GetActivityById(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newLotteryServiceGetActivityByIdArgs() interface{} {
	return lottery.NewLotteryServiceGetActivityByIdArgs()
}

func newLotteryServiceGetActivityByIdResult() interface{} {
	return lottery.NewLotteryServiceGetActivityByIdResult()
}

func prizeAddHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*lottery.LotteryServicePrizeAddArgs)
	realResult := result.(*lottery.LotteryServicePrizeAddResult)
	success, err := handler.(lottery.LotteryService).PrizeAdd(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newLotteryServicePrizeAddArgs() interface{} {
	return lottery.NewLotteryServicePrizeAddArgs()
}

func newLotteryServicePrizeAddResult() interface{} {
	return lottery.NewLotteryServicePrizeAddResult()
}

func prizeDelHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*lottery.LotteryServicePrizeDelArgs)
	realResult := result.(*lottery.LotteryServicePrizeDelResult)
	success, err := handler.(lottery.LotteryService).PrizeDel(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newLotteryServicePrizeDelArgs() interface{} {
	return lottery.NewLotteryServicePrizeDelArgs()
}

func newLotteryServicePrizeDelResult() interface{} {
	return lottery.NewLotteryServicePrizeDelResult()
}

func prizeUpdateHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*lottery.LotteryServicePrizeUpdateArgs)
	realResult := result.(*lottery.LotteryServicePrizeUpdateResult)
	success, err := handler.(lottery.LotteryService).PrizeUpdate(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newLotteryServicePrizeUpdateArgs() interface{} {
	return lottery.NewLotteryServicePrizeUpdateArgs()
}

func newLotteryServicePrizeUpdateResult() interface{} {
	return lottery.NewLotteryServicePrizeUpdateResult()
}

func getPrizeByAidHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*lottery.LotteryServiceGetPrizeByAidArgs)
	realResult := result.(*lottery.LotteryServiceGetPrizeByAidResult)
	success, err := handler.(lottery.LotteryService).GetPrizeByAid(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newLotteryServiceGetPrizeByAidArgs() interface{} {
	return lottery.NewLotteryServiceGetPrizeByAidArgs()
}

func newLotteryServiceGetPrizeByAidResult() interface{} {
	return lottery.NewLotteryServiceGetPrizeByAidResult()
}

func getPrizeByIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*lottery.LotteryServiceGetPrizeByIdArgs)
	realResult := result.(*lottery.LotteryServiceGetPrizeByIdResult)
	success, err := handler.(lottery.LotteryService).GetPrizeById(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newLotteryServiceGetPrizeByIdArgs() interface{} {
	return lottery.NewLotteryServiceGetPrizeByIdArgs()
}

func newLotteryServiceGetPrizeByIdResult() interface{} {
	return lottery.NewLotteryServiceGetPrizeByIdResult()
}

func chooseHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {

	realResult := result.(*lottery.LotteryServiceChooseResult)
	success, err := handler.(lottery.LotteryService).Choose(ctx)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newLotteryServiceChooseArgs() interface{} {
	return lottery.NewLotteryServiceChooseArgs()
}

func newLotteryServiceChooseResult() interface{} {
	return lottery.NewLotteryServiceChooseResult()
}

func getUserOrderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*lottery.LotteryServiceGetUserOrderArgs)
	realResult := result.(*lottery.LotteryServiceGetUserOrderResult)
	success, err := handler.(lottery.LotteryService).GetUserOrder(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newLotteryServiceGetUserOrderArgs() interface{} {
	return lottery.NewLotteryServiceGetUserOrderArgs()
}

func newLotteryServiceGetUserOrderResult() interface{} {
	return lottery.NewLotteryServiceGetUserOrderResult()
}

func getAllOrderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*lottery.LotteryServiceGetAllOrderArgs)
	realResult := result.(*lottery.LotteryServiceGetAllOrderResult)
	success, err := handler.(lottery.LotteryService).GetAllOrder(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newLotteryServiceGetAllOrderArgs() interface{} {
	return lottery.NewLotteryServiceGetAllOrderArgs()
}

func newLotteryServiceGetAllOrderResult() interface{} {
	return lottery.NewLotteryServiceGetAllOrderResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) ActivityAdd(ctx context.Context, req *lottery.ActivityAddRequest) (r *lottery.BaseResponse, err error) {
	var _args lottery.LotteryServiceActivityAddArgs
	_args.Req = req
	var _result lottery.LotteryServiceActivityAddResult
	if err = p.c.Call(ctx, "ActivityAdd", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ActivityDel(ctx context.Context, req *lottery.ActivityDelRequest) (r *lottery.BaseResponse, err error) {
	var _args lottery.LotteryServiceActivityDelArgs
	_args.Req = req
	var _result lottery.LotteryServiceActivityDelResult
	if err = p.c.Call(ctx, "ActivityDel", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ActivityUpdate(ctx context.Context, req *lottery.ActivityUpdateRequest) (r *lottery.BaseResponse, err error) {
	var _args lottery.LotteryServiceActivityUpdateArgs
	_args.Req = req
	var _result lottery.LotteryServiceActivityUpdateResult
	if err = p.c.Call(ctx, "ActivityUpdate", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetActivityByGid(ctx context.Context, req *lottery.GetActivityByGidRequest) (r *lottery.ActivitysResponse, err error) {
	var _args lottery.LotteryServiceGetActivityByGidArgs
	_args.Req = req
	var _result lottery.LotteryServiceGetActivityByGidResult
	if err = p.c.Call(ctx, "GetActivityByGid", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetAllActivity(ctx context.Context, req *lottery.GetAllActivityRequest) (r *lottery.ActivitysResponse, err error) {
	var _args lottery.LotteryServiceGetAllActivityArgs
	_args.Req = req
	var _result lottery.LotteryServiceGetAllActivityResult
	if err = p.c.Call(ctx, "GetAllActivity", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetActivityById(ctx context.Context, req *lottery.GetActivityByIdRequest) (r *lottery.ActivityResponse, err error) {
	var _args lottery.LotteryServiceGetActivityByIdArgs
	_args.Req = req
	var _result lottery.LotteryServiceGetActivityByIdResult
	if err = p.c.Call(ctx, "GetActivityById", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PrizeAdd(ctx context.Context, req *lottery.PrizeAddRequest) (r *lottery.BaseResponse, err error) {
	var _args lottery.LotteryServicePrizeAddArgs
	_args.Req = req
	var _result lottery.LotteryServicePrizeAddResult
	if err = p.c.Call(ctx, "PrizeAdd", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PrizeDel(ctx context.Context, req *lottery.PrizeDelRequest) (r *lottery.BaseResponse, err error) {
	var _args lottery.LotteryServicePrizeDelArgs
	_args.Req = req
	var _result lottery.LotteryServicePrizeDelResult
	if err = p.c.Call(ctx, "PrizeDel", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PrizeUpdate(ctx context.Context, req *lottery.PrizeUpdateRequest) (r *lottery.BaseResponse, err error) {
	var _args lottery.LotteryServicePrizeUpdateArgs
	_args.Req = req
	var _result lottery.LotteryServicePrizeUpdateResult
	if err = p.c.Call(ctx, "PrizeUpdate", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetPrizeByAid(ctx context.Context, req *lottery.GetPrizeByAidRequest) (r *lottery.PrizesResponse, err error) {
	var _args lottery.LotteryServiceGetPrizeByAidArgs
	_args.Req = req
	var _result lottery.LotteryServiceGetPrizeByAidResult
	if err = p.c.Call(ctx, "GetPrizeByAid", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetPrizeById(ctx context.Context, req *lottery.GetPrizeByIdRequest) (r *lottery.PrizeResponse, err error) {
	var _args lottery.LotteryServiceGetPrizeByIdArgs
	_args.Req = req
	var _result lottery.LotteryServiceGetPrizeByIdResult
	if err = p.c.Call(ctx, "GetPrizeById", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Choose(ctx context.Context) (r *lottery.ChooseResponse, err error) {
	var _args lottery.LotteryServiceChooseArgs
	var _result lottery.LotteryServiceChooseResult
	if err = p.c.Call(ctx, "Choose", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserOrder(ctx context.Context, req *lottery.GetUserOrderRequest) (r *lottery.OrdersResponse, err error) {
	var _args lottery.LotteryServiceGetUserOrderArgs
	_args.Req = req
	var _result lottery.LotteryServiceGetUserOrderResult
	if err = p.c.Call(ctx, "GetUserOrder", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetAllOrder(ctx context.Context, req *lottery.GetAllOrderRequest) (r *lottery.OrdersResponse, err error) {
	var _args lottery.LotteryServiceGetAllOrderArgs
	_args.Req = req
	var _result lottery.LotteryServiceGetAllOrderResult
	if err = p.c.Call(ctx, "GetAllOrder", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
