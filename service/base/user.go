package base

import (
	"context"
	"gorm.io/gorm"
	"sign-lottery/dao/db/model"
	"sign-lottery/kitex_gen/user"
	"sign-lottery/pkg/errmsg"
	"sign-lottery/pkg/jwt"
	. "sign-lottery/pkg/log"
	"sign-lottery/utils"
	"strconv"
	"time"
)

// SendEmail implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) SendEmail(ctx context.Context, req *user.EmailRequest) (resp *user.BaseResponse, err error) {
	email := req.GetEmail()
	resp = new(user.BaseResponse)
	code := utils.RandCode(6)
	err = s.cache.User.StoreCode(ctx, email, code)
	if err != nil {
		resp.Code = errmsg.Error
		resp.Msg = errmsg.GetMsg(errmsg.Error)
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
	if !s.cache.User.ExistCode(ctx, email) {
		resp.Code = errmsg.CodeNotExist
		resp.Msg = errmsg.GetMsg(errmsg.CodeNotExist)
		return nil, err
	}
	real_code, err := s.cache.User.GetCode(ctx, email)
	if err != nil {
		resp.Code = errmsg.CodeIsExpired
		resp.Msg = errmsg.GetMsg(errmsg.CodeIsExpired)
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
	id, exist := s.dao.User.CheckUserIsExist(ctx, email)
	if exist {
		resp.Code = errmsg.EmailHasExist
		resp.Msg = errmsg.GetMsg(errmsg.EmailHasExist)
		return
	}
	err, id = s.dao.User.Registe(ctx, u)
	if err != nil {
		Log.Errorln("registe err:", err)
		resp.Code = errmsg.Error
		resp.Msg = errmsg.GetMsg(errmsg.Error)
		return nil, err
	}
	err = s.cache.User.StoreLoginInfo(ctx, id, email, password)
	if err != nil {
		Log.Errorln("store user login info to cache err:", err)
	}
	err = s.cache.User.StoreUserInfo(ctx, id, u)
	if err != nil {
		Log.Errorln("store user info to cache err:", err)
	}
	resp.Code = errmsg.Success
	resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

// Login implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	resp = new(user.LoginResponse)
	email, password := req.GetEmail(), req.GetPassword()
	if !s.cache.User.ExistsLoginInfo(ctx, email) {
		user, err := s.dao.User.GetPasswordAndIdByEmail(ctx, email)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				resp.Resp.Code = errmsg.EmailNotExist
				resp.Resp.Msg = errmsg.GetMsg(errmsg.EmailNotExist)
				return nil, err
			}
			Log.Errorln("get password from db err:", err)
			resp.Resp.Code = errmsg.Error
			resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
			return nil, err
		}
		err = s.cache.User.StoreLoginInfo(ctx, user.ID, user.Email, user.Password)
		if err != nil {
			Log.Errorln("store login info to cache err:", err)
			resp.Resp.Code = errmsg.Error
			resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
			return nil, err
		}
	}
	LoginInfo := s.cache.User.GetLoginInfo(ctx, email)
	if password != LoginInfo["password"] {
		resp.Resp.Code = errmsg.PasswordError
		resp.Resp.Msg = errmsg.GetMsg(errmsg.PasswordError)
		return nil, nil
	}
	id, _ := strconv.ParseInt(LoginInfo["id"], 10, 64)
	resp.Token, err = jwt.GenToken(id, email)
	if err != nil {
		Log.Errorln("get jwt err:", err)
		resp.Resp.Code = errmsg.Error
		resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
		return nil, err
	}
	resp.Resp.Code = errmsg.Success
	resp.Resp.Msg = errmsg.GetMsg(errmsg.Success)
	resp.Id = id
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
	resp = new(user.UsersResponse)
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
	resp = new(user.UsersResponse)
	gid, offset, limit := req.GetGid(), req.GetOffset(), req.GetLimit()
	var users []*model.User
	var count int64
	if s.cache.Group.GroupOffsetExist(ctx, gid, offset, limit) {
		var uids []string
		uids, count, err = s.cache.Group.GetUserByGid(ctx, gid, offset, limit)
		if err != nil {
			Log.Errorln("get group offset err:", err)
			resp.Resp.Code = errmsg.Error
			resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
			return
		}
		users, err = s.cache.User.GetUsersById(ctx, uids)
		if err != nil {
			Log.Errorln("get gourp users err:", err)
			resp.Resp.Code = errmsg.Error
			resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
			return
		}
	} else {
		users, count, err = s.dao.User.GetUserByGid(ctx, gid, int(offset), int(limit))
		if err != nil {
			Log.Errorln("get users form db err:", err)
			resp.Resp.Code = errmsg.Error
			resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
			return
		}
		err = s.cache.User.StoreUsersInfo(ctx, users)
		if err != nil {
			Log.Errorln("store users to cache err:", err)
			resp.Resp.Code = errmsg.Error
			resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
			return
		}
		group, err := s.dao.Group.GetGroupById(ctx, gid)
		if err != nil {
			Log.Errorln("get group info from db err:", err)
			resp.Resp.Code = errmsg.Error
			resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
			return
		}
		s.cache.Group.StoreGroupOffset(ctx, gid, offset, limit, group)
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
	resp.Resp.Code = errmsg.Success
	resp.Resp.Msg = errmsg.GetMsg(errmsg.Success)
	resp.Users = resp_users
	resp.Total = count
	return
}

