package lottery

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"sign-lottery/cmd/api/biz/model/lottery"
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

	resp := new(lottery.BaseResponse)

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

	resp := new(lottery.BaseResponse)

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

	resp := new(lottery.BaseResponse)

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

	resp := new(lottery.ActivitysResponse)

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

	resp := new(lottery.ActivitysResponse)

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

	resp := new(lottery.ActivityResponse)

	c.JSON(consts.StatusOK, resp)
}
