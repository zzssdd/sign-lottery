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

// PrizeAdd .
// @router /prize/add/ [POST]
func PrizeAdd(ctx context.Context, c *app.RequestContext) {
	var err error
	var req lottery.PrizeAddRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	var rpcReq *lottery2.PrizeAddRequest
	err = common.BindRpcOption(req, rpcReq)
	if err != nil {
		Log.Errorln("bind rpc option err:", err)
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp, err := rpc.LotteryClient.PrizeAdd(ctx, rpcReq)
	if err != nil {
		Log.Errorln("prize add err:", err)
	}
	c.JSON(consts.StatusOK, resp)
}

// PrizeDel .
// @router /prize/del/ [DELETE]
func PrizeDel(ctx context.Context, c *app.RequestContext) {
	var err error
	var req lottery.PrizeDelRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	var rpcReq *lottery2.PrizeDelRequest
	err = common.BindRpcOption(req, rpcReq)
	if err != nil {
		Log.Errorln("bind rpc option err:", err)
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp, err := rpc.LotteryClient.PrizeDel(ctx, rpcReq)
	if err != nil {
		Log.Errorln("del prize err:", err)
	}
	c.JSON(consts.StatusOK, resp)
}

// PrizeUpdate .
// @router /prize/update/ [PUT]
func PrizeUpdate(ctx context.Context, c *app.RequestContext) {
	var err error
	var req lottery.PrizeUpdateRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	var rpcReq *lottery2.PrizeUpdateRequest
	err = common.BindRpcOption(req, rpcReq)
	if err != nil {
		Log.Errorln("bind rpc option err:", err)
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp, err := rpc.LotteryClient.PrizeUpdate(ctx, rpcReq)
	if err != nil {
		Log.Errorln("update prize err:", err)
	}
	c.JSON(consts.StatusOK, resp)
}

// GetPrizeByAid .
// @router /prize/aid/ [GET]
func GetPrizeByAid(ctx context.Context, c *app.RequestContext) {
	var err error
	var req lottery.GetPrizeByAidRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	var rpcReq *lottery2.GetPrizeByAidRequest
	err = common.BindRpcOption(req, rpcReq)
	if err != nil {
		Log.Errorln("bind rpc option err:", err)
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp, err := rpc.LotteryClient.GetPrizeByAid(ctx, rpcReq)
	if err != nil {
		Log.Errorln("get prize by aid err:", err)
	}
	c.JSON(consts.StatusOK, resp)
}

// GetPrizeById .
// @router /prize/id/ [GET]
func GetPrizeById(ctx context.Context, c *app.RequestContext) {
	var err error
	var req lottery.GetPrizeByIdRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	var rpcReq *lottery2.GetPrizeByIdRequest
	err = common.BindRpcOption(req, rpcReq)
	if err != nil {
		Log.Errorln("bind rpc option err:", err)
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp, err := rpc.LotteryClient.GetPrizeById(ctx, rpcReq)
	if err != nil {
		Log.Errorln("get prize by id err:", err)
	}
	c.JSON(consts.StatusOK, resp)
}