// ChangeUserAvater implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) ChangeUserAvater(ctx context.Context, req *user.ChangeUserAvaterRequest) (resp *user.BaseResponse, err error) {
	resp = new(user.BaseResponse)
	id := req.GetId()
	avater := req.GetAvater()
	err = s.dao.User.ChangeUserAvater(ctx, id, avater)
	if err != nil {
		Log.Errorln("change avater from db err:", err)
		resp.Code = errmsg.Error
		resp.Msg = errmsg.GetMsg(errmsg.Error)
		return
	}
	if s.cache.User.UserInfoExist(ctx, id) {
		err = s.cache.User.ClearUserInfo(ctx, id)
		if err != nil {
			Log.Errorln("clear user info from cache err:", err)
		}
		go func() {
			time.Sleep(time.Millisecond)
			s.cache.User.ClearUserInfo(ctx, id)
		}()
	}
	resp.Code = errmsg.Success
	resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

// ChangeUserPassword implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) ChangeUserPassword(ctx context.Context, req *user.ChangePasswordRequest) (resp *user.BaseResponse, err error) {
	resp = new(user.BaseResponse)
	id, oldP, newP := req.GetId(), req.GetOld(), req.GetNew_()
	if oldP == newP {
		resp.Code = errmsg.NewEqualOld
		resp.Msg = errmsg.GetMsg(errmsg.NewEqualOld)
		return
	}
	email, err := s.dao.User.ChangePassword(ctx, id, oldP, newP)
	if err != nil {
		resp.Code = errmsg.OldPassworError
		resp.Msg = errmsg.GetMsg(errmsg.OldPassworError)
		return
	}
	if s.cache.User.ExistsLoginInfo(ctx, *email) {
		err = s.cache.User.ClearLoginInfo(ctx, *email)
		if err != nil {
			Log.Errorln("clear user info from cache err:", err)
		}
		go func() {
			time.Sleep(time.Millisecond)
			s.cache.User.ClearLoginInfo(ctx, *email)
		}()
	}
	resp.Code = errmsg.Success
	resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

// ChangeUserAddress implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) ChangeUserAddress(ctx context.Context, req *user.ChangeAddressRequest) (resp *user.BaseResponse, err error) {
	resp = new(user.BaseResponse)
	id, address := req.GetId(), req.GetAddress()
	err = s.dao.User.ChangeAddress(ctx, id, address)
	if err != nil {
		Log.Errorln("change address err:", err)
		resp.Code = errmsg.Error
		resp.Msg = errmsg.GetMsg(errmsg.Error)
		return
	}
	if s.cache.User.UserInfoExist(ctx, id) {
		err = s.cache.User.ClearUserInfo(ctx, id)
		if err != nil {
			Log.Errorln("clear user info from cache err:", err)
		}
		go func() {
			time.Sleep(time.Millisecond)
			s.cache.User.ClearUserInfo(ctx, id)
		}()
	}
	resp.Code = errmsg.Success
	resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

// UserDel implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) UserDel(ctx context.Context, req *user.UserDelRequest) (resp *user.BaseResponse, err error) {
	// TODO: Your code here...
	resp = new(user.BaseResponse)
	id := req.GetId()
	err = s.dao.User.UserDel(ctx, id)
	if err != nil {
		Log.Errorln("delete user err:", err)
		resp.Code = errmsg.Error
		resp.Msg = errmsg.GetMsg(errmsg.Error)
		return
	}
	if s.cache.User.UserInfoExist(ctx, id) {
		err = s.cache.User.ClearUserInfo(ctx, id)
		if err != nil {
			Log.Errorln("clear user info from cache err:", err)
		}
		go func() {
			time.Sleep(time.Millisecond)
			s.cache.User.ClearUserInfo(ctx, id)
		}()
	}
	resp.Code = errmsg.Success
	resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}
