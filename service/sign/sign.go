package sign

import (
	"context"
	"sign-lottery/dao/db/model"
	"sign-lottery/kitex_gen/sign"
	"sign-lottery/pkg/errmsg"
	. "sign-lottery/pkg/log"
	"sign-lottery/rabbitmq/consumer"
	model2 "sign-lottery/rabbitmq/model"
	"time"
)

// Sign implements the SignServiceImpl interface.
func (s *SignServiceImpl) Sign(ctx context.Context) (resp *sign.BaseResponse, err error) {
	// TODO: Your code here...
	signChan := make(chan model2.Sign)
	err = consumer.NewConsumer().Sign.ConsumerSign(signChan)
	if err != nil {
		Log.Fatalln("create consumer err:", err)
	}
	for signInfo := range signChan {
		uid := signInfo.Uid
		gid := signInfo.Gid
		latitude := signInfo.Latitude
		longtitude := signInfo.Longtitude
		ip := signInfo.Ip
		ok := s.cache.Sign.PosLimit(ctx, gid, latitude, longtitude)
		if !ok {
			err = s.cache.HandlerErr.ReturnSignErr(ctx, uid, gid, errmsg.SignNotInPos)
			if err != nil {
				Log.Errorln("store return sign code err:", err)
			}
			continue
		}
		ok = s.cache.Sign.IpLimit(ctx, ip, uid)
		if !ok {
			err = s.cache.HandlerErr.ReturnSignErr(ctx, uid, gid, errmsg.SignIpUsed)
			if err != nil {
				Log.Errorln("store return sign code err:", err)
			}
			continue
		}
		var group *model.SignGroup
		if !s.cache.Group.GroupInfoExist(ctx, gid) {
			group, err = s.dao.Group.GetGroupById(ctx, gid)
			if err != nil {
				Log.Errorln("get group info from db err:", err)
				err2 := s.cache.HandlerErr.ReturnSignErr(ctx, uid, gid, errmsg.Error)
				if err2 != nil {
					Log.Errorln("store return sign code err:", err)
				}
				continue
			}
			err = s.cache.Group.StoreGroupInfo(ctx, group)
			if err != nil {
				Log.Errorln("store group to cache err:", err)
				continue
			}
		} else {
			group, err = s.cache.Group.GetGroupInfo(ctx, gid)
			if err != nil {
				Log.Errorln("get group info from cache err:", err)
				err2 := s.cache.HandlerErr.ReturnSignErr(ctx, uid, gid, errmsg.Error)
				if err2 != nil {
					Log.Errorln("store return code err:", err)
				}
				continue
			}
		}
		now := time.Now()
		if now.Day() == group.Start.Day() {
			if now.Before(group.Start) {
				err = s.cache.Sign.UserSignStart(ctx, uid, gid)
			} else if now.After(group.End) {
				err = s.cache.Sign.UserSignEnd(ctx, uid, gid)
			} else {
				err2 := s.cache.HandlerErr.ReturnSignErr(ctx, uid, gid, errmsg.SignNotInTime)
				if err2 != nil {
					Log.Errorln("store return code err:", err)
				}
				continue
			}
		} else {
			err2 := s.cache.HandlerErr.ReturnSignErr(ctx, uid, gid, errmsg.SignNotInTime)
			if err2 != nil {
				Log.Errorln("store return code err:", err)
			}
			continue
		}
		if err != nil {
			err2 := s.cache.HandlerErr.ReturnSignErr(ctx, uid, gid, errmsg.Error)
			if err2 != nil {
				Log.Errorln("store return code err:", err)
			}
			Log.Errorln("add sign info to cache err:", err)
			continue
		}
		err2 := s.cache.HandlerErr.ReturnSignErr(ctx, uid, gid, errmsg.Success)
		if err2 != nil {
			Log.Errorln("store return code err:", err)
		}
	}
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
