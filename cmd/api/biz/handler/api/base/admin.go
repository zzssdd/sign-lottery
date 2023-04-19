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
	"sign-lottery/utils"
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
	req.Password = utils.Crypto(req.Password)
	var rpcReq *base2.AdminLoginRequest
	err = common.BindRpcOption(req, rpcReq)
	if err != nil {
		Log.Errorln("bind rpc option err:", err)
	}
	resp, err := base.BaseClient.AdminLogin(ctx, rpcReq)
	if err != nil {
		Log.Errorln("admin login err:", err)
	}
	c.JSON(consts.StatusOK, resp)
}
