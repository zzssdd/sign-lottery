// Code generated by hertz generator.

package api

import (
	"context"
	"sign-lottery/cmd/api/biz/model/lottery"
	"sign-lottery/cmd/api/biz/model/sign"
	"sign-lottery/cmd/api/biz/model/user"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
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

// CreateGroup .
// @router /group/add/ [POST]
func CreateGroup(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.CreateGroupRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.BaseResponse)

	c.JSON(consts.StatusOK, resp)
}

// JoinGroup .
// @router /group/join/ [POST]
func JoinGroup(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.JoinGroupRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.BaseResponse)

	c.JSON(consts.StatusOK, resp)
}

// GetGroupById .
// @router /group/id/ [GET]
func GetGroupById(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.GetGroupByIdRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.GroupResponse)

	c.JSON(consts.StatusOK, resp)
}

// GetAllGroup .
// @router /group/list/ [GET]
func GetAllGroup(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.GetAllGroupRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.GroupsResponse)

	c.JSON(consts.StatusOK, resp)
}

// GroupUpdate .
// @router /group/put/ [PUT]
func GroupUpdate(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.GroupUpdateRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.BaseResponse)

	c.JSON(consts.StatusOK, resp)
}

// GroupDel .
// @router /group/del/ [DELETE]
func GroupDel(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.GroupDelRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(user.BaseResponse)

	c.JSON(consts.StatusOK, resp)
}

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