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

// SignPosAdd .
// @router /sign/pos/ [POST]
func SignPosAdd(ctx context.Context, c *app.RequestContext) {
	var err error
	var req sign.SignPosAddRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	var rpcReq *sign2.SignPosAddRequest
	err = common.BindRpcOption(req, rpcReq)
	if err != nil {
		Log.Errorln("bind rpc option err:", err)
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp, err := rpc.SignClient.SignPosAdd(ctx, rpcReq)
	if err != nil {
		Log.Errorln("sign add pos err:", err)
	}

	c.JSON(consts.StatusOK, resp)
}

// SignPosDel .
// @router /sign/pos/ [DELETE]
func SignPosDel(ctx context.Context, c *app.RequestContext) {
	var err error
	var req sign.SignPosDelRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	var rpcReq *sign2.SignPosDelRequest
	err = common.BindRpcOption(req, rpcReq)
	if err != nil {
		Log.Errorln("bind rpc option err:", err)
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp, err := rpc.SignClient.SignPosDel(ctx, rpcReq)
	if err != nil {
		Log.Errorln("sign del pos err:", err)
	}
	c.JSON(consts.StatusOK, resp)
}

// GetSignPos .
// @router /sign/pos/ [GET]
func GetSignPos(ctx context.Context, c *app.RequestContext) {
	var err error
	var req sign.GetSignPosRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	var rpcReq *sign2.GetSignPosRequest
	err = common.BindRpcOption(req, rpcReq)
	if err != nil {
		Log.Errorln("bind rpc option err:", err)
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp, err := rpc.SignClient.GetSignPos(ctx, rpcReq)
	if err != nil {
		Log.Errorln("get sign pos err:", err)
	}
	c.JSON(consts.StatusOK, resp)
}
