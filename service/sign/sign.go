package sign

import (
	"context"
	"sign-lottery/dao/db/model"
	"sign-lottery/kitex_gen/sign"
	"sign-lottery/pkg/errmsg"
	. "sign-lottery/pkg/log"
	"time"
)

// Sign implements the SignServiceImpl interface.
func (s *SignServiceImpl) Sign(ctx context.Context, req *sign.SignRequest) (resp *sign.BaseResponse, err error) {
	// TODO: Your code here...
	resp = new(sign.BaseResponse)
	uid := req.GetUid()
	gid := req.GetGid()
	latitude := req.GetLatitude()
	longtitude := req.GetLongtitude()
	ip := req.GetIp()
	ok := s.cache.Sign.PosLimit(ctx, gid, latitude, longtitude)
	if !ok {
		resp.Code = errmsg.SignNotInPos
		resp.Msg = errmsg.GetMsg(errmsg.SignNotInPos)
		return
	}
	ok = s.cache.Sign.IpLimit(ctx, ip, uid)
	if !ok {
		resp.Code = errmsg.SignIpUsed
		resp.Msg = errmsg.GetMsg(errmsg.SignIpUsed)
		return
	}
	var group *model.SignGroup
	if !s.cache.Group.GroupInfoExist(ctx, gid) {
		group, err = s.dao.Group.GetGroupById(ctx, gid)
		if err != nil {
			Log.Errorln("get group info from db err:", err)
			resp.Code = errmsg.Error
			resp.Msg = errmsg.GetMsg(errmsg.Error)
			return nil, err
		}
		err = s.cache.Group.StoreGroupInfo(ctx, group)
		if err != nil {
			Log.Errorln("store group to cache err:", err)
		}
	} else {
		group, err = s.cache.Group.GetGroupInfo(ctx, gid)
		if err != nil {
			Log.Errorln("get group info from cache err:", err)
			resp.Code = errmsg.Error
			resp.Msg = errmsg.GetMsg(errmsg.Error)
			return nil, err
		}
	}
	now := time.Now()
	if now.Day() == group.Start.Day() {
		if now.Before(group.Start) {
			err = s.cache.Sign.UserSignStart(ctx, uid, gid)
		} else if now.After(group.End) {
			err = s.cache.Sign.UserSignEnd(ctx, uid, gid)
		} else {
			resp.Code = errmsg.SignNotInTime
			resp.Msg = errmsg.GetMsg(errmsg.SignNotInTime)
			return
		}
	} else {
		resp.Code = errmsg.SignNotInTime
		resp.Msg = errmsg.GetMsg(errmsg.SignNotInTime)
		return
	}
	if err != nil {
		Log.Errorln("add sign info to cache err:", err)
		resp.Code = errmsg.Error
		resp.Msg = errmsg.GetMsg(errmsg.Error)
		return nil, err
	}
	s.cache.Sign.IpLimitAdd(ctx, ip, uid)
	resp.Code = errmsg.Success
	resp.Msg = errmsg.GetMsg(errmsg.Success)
	//month := strconv.Itoa(time.Now().Year()) + "-" + time.Now().Month().String() + "-" + strconv.Itoa(time.Now().Day())
	return
}

// AskLeave implements the SignServiceImpl interface.
func (s *SignServiceImpl) AskLeave(ctx context.Context, req *sign.AskLeaveRequest) (resp *sign.BaseResponse, err error) {
	// TODO: Your code here...
	resp = new(sign.BaseResponse)
	uid := req.GetUid()
	gid := req.GetGid()
	leave_time := time.Now()
	issue := req.Issue
	leaveInfo := &model.AskLeave{
		UID:   uid,
		Issue: issue,
		Time:  &leave_time,
		Gid:   int(gid),
	}
	err = s.dao.Sign.AskLeave(ctx, leaveInfo)
	if err != nil {
		Log.Errorln("ask leave to db err:", err)
		resp.Code = errmsg.Error
		resp.Msg = errmsg.GetMsg(errmsg.Error)
		return nil, err
	}
	resp.Code = errmsg.Success
	resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}
