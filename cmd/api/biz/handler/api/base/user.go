package base

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"sign-lottery/cmd/api/biz/handler/common"
	"sign-lottery/cmd/api/biz/model/user"
	"sign-lottery/cmd/rpc/base"
	"sign-lottery/dao/cache"
	base2 "sign-lottery/kitex_gen/user"
	"sign-lottery/pkg/errmsg"
	. "sign-lottery/pkg/log"
	"sign-lottery/rabbitmq/producer"
	"sign-lottery/utils"
	"time"
)

// SendEmail .
// @router /user/email/ [POST]
func SendEmail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.EmailRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	err = producer.NewProcuer().Email.ProducerEmail(req.Email)
	if err != nil {
		Log.Errorln("send email to rabbit err:", err)
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	Cache := cache.NewCache()
	for !Cache.HandlerErr.ExistEmailErr(ctx, req.Email) {
		time.Sleep(time.Second)
	}
	code, err := Cache.HandlerErr.GetEmailErr(ctx, req.Email)
	var resp *user.BaseResponse
	if err != nil {
		Log.Errorln("get email code from cache err:", err)
		resp = &user.BaseResponse{
			Code: errmsg.Error,
			Msg:  errmsg.GetMsg(errmsg.Error),
		}
	} else {
		resp = &user.BaseResponse{
			Code: int32(code),
			Msg:  errmsg.GetMsg(code),
		}
	}
	c.JSON(consts.StatusOK, resp)
}

// Registe .
// @router /user/registe/ [POST]
func Registe(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.RegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	req.Password = utils.Crypto(req.Password)
	var rpcReq *base2.RegisterRequest
	err = common.BindRpcOption(rpcReq, req)
	if err != nil {
		Log.Errorln("bind rpc option err:", err)
		c.JSON(consts.StatusBadRequest, err.Error())
	}
	resp, err := base.BaseClient.Registe(ctx, rpcReq)
	if err != nil {
		Log.Errorln("register err:", err)
	}
	c.JSON(consts.StatusOK, resp)
}

// Login .
// @router /user/login/ [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.LoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	req.Password = utils.Crypto(req.Password)
	var rpcReq *base2.LoginRequest
	common.BindRpcOption(req, rpcReq)
	resp, err := base.BaseClient.Login(ctx, rpcReq)
	if err != nil {
		Log.Errorln("login err:", err)
	}
	c.JSON(consts.StatusOK, resp)
}

// GetUserById .
// @router /user/id/ [GET]
func GetUserById(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.GetUserByIdRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	var rpcReq *base2.GetUserByIdRequest
	err = common.BindRpcOption(req, rpcReq)
	if err != nil {
		Log.Errorln("bind rpc option err:", err)
	}
	resp, err := base.BaseClient.GetUserById(ctx, rpcReq)
	if err != nil {
		Log.Errorln("get user by id err:", err)
	}
	c.JSON(consts.StatusOK, resp)
}

// GetUserByGid .
// @router /user/gid/ [GET]
func GetUserByGid(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.GetUserByGidRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	var rpcReq *base2.GetUserByGidRequest

	err = common.BindRpcOption(req, rpcReq)
	if err != nil {
		Log.Errorln("bind rpc option err:", err)
	}
	resp, err := base.BaseClient.GetUserByGid(ctx, rpcReq)
	if err != nil {
		Log.Errorln("get user by gid err:", err)
	}

	c.JSON(consts.StatusOK, resp)
}

// ChangeUserAvater .
// @router /user/avater/ [PUT]
func ChangeUserAvater(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.ChangeUserAvaterRequest
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
	req.ID = &id
	var rpcReq *base2.ChangeUserAvaterRequest
	err = common.BindRpcOption(req, rpcReq)
	if err != nil {
		Log.Errorln("bind rpc option err:", err)
	}
	resp, err := base.BaseClient.ChangeUserAvater(ctx, rpcReq)
	if err != nil {
		Log.Errorln("change user avater err:", err)
	}
	c.JSON(consts.StatusOK, resp)
}

// ChangeUserPassword .
// @router /user/password/ [PUT]
func ChangeUserPassword(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.ChangePasswordRequest
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
	req.ID = &id
	var rpcReq *base2.ChangePasswordRequest
	err = common.BindRpcOption(req, rpcReq)
	if err != nil {
		Log.Errorln("bind rpc option err:", err)
	}
	resp, err := base.BaseClient.ChangeUserPassword(ctx, rpcReq)
	if err != nil {
		Log.Errorln("change user password err:", err)
	}
	c.JSON(consts.StatusOK, resp)
}

// ChangeUserAddress .
// @router /user/address/ [PUT]
func ChangeUserAddress(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.ChangeAddressRequest
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
	req.ID = &id
	var rpcReq *base2.ChangeAddressRequest
	err = common.BindRpcOption(req, rpcReq)
	if err != nil {
		Log.Errorln("bind rpc option err:", err)
	}
	resp, err := base.BaseClient.ChangeUserAddress(ctx, rpcReq)
	if err != nil {
		Log.Errorln("change user address err:", err)
	}
	c.JSON(consts.StatusOK, resp)
}

// UserDel .
// @router /admin/user/ [DELETE]
func UserDel(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserDelRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	var rpcReq *base2.UserDelRequest
	err = common.BindRpcOption(req, rpcReq)
	if err != nil {
		Log.Errorln("bind rpc option err:", err)
	}
	resp, err := base.BaseClient.UserDel(ctx, rpcReq)
	if err != nil {
		Log.Errorln("user del err:", err)
	}
	c.JSON(consts.StatusOK, resp)
}

// GetAllUser .
// @router /admin/userlist/ [GET]
func GetAllUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.GetAllUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	var rpcReq *base2.GetAllUserRequest
	err = common.BindRpcOption(req, rpcReq)
	if err != nil {
		Log.Errorln("bind rpc option err:", err)
	}

	resp, err := base.BaseClient.GetAllUser(ctx, rpcReq)
	if err != nil {
		Log.Errorln("get all user err:", err)
	}
	c.JSON(consts.StatusOK, resp)
}
