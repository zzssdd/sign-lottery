package lottery

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"sign-lottery/cmd/api/biz/model/lottery"
)

// Choose .
// @router /prize/choose/ [GET]
func Choose(ctx context.Context, c *app.RequestContext) {
	var err error
	var req lottery.ChooseRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(lottery.ChooseResponse)

	c.JSON(consts.StatusOK, resp)
}
