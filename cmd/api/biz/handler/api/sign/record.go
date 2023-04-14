package sign

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"sign-lottery/cmd/api/biz/model/sign"
)

// SignPosAdd .
// @router /sign/pos/ [POST]
func SignPosAdd(ctx context.Context, c *app.RequestContext) {
	var err error
	var req sign.SignPosAddRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(sign.BaseResponse)

	c.JSON(consts.StatusOK, resp)
}

// SignPosDel .
// @router /sign/pos/ [DELETE]
func SignPosDel(ctx context.Context, c *app.RequestContext) {
	var err error
	var req sign.SignPosDelRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(sign.BaseResponse)

	c.JSON(consts.StatusOK, resp)
}

// GetSignPos .
// @router /sign/pos/ [GET]
func GetSignPos(ctx context.Context, c *app.RequestContext) {
	var err error
	var req sign.GetSignPosRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(sign.GetSignPosResponse)

	c.JSON(consts.StatusOK, resp)
}
