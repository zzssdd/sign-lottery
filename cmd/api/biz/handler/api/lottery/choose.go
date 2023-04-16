package lottery

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"sign-lottery/cmd/api/biz/model/lottery"
	"sign-lottery/dao/cache"
	"sign-lottery/dao/db"
	"sign-lottery/pkg/errmsg"
	. "sign-lottery/pkg/log"
	model2 "sign-lottery/rabbitmq/model"
	"sign-lottery/rabbitmq/producer"
	"time"
)

// Choose .
// @router /prize/choose/ [GET]
func Choose(ctx context.Context, c *app.RequestContext) {
	var err error
	var req lottery.ChooseRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	Cache := cache.NewCache()
	Db := db.NewDao()
	var resp *lottery.ChooseResponse
	if !Cache.Activity.ExistActivityInfo(ctx, req.Aid) {
		activity, err := Db.Activity.GetActivityById(ctx, req.Aid)
		if err != nil {
			Log.Errorln("get activity from db err:", err)
			c.String(consts.StatusBadRequest, err.Error())
			return
		}
		err = Cache.Activity.StoreActivityInfo(ctx, req.Aid, activity)
		if err != nil {
			Log.Errorln("store activity into cache err:", err)
			c.String(consts.StatusBadRequest, err.Error())
			return
		}
	}
	now := time.Now()
	is_ok := Cache.Activity.CheckActivity(ctx, req.Aid, now)
	if !is_ok {
		resp = &lottery.ChooseResponse{
			Resp: &lottery.BaseResponse{
				Code: errmsg.PrizeIsNull,
				Msg:  errmsg.GetMsg(errmsg.PrizeIsNull),
			},
		}
	}
	choose := &model2.Choose{
		Uid: *req.UID,
		Aid: req.Aid,
	}
	err = producer.NewProcuer().Choose.ProducerChoose(choose)
	if err != nil {
		resp = &lottery.ChooseResponse{
			Resp: &lottery.BaseResponse{
				Code: errmsg.Error,
				Msg:  errmsg.GetMsg(errmsg.Error),
			},
		}
	}
	for !Cache.HandlerErr.ExistChooseErr(ctx, *req.UID, req.Aid) {
		time.Sleep(time.Second)
	}
	prizeName, err := Cache.HandlerErr.GetChooseErr(ctx, *req.UID, req.Aid)
	if err != nil {
		if err != nil {
			Log.Errorln("get code from cache err:", err)
			c.String(consts.StatusBadRequest, err.Error())
			return
		}
	}
	resp = &lottery.ChooseResponse{
		Resp: &lottery.BaseResponse{
			Code: errmsg.Success,
			Msg:  errmsg.GetMsg(errmsg.Success),
		},
		Name: prizeName,
	}
	c.JSON(consts.StatusOK, resp)
}
