package lottery

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"sign-lottery/cmd/api/biz/handler/common"
	"sign-lottery/cmd/api/biz/model/lottery"
	rpc "sign-lottery/cmd/rpc/lottery"
	lottery2 "sign-lottery/kitex_gen/lottery"
	. "sign-lottery/pkg/log"
)

// GetUserOrder .
// @router /order/uid/ [GET]
func GetUserOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req lottery.GetUserOrderRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	var rpcReq *lottery2.GetUserOrderRequest
	err = common.BindRpcOption(req, rpcReq)
	if err != nil {
		Log.Errorln("bind rpc option err:", err)
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp, err := rpc.LotteryClient.GetUserOrder(ctx, rpcReq)
	if err != nil {
		Log.Errorln("get user order err:", err)
	}
	c.JSON(consts.StatusOK, resp)
}

// GetAllOrder .
// @router /order/list/ [GET]
func GetAllOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req lottery.GetAllOrderRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	var rpcReq *lottery2.GetAllOrderRequest
	err = common.BindRpcOption(req, rpcReq)
	if err != nil {
		Log.Errorln("bind rpc option err:", err)
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp, err := rpc.LotteryClient.GetAllOrder(ctx, rpcReq)
	if err != nil {
		Log.Errorln("get all order err:", err)
	}
	c.JSON(consts.StatusOK, resp)
}
