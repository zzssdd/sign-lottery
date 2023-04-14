package base

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"sign-lottery/cmd/api/biz/model/user"
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

	resp := new(user.BaseResponse)

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

	resp := new(user.BaseResponse)

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

	resp := new(user.GroupResponse)

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

	resp := new(user.GroupsResponse)

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

	resp := new(user.BaseResponse)

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

	resp := new(user.BaseResponse)

	c.JSON(consts.StatusOK, resp)
}
