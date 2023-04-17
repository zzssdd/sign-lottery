package sign

import (
	"context"
	"sign-lottery/dao/db/model"
	"sign-lottery/kitex_gen/sign"
	"sign-lottery/pkg/errmsg"
	. "sign-lottery/pkg/log"
	"time"
)

// SignPosAdd implements the SignServiceImpl interface.
func (s *SignServiceImpl) SignPosAdd(ctx context.Context, req *sign.SignPosAddRequest) (resp *sign.BaseResponse, err error) {
	// TODO: Your code here...
	resp = new(sign.BaseResponse)
	gid := req.GetGid()
	name := req.GetName()
	longtitude := req.GetLongtitude()
	latitude := req.GetLatitude()
	uid := req.GetUid()
	previlege := s.dao.Group.CheckGroupPrevilege(ctx, uid, gid)
	if !previlege {
		resp.Code = errmsg.NoPreviledge
		resp.Msg = errmsg.GetMsg(errmsg.NoPreviledge)
		return
	}
	pos := &model.SignGroupPos{
		Name:       name,
		Gid:        int(gid),
		Latitude:   latitude,
		Longtitude: longtitude,
	}
	err = s.dao.Sign.SignPosAdd(ctx, pos)
	if err != nil {
		resp.Code = errmsg.Error
		resp.Msg = errmsg.GetMsg(errmsg.Error)
		return nil, err
	}
	err = s.cache.Sign.SignPosAdd(ctx, gid, pos)
	if err != nil {
		Log.Errorln("add pos to cache err:", err)
		resp.Code = errmsg.Error
		resp.Msg = errmsg.GetMsg(errmsg.Error)
		return nil, err
	}
	resp.Code = errmsg.Success
	resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

// SignPosDel implements the SignServiceImpl interface.
func (s *SignServiceImpl) SignPosDel(ctx context.Context, req *sign.SignPosDelRequest) (resp *sign.BaseResponse, err error) {
	// TODO: Your code here...
	resp = new(sign.BaseResponse)
	uid, gid, name := req.GetUid(), req.GetGid(), req.GetName()
	previlege := s.dao.Group.CheckGroupPrevilege(ctx, uid, gid)
	if !previlege {
		resp.Code = errmsg.NoPreviledge
		resp.Msg = errmsg.GetMsg(errmsg.NoPreviledge)
		return
	}
	err = s.dao.Sign.SignPosDel(ctx, gid, name)
	if err != nil {
		Log.Errorln("delete sign pos from db err:", err)
		resp.Code = errmsg.Error
		resp.Msg = errmsg.GetMsg(errmsg.Error)
		return nil, err
	}
	err = s.cache.Sign.SignPosDel(ctx, gid, name)
	if err != nil {
		Log.Errorln("delete sign pos err:", err)
	}
	go func() {
		time.Sleep(time.Millisecond)
		s.cache.Sign.SignPosDel(ctx, gid, name)
	}()
	resp.Code = errmsg.Success
	resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

// GetSignPos implements the SignServiceImpl interface.
func (s *SignServiceImpl) GetSignPos(ctx context.Context, req *sign.GetSignPosRequest) (resp *sign.GetSignPosResponse, err error) {
	// TODO: Your code here...
	resp = new(sign.GetSignPosResponse)
	gid, limit, offset := req.GetGid(), req.GetLimit(), req.GetOffset()
	var count int64
	var signpos []*model.SignGroupPos
	signpos, count, err = s.dao.Sign.SignPosGet(ctx, gid, int(offset), int(limit))
	if err != nil {
		Log.Errorln("get sign pos from db err:", err)
		resp.Resp.Code = errmsg.Error
		resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
		return nil, err
	}
	for _, v := range signpos {
		err = s.cache.Sign.SignPosAdd(ctx, gid, v)
		if err != nil {
			Log.Errorln("store pos into cache err:", err)
			resp.Resp.Code = errmsg.Error
			resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
			return nil, err
		}
	}

	ret_poses := []*sign.PosInfo{}
	for _, v := range signpos {
		ret_pos := &sign.PosInfo{
			Name:       v.Name,
			Longtitude: v.Longtitude,
			Latitude:   v.Latitude,
		}
		ret_poses = append(ret_poses, ret_pos)
	}
	resp.Pos = ret_poses
	resp.Total = count
	resp.Resp.Code = errmsg.Success
	resp.Resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}
