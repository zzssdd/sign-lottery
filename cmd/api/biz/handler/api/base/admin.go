package base

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"sign-lottery/cmd/api/biz/model/user"
)

// AdminLogin .
// @router /admin/login/ [POST]
func AdminLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.AdminLoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.AdminLoginResponse)

	c.JSON(consts.StatusOK, resp)
}
