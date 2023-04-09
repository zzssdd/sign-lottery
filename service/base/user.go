package base

import (
	"context"
	"sign-lottery/dao/db/model"
	"sign-lottery/kitex_gen/user"
	"sign-lottery/pkg/errmsg"
	"sign-lottery/pkg/jwt"
	. "sign-lottery/pkg/log"
	"sign-lottery/utils"
)

// SendEmail implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) SendEmail(ctx context.Context, req *user.EmailRequest) (resp *user.BaseResponse, err error) {
	email := req.GetEmail()
	resp = new(user.BaseResponse)
	code := utils.RandCode(6)
	err = s.cache.User.StoreCode(ctx, email, code)
	if err != nil {
		resp.Code = errmsg.CodeIsExpired
		resp.Msg = errmsg.GetMsg(errmsg.CodeIsExpired)
		return nil, nil
	}
	err = utils.SendEmail(email, code)
	if err != nil {
		Log.Errorln("send email err:", err)
		resp.Code = errmsg.SendEmailFailed
		resp.Msg = errmsg.GetMsg(errmsg.SendEmailFailed)
		return nil, err
	}
	resp.Code = errmsg.Success
	resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

// Registe implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) Registe(ctx context.Context, req *user.RegisterRequest) (resp *user.BaseResponse, err error) {
	resp = new(user.BaseResponse)
	email := req.GetEmail()
	password := req.GetPassword()
	name := req.GetName()
	code := req.Code
	real_code, err := s.cache.User.GetCode(ctx, email)
	if err != nil {
		resp.Code = errmsg.CodeNotExist
		resp.Msg = errmsg.GetMsg(errmsg.CodeNotExist)
		return nil, err
	}
	if real_code != code {
		resp.Code = errmsg.CodeIncorrect
		resp.Msg = errmsg.GetMsg(errmsg.CodeIncorrect)
		return nil, nil
	}
	u := &model.User{
		Email:    email,
		Name:     name,
		Password: password,
	}
	err = s.dao.User.Registe(ctx, u)
	if err != nil {
		Log.Errorln("registe err:", err)
		resp.Code = errmsg.Error
		resp.Msg = errmsg.GetMsg(errmsg.Error)
		return nil, err
	}
	resp.Code = errmsg.Success
	resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

// Login implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	resp = new(user.LoginResponse)
	email, password := req.GetEmail(), req.GetPassword()
	is_ok, id := s.dao.User.Login(ctx, email, password)
	if !is_ok {
		resp.Resp.Code = errmsg.EmailOrPasswordError
		resp.Resp.Msg = errmsg.GetMsg(errmsg.EmailOrPasswordError)
		return nil, nil
	}
	resp.Token, err = jwt.GenToken(id, email)
	if err != nil {
		Log.Errorln("get jwt err:", err)
		resp.Resp.Code = errmsg.Error
		resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
		return nil, err
	}
	resp.Resp.Code = errmsg.Success
	resp.Resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

// GetUserById implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) GetUserById(ctx context.Context, req *user.GetUserByIdRequest) (resp *user.UserResponse, err error) {
	resp = new(user.UserResponse)
	id := req.GetId()
	userInfo := new(model.User)
	if s.cache.User.UserInfoExist(ctx, id) {
		userInfo, err = s.cache.User.GetUserById(ctx, id)
		if err != nil {
			Log.Errorln("get info err:", err)
			resp.Resp.Code = errmsg.Error
			resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
			return nil, err
		}
	} else {
		userInfo, err = s.dao.User.GetUserById(ctx, id)
		if err != nil {
			Log.Errorln("get user from db err:", err)
			resp.Resp.Code = errmsg.Error
			resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
			return nil, err
		} else {
			err = s.cache.User.StoreUserInfo(ctx, id, userInfo)
			if err != nil {
				Log.Errorln("cache store user err:", err)
			}
		}
	}
	u := &user.UserInfo{
		Id:        id,
		CreatTime: userInfo.CreatedAt.Format("2006-01-02 15:04:05"),
		Email:     userInfo.Email,
		Name:      userInfo.Name,
		Avater:    *userInfo.Avater,
	}
	resp.User = u
	resp.Resp.Code = errmsg.Success
	resp.Resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

// GetAllUser implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) GetAllUser(ctx context.Context, req *user.GetAllUserRequest) (resp *user.UsersResponse, err error) {
	offset, limit := req.GetOffset(), req.GetLimit()
	users, count, err := s.dao.User.GetAllUser(ctx, int(offset), int(limit))
	if err != nil {
		Log.Errorln("get all user err:", err)
		resp.Resp.Code = errmsg.Error
		resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
		return nil, err
	}
	resp_users := []*user.UserInfo{}
	for _, v := range users {
		resp_user := &user.UserInfo{
			Id:        v.ID,
			CreatTime: v.CreatedAt.Format("2006-01-02 15:04:05"),
			Email:     v.Email,
			Name:      v.Name,
			Avater:    *v.Avater,
		}
		resp_users = append(resp_users, resp_user)
	}
	resp.Total = count
	resp.Resp.Code = errmsg.Success
	resp.Resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

// GetUserByGid implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) GetUserByGid(ctx context.Context, req *user.GetUserByGidRequest) (resp *user.UsersResponse, err error) {

	return
}

// ChangeUserAvater implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) ChangeUserAvater(ctx context.Context, req *user.ChangeUserAvaterRequest) (resp *user.BaseResponse, err error) {
	// TODO: Your code here...
	return
}

// ChangeUserPassword implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) ChangeUserPassword(ctx context.Context, req *user.ChangePasswordRequest) (resp *user.BaseResponse, err error) {
	// TODO: Your code here...
	return
}

// ChangeUserAddress implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) ChangeUserAddress(ctx context.Context, req *user.ChangeAddressRequest) (resp *user.BaseResponse, err error) {
	// TODO: Your code here...
	return
}

// UserDel implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) UserDel(ctx context.Context, req *user.UserDelRequest) (resp *user.BaseResponse, err error) {
	// TODO: Your code here...
	return
}
