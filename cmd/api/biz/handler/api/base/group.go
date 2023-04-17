package base

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"sign-lottery/cmd/api/biz/handler/common"
	"sign-lottery/cmd/api/biz/model/user"
	"sign-lottery/cmd/rpc/base"
	base2 "sign-lottery/kitex_gen/user"
	. "sign-lottery/pkg/log"
)

// CreateGroup .
// @router /group/add/ [POST]
func CreateGroup(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.CreateGroupRequest
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
	req.Owner = value.(int64)
	var rpcReq *base2.CreateGroupRequest
	err = common.BindRpcOption(req, rpcReq)
	if err != nil {
		Log.Errorln("bind rpc option err:", err)
	}
	resp, err := base.BaseClient.CreateGroup(ctx, rpcReq)
	if err != nil {
		Log.Errorln("create group err:", err)
	}
	c.JSON(consts.StatusOK, resp)
}

// JoinGroup .
// @router /group/join/ [POST]
func JoinGroup(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.JoinGroupRequest
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
	var rpcReq *base2.JoinGroupRequest
	err = common.BindRpcOption(req, rpcReq)
	if err != nil {
		Log.Errorln("bind rpc option err:", err)
	}
	resp, err := base.BaseClient.JoinGroup(ctx, rpcReq)
	if err != nil {
		Log.Errorln("join group err:", err)
	}
	c.JSON(consts.StatusOK, resp)
}

// GetGroupById .
// @router /group/id/ [GET]
func GetGroupById(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.GetGroupByIdRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	var rpcReq *base2.GetGroupByIdRequest
	err = common.BindRpcOption(req, rpcReq)
	if err != nil {
		Log.Errorln("bind rpc option err:", err)
	}
	resp, err := base.BaseClient.GetGroupById(ctx, rpcReq)
	if err != nil {
		Log.Errorln("get group by id err:", err)
	}
	c.JSON(consts.StatusOK, resp)
}

// GetAllGroup .
// @router /group/list/ [GET]
func GetAllGroup(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.GetAllGroupRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	var rpcReq *base2.GetAllGroupRequest
	err = common.BindRpcOption(req, rpcReq)
	if err != nil {
		Log.Errorln("bind rpc option err:", err)
	}
	resp, err := base.BaseClient.GetAllGroup(ctx, rpcReq)
	if err != nil {
		Log.Errorln("get all group err:", err)
	}
	c.JSON(consts.StatusOK, resp)
}

// GroupUpdate .
// @router /group/put/ [PUT]
func GroupUpdate(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.GroupUpdateRequest
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
	req.UID = value.(int64)
	var rpcReq *base2.GroupUpdateRequest
	err = common.BindRpcOption(req, rpcReq)
	if err != nil {
		Log.Errorln("bind rpc option err:", err)
	}
	resp, err := base.BaseClient.GroupUpdate(ctx, rpcReq)

	c.JSON(consts.StatusOK, resp)
}

// GroupDel .
// @router /group/del/ [DELETE]
func GroupDel(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.GroupDelRequest
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
	req.UID = value.(int64)
	var rpcReq *base2.GroupDelRequest
	err = common.BindRpcOption(req, rpcReq)
	if err != nil {
		Log.Errorln("bind rpc option err:", err)
	}
	resp, err := base.BaseClient.GroupDel(ctx, rpcReq)

	c.JSON(consts.StatusOK, resp)
}
