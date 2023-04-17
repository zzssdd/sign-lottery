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

// ActivityAdd .
// @router /activity/add/ [POST]
func ActivityAdd(ctx context.Context, c *app.RequestContext) {
	var err error
	var req lottery.ActivityAddRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	value, exists := c.Get("id")
	if !exists {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	id := value.(int64)
	req.UID = &id
	var rpcReq *lottery2.ActivityAddRequest
	err = common.BindRpcOption(req, rpcReq)
	if err != nil {
		Log.Errorln("bind rpc option err:", err)
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp, err := rpc.LotteryClient.ActivityAdd(ctx, rpcReq)
	if err != nil {
		Log.Errorln("add activity err:", err)
	}
	c.JSON(consts.StatusOK, resp)
}

// ActivityDel .
// @router /activity/del/ [DELETE]
func ActivityDel(ctx context.Context, c *app.RequestContext) {
	var err error
	var req lottery.ActivityDelRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	value, exists := c.Get("id")
	if !exists {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	id := value.(int64)
	req.UID = id
	var rpcReq *lottery2.ActivityDelRequest
	err = common.BindRpcOption(req, rpcReq)
	if err != nil {
		Log.Errorln("bind rpc option err:", err)
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp, err := rpc.LotteryClient.ActivityDel(ctx, rpcReq)
	if err != nil {
		Log.Errorln("del activity err:", err)
	}
	c.JSON(consts.StatusOK, resp)
}

// ActivityUpdate .
// @router /activity/update/ [PUT]
func ActivityUpdate(ctx context.Context, c *app.RequestContext) {
	var err error
	var req lottery.ActivityUpdateRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	value, exists := c.Get("id")
	if !exists {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	id := value.(int64)
	req.UID = id
	var rpcReq *lottery2.ActivityUpdateRequest
	err = common.BindRpcOption(req, rpcReq)
	if err != nil {
		Log.Errorln("bind rpc option err:", err)
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp, err := rpc.LotteryClient.ActivityUpdate(ctx, rpcReq)
	if err != nil {
		Log.Errorln("update activity err:", err)
	}
	c.JSON(consts.StatusOK, resp)
}

// GetActivityByGid .
// @router /activity/gid/ [GET]
func GetActivityByGid(ctx context.Context, c *app.RequestContext) {
	var err error
	var req lottery.GetActivityByGidRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	var rpcReq *lottery2.GetActivityByGidRequest
	err = common.BindRpcOption(req, rpcReq)
	if err != nil {
		Log.Errorln("bind rpc option err:", err)
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp, err := rpc.LotteryClient.GetActivityByGid(ctx, rpcReq)
	if err != nil {
		Log.Errorln("get activity by gid err:", err)
	}
	c.JSON(consts.StatusOK, resp)
}

// GetAllActivity .
// @router /activity/list/ [GET]
func GetAllActivity(ctx context.Context, c *app.RequestContext) {
	var err error
	var req lottery.GetAllActivityRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	var rpcReq *lottery2.GetAllActivityRequest
	err = common.BindRpcOption(req, rpcReq)
	if err != nil {
		Log.Errorln("bind rpc option err:", err)
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp, err := rpc.LotteryClient.GetAllActivity(ctx, rpcReq)
	if err != nil {
		Log.Errorln("get all activity err:", err)
	}
	c.JSON(consts.StatusOK, resp)
}

// GetActivityById .
// @router /activity/id/ [GET]
func GetActivityById(ctx context.Context, c *app.RequestContext) {
	var err error
	var req lottery.GetActivityByIdRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	var rpcReq *lottery2.GetActivityByIdRequest
	err = common.BindRpcOption(req, rpcReq)
	if err != nil {
		Log.Errorln("bind rpc option err:", err)
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp, err := rpc.LotteryClient.GetActivityById(ctx, rpcReq)
	if err != nil {
		Log.Errorln("get activity by id err:", err)
	}
	c.JSON(consts.StatusOK, resp)
}
