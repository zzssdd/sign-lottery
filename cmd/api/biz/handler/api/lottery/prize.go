package lottery

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"sign-lottery/cmd/api/biz/model/lottery"
)

// PrizeAdd .
// @router /prize/add/ [POST]
func PrizeAdd(ctx context.Context, c *app.RequestContext) {
	var err error
	var req lottery.PrizeAddRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(lottery.BaseResponse)

	c.JSON(consts.StatusOK, resp)
}

// PrizeDel .
// @router /prize/del/ [DELETE]
func PrizeDel(ctx context.Context, c *app.RequestContext) {
	var err error
	var req lottery.PrizeDelRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(lottery.BaseResponse)

	c.JSON(consts.StatusOK, resp)
}

// PrizeUpdate .
// @router /prize/update/ [PUT]
func PrizeUpdate(ctx context.Context, c *app.RequestContext) {
	var err error
	var req lottery.PrizeUpdateRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(lottery.BaseResponse)

	c.JSON(consts.StatusOK, resp)
}

// GetPrizeByAid .
// @router /prize/aid/ [GET]
func GetPrizeByAid(ctx context.Context, c *app.RequestContext) {
	var err error
	var req lottery.GetPrizeByAidRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(lottery.PrizesResponse)

	c.JSON(consts.StatusOK, resp)
}

// GetPrizeById .
// @router /prize/id/ [GET]
func GetPrizeById(ctx context.Context, c *app.RequestContext) {
	var err error
	var req lottery.GetPrizeByIdRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(lottery.PrizeResponse)

	c.JSON(consts.StatusOK, resp)
}
