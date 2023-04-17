package sign

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"sign-lottery/cmd/api/biz/handler/common"
	"sign-lottery/cmd/api/biz/model/sign"
	rpc "sign-lottery/cmd/rpc/sign"
	sign2 "sign-lottery/kitex_gen/sign"
	. "sign-lottery/pkg/log"
)

// GetMonthSign .
// @router /sign/month/ [GET]
func GetMonthSign(ctx context.Context, c *app.RequestContext) {
	var err error
	var req sign.GetMonthSignRequest
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
	var rpcReq *sign2.GetMonthSignRequest
	err = common.BindRpcOption(req, rpcReq)
	if err != nil {
		Log.Errorln("bind rpc option err:", err)
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp, err := rpc.SignClient.GetMonthSign(ctx, rpcReq)
	if err != nil {
		Log.Errorln("get month sign err:", err)
	}
	c.JSON(consts.StatusOK, resp)
}

// GetMonthSignByGid .
// @router /sign/gmonth [GET]
func GetMonthSignByGid(ctx context.Context, c *app.RequestContext) {
	var err error
	var req sign.GetMonthSignsByGid
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	var rpcReq *sign2.GetMonthSignsByGid
	err = common.BindRpcOption(req, rpcReq)
	if err != nil {
		Log.Errorln("bind rpc option err:", err)
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp, err := rpc.SignClient.GetMonthSignByGid(ctx, rpcReq)
	if err != nil {
		Log.Errorln("get month sign by gid err:", err)
	}

	c.JSON(consts.StatusOK, resp)
}

// GetAllRecord .
// @router /sign/recordlist/ [GET]
func GetAllRecord(ctx context.Context, c *app.RequestContext) {
	var err error
	var req sign.GetAllRecordRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	var rpcReq *sign2.GetAllRecordRequest
	err = common.BindRpcOption(req, rpcReq)
	if err != nil {
		Log.Errorln("bind rpc option err:", err)
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp, err := rpc.SignClient.GetAllRecord(ctx, rpcReq)
	if err != nil {
		Log.Errorln("get all record err:", err)
	}
	c.JSON(consts.StatusOK, resp)
}

// GetUserRecord .
// @router /sign/record/ [GET]
func GetUserRecord(ctx context.Context, c *app.RequestContext) {
	var err error
	var req sign.GetUserRecordRequest
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
	var rpcReq *sign2.GetUserRecordRequest

	err = common.BindRpcOption(req, rpcReq)
	if err != nil {
		Log.Errorln("bind rpc option err:", err)
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp, err := rpc.SignClient.GetUserRecord(ctx, rpcReq)
	if err != nil {
		Log.Errorln("get user record err:", err)
	}
	c.JSON(consts.StatusOK, resp)
}
