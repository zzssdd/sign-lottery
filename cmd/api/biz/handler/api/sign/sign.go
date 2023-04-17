package sign

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"sign-lottery/cmd/api/biz/handler/common"
	"sign-lottery/cmd/api/biz/model/sign"
	rpc "sign-lottery/cmd/rpc/sign"
	"sign-lottery/dao/cache"
	sign2 "sign-lottery/kitex_gen/sign"
	"sign-lottery/pkg/errmsg"
	. "sign-lottery/pkg/log"
	model2 "sign-lottery/rabbitmq/model"
	"sign-lottery/rabbitmq/producer"
	"time"
)

// Sign .
// @router /sign/add/ [POST]
func Sign(ctx context.Context, c *app.RequestContext) {
	var err error
	var req sign.SignRequest
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
	signInfo := &model2.Sign{
		Uid:        id,
		Gid:        req.Gid,
		Latitude:   req.Latitude,
		Longtitude: req.Longtitude,
		Ip:         c.ClientIP(),
	}
	err = producer.NewProcuer().Sign.ProducerSign(signInfo)
	if err != nil {
		Log.Errorln("send email to rabbit err:", err)
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	Cache := cache.NewCache()
	for !Cache.HandlerErr.ExistSignErr(ctx, signInfo.Uid, signInfo.Gid) {
		time.Sleep(time.Second)
	}
	code, err := Cache.HandlerErr.GetSignErr(ctx, signInfo.Uid, signInfo.Gid)
	var resp *sign.BaseResponse
	if err != nil {
		Log.Errorln("get sign code from cache err:", err)
		resp = &sign.BaseResponse{
			Code: errmsg.Error,
			Msg:  errmsg.GetMsg(errmsg.Error),
		}
	} else {
		resp = &sign.BaseResponse{
			Code: int32(code),
			Msg:  errmsg.GetMsg(code),
		}
	}
	c.JSON(consts.StatusOK, resp)
}

// AskLeave .
// @router /sign/leave/ [POST]
func AskLeave(ctx context.Context, c *app.RequestContext) {
	var err error
	var req sign.AskLeaveRequest
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
	var rpcReq *sign2.AskLeaveRequest
	err = common.BindRpcOption(req, rpcReq)
	if err != nil {
		Log.Errorln("bind rpc option err:", err)
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	resp, err := rpc.SignClient.AskLeave(ctx, rpcReq)
	if err != nil {
		Log.Errorln("ask leave err:", err)
	}
	c.JSON(consts.StatusOK, resp)
}
