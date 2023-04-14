package sign

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"sign-lottery/cmd/api/biz/model/sign"
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

	resp := new(sign.BaseResponse)

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

	resp := new(sign.BaseResponse)

	c.JSON(consts.StatusOK, resp)
}
