package lottery

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"sign-lottery/cmd/api/biz/model/lottery"
)

// GetUserOrder .
// @router /order/uid/ [GET]
func GetUserOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req lottery.GetUserOrderRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(lottery.OrdersResponse)

	c.JSON(consts.StatusOK, resp)
}

// GetAllOrder .
// @router /order/list/ [GET]
func GetAllOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req lottery.GetAllOrderRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(lottery.OrdersResponse)

	c.JSON(consts.StatusOK, resp)
}
