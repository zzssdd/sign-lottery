package base

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"sign-lottery/cmd/api/biz/model/user"
)

// SendEmail .
// @router /user/email/ [POST]
func SendEmail(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.EmailRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.BaseResponse)

	c.JSON(consts.StatusOK, resp)
}

// Registe .
// @router /user/registe/ [POST]
func Registe(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.RegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.BaseResponse)

	c.JSON(consts.StatusOK, resp)
}

// Login .
// @router /user/login/ [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.LoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.LoginResponse)

	c.JSON(consts.StatusOK, resp)
}

// GetUserById .
// @router /user/id/ [GET]
func GetUserById(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.GetUserByIdRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.UsersResponse)

	c.JSON(consts.StatusOK, resp)
}

// GetUserByGid .
// @router /user/gid/ [GET]
func GetUserByGid(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.GetUserByGidRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.UsersResponse)

	c.JSON(consts.StatusOK, resp)
}

// ChangeUserAvater .
// @router /user/avater/ [PUT]
func ChangeUserAvater(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.ChangeUserAvaterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.BaseResponse)

	c.JSON(consts.StatusOK, resp)
}

// ChangeUserPassword .
// @router /user/password/ [PUT]
func ChangeUserPassword(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.ChangePasswordRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.BaseResponse)

	c.JSON(consts.StatusOK, resp)
}

// ChangeUserAddress .
// @router /user/address/ [PUT]
func ChangeUserAddress(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.ChangeAddressRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.BaseResponse)

	c.JSON(consts.StatusOK, resp)
}

// UserDel .
// @router /admin/user/ [DELETE]
func UserDel(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserDelRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.BaseResponse)

	c.JSON(consts.StatusOK, resp)
}

// GetAllUser .
// @router /admin/userlist/ [GET]
func GetAllUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.GetAllUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.UsersResponse)

	c.JSON(consts.StatusOK, resp)
}
