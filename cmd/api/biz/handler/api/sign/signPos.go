package sign

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"sign-lottery/cmd/api/biz/model/sign"
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

	resp := new(sign.MonthSignResponse)

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

	resp := new(sign.MonthSignsResponse)

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

	resp := new(sign.RecordsResponse)

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

	resp := new(sign.RecordsResponse)

	c.JSON(consts.StatusOK, resp)
}
