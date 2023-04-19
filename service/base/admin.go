package base

import (
	"context"
	"sign-lottery/kitex_gen/user"
	"sign-lottery/pkg/errmsg"
	"sign-lottery/pkg/jwt"
	. "sign-lottery/pkg/log"
)

// AdminLogin implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) AdminLogin(ctx context.Context, req *user.AdminLoginRequest) (resp *user.AdminLoginResponse, err error) {
	// TODO: Your code here...
	resp = new(user.AdminLoginResponse)
	name, password := req.GetName(), req.GetPassword()
	is_ok := s.dao.Admin.Login(ctx, name, password)
	if !is_ok {
		resp.Resp.Code = errmsg.NameOrPasswordError
		resp.Resp.Msg = errmsg.GetMsg(errmsg.NameOrPasswordError)
		return
	}
	resp.Token, err = jwt.GenAdminToken(name)
	if err != nil {
		Log.Errorln("get admin token err:", err)
		resp.Resp.Code = errmsg.Error
		resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
		return nil, err
	}
	err = s.cache.User.StoreToken(ctx, resp.Token, resp.Token)
	if err != nil {
		Log.Errorln("store admin token into cache err:", err)
		resp.Resp.Code = errmsg.Error
		resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
		return nil, err
	}
	resp.Resp.Code = errmsg.Success
	resp.Resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}
