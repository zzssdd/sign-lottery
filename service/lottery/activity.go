package lottery

import (
	"context"
	"sign-lottery/dao/db/model"
	"sign-lottery/kitex_gen/lottery"
	"sign-lottery/pkg/errmsg"
	. "sign-lottery/pkg/log"
	"time"
)

func (s *LotteryServiceImpl) ActivityAdd(ctx context.Context, req *lottery.ActivityAddRequest) (resp *lottery.BaseResponse, err error) {
	//TODO implement me
	resp = new(lottery.BaseResponse)
	name := req.GetName()
	picture := req.GetPicture()
	desc := req.GetDec()
	cost := req.GetCost()
	uid := req.GetUid()
	gid := int(req.GetGid())
	start, _ := time.Parse("2006-01-02 15:04:05", req.GetStart())
	end, _ := time.Parse("2006-01-02 15:04:05", req.GetEnd())
	activity := &model.Activity{
		Name:    name,
		Picture: &picture,
		Des:     &desc,
		Cost:    &cost,
		UID:     &uid,
		Gid:     gid,
		Start:   &start,
		End:     &end,
	}
	err, id := s.dao.Activity.ActivityAdd(ctx, activity)
	if err != nil {
		Log.Errorln("activity add err:", err)
		resp.Code = errmsg.Error
		resp.Msg = errmsg.GetMsg(errmsg.Error)
		return nil, err
	}
	err = s.cache.Activity.StoreActivityInfo(ctx, int32(id), activity)
	if err != nil {
		Log.Errorln("store activity to cache err:", err)
	}
	resp.Code = errmsg.Success
	resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

func (s *LotteryServiceImpl) ActivityDel(ctx context.Context, req *lottery.ActivityDelRequest) (resp *lottery.BaseResponse, err error) {
	//TODO implement me
	resp = new(lottery.BaseResponse)
	id := req.GetId()
	err = s.dao.Activity.ActivityDel(ctx, id)
	if err != nil {
		Log.Errorln("delete activity from db err:", err)
		resp.Code = errmsg.Error
		resp.Msg = errmsg.GetMsg(errmsg.Error)
		return nil, err
	}
	if s.cache.Activity.ExistActivityInfo(ctx, id) {
		err = s.cache.Activity.ClearActivityInfo(ctx, id)
		if err != nil {
			Log.Errorln("clear activity info from cache err:", err)
		}
		go func() {
			time.Sleep(time.Millisecond)
			s.cache.Activity.ClearActivityInfo(ctx, id)
		}()
	}
	resp.Code = errmsg.Success
	resp.Msg = errmsg.GetMsg(errmsg.Success)
	return nil, err
}

func (s *LotteryServiceImpl) ActivityUpdate(ctx context.Context, req *lottery.ActivityUpdateRequest) (resp *lottery.BaseResponse, err error) {
	//TODO implement me
	resp = new(lottery.BaseResponse)
	id := req.GetId()
	name := req.GetName()
	picture := req.GetPicture()
	desc := req.GetDesc()
	cost := req.GetCost()
	uid := req.GetUid()
	gid := int(req.GetGid())
	start, _ := time.Parse("2006-01-02 15:04:05", req.GetStart())
	end, _ := time.Parse("2006-01-02 15:04:05", req.GetEnd())
	activity := &model.Activity{
		Name:    name,
		Picture: &picture,
		Des:     &desc,
		Cost:    &cost,
		UID:     &uid,
		Gid:     gid,
		Start:   &start,
		End:     &end,
	}
	err = s.dao.Activity.ActivityUpdate(ctx, id, activity)
	if err != nil {
		Log.Errorln("update activity to db err:", err)
		resp.Code = errmsg.Error
		resp.Msg = errmsg.GetMsg(errmsg.Error)
		return nil, err
	}
	if s.cache.Activity.ExistActivityInfo(ctx, id) {
		err = s.cache.Activity.ClearActivityInfo(ctx, id)
		if err != nil {
			Log.Errorln("clear activity from cache err:", err)
		}
		go func() {
			time.Sleep(time.Millisecond)
			s.cache.Activity.ClearActivityInfo(ctx, id)
		}()
	}
	resp.Code = errmsg.Success
	resp.Msg = errmsg.GetMsg(errmsg.Success)
	return nil, err
}

func (s *LotteryServiceImpl) GetActivityByGid(ctx context.Context, req *lottery.GetActivityByGidRequest) (resp *lottery.ActivitysResponse, err error) {
	//TODO implement me
	resp = new(lottery.ActivitysResponse)
	gid, offset, limit := req.GetGid(), req.GetOffset(), req.GetLimit()
	var activitys []*model.Activity
	var count int64
	if !s.cache.Activity.ExistActivityOffset(ctx, gid, offset, limit) {
		activitys, count, err = s.dao.Activity.GetActivityByGid(ctx, gid, int(offset), int(limit))
		if err != nil {
			Log.Errorln("get activity from db err:", err)
			resp.Resp.Code = errmsg.Error
			resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
			return nil, err
		}
		err = s.cache.Activity.StoreActivityOffset(ctx, gid, offset, limit, activitys)
		if err != nil {
			Log.Errorln("store activity to cache err:", err)
		}
	} else {
		activitys, err = s.cache.Activity.GetActivityOffset(ctx, gid, offset, limit)
		if err != nil {
			Log.Errorln("get activity from cache err:", err)
			resp.Resp.Code = errmsg.Error
			resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
			return nil, err
		}
		count = int64(len(activitys))
	}
	ret_activitys := []*lottery.Activity{}
	for _, v := range activitys {
		ret_activity := &lottery.Activity{
			Id:        int32(v.ID),
			CreatTime: v.CreatedAt.Format("2006-01-02 15:04:05"),
			Name:      v.Name,
			Picture:   *v.Picture,
			Desc:      *v.Des,
			Cost:      *v.Cost,
			Uid:       *v.UID,
			Gid:       int32(v.Gid),
			Start:     v.Start.Format("2006-01-02 15:04:05"),
			End:       v.End.Format("2006-01-02 15:04:05"),
			Count:     *v.Num,
		}
		ret_activitys = append(ret_activitys, ret_activity)
	}
	resp.Activity = ret_activitys
	resp.Total = count
	resp.Resp.Code = errmsg.Success
	resp.Resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

func (s *LotteryServiceImpl) GetAllActivity(ctx context.Context, req *lottery.GetAllActivityRequest) (resp *lottery.ActivitysResponse, err error) {
	//TODO implement me
	resp = new(lottery.ActivitysResponse)
	offset, limit := int(req.GetOffset()), int(req.GetLimit())
	activitys, count, err := s.dao.Activity.GetAllActivity(ctx, offset, limit)
	if err != nil {
		Log.Errorln("get all activity from db err:", err)
		resp.Resp.Code = errmsg.Error
		resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
		return nil, err
	}
	ret_activitys := []*lottery.Activity{}
	for _, v := range activitys {
		ret_activity := &lottery.Activity{
			Id:        int32(v.ID),
			CreatTime: v.CreatedAt.Format("2006-01-02 15:04:05"),
			Name:      v.Name,
			Picture:   *v.Picture,
			Desc:      *v.Des,
			Cost:      *v.Cost,
			Uid:       *v.UID,
			Gid:       int32(v.Gid),
			Start:     v.Start.Format("2006-01-02 15:04:05"),
			End:       v.End.Format("2006-01-02 15:04:05"),
			Count:     *v.Num,
		}
		ret_activitys = append(ret_activitys, ret_activity)
	}
	resp.Activity = ret_activitys
	resp.Total = count
	resp.Resp.Code = errmsg.Success
	resp.Resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

func (s *LotteryServiceImpl) GetActivityById(ctx context.Context, req *lottery.GetActivityByIdRequest) (resp *lottery.ActivityResponse, err error) {
	//TODO implement me
	resp = new(lottery.ActivityResponse)
	id := req.GetId()
	var activity *model.Activity
	if !s.cache.Activity.ExistActivityInfo(ctx, id) {
		activity, err = s.dao.Activity.GetActivityById(ctx, id)
		if err != nil {
			Log.Errorln("get activity from db err:", err)
			resp.Resp.Code = errmsg.Error
			resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
			return nil, err
		}
		err = s.cache.Activity.StoreActivityInfo(ctx, id, activity)
		if err != nil {
			Log.Errorln("store activity to cache err:", err)
		}
	} else {
		activity, err = s.cache.Activity.GetActivityById(ctx, id)
		if err != nil {
			Log.Errorln("get activity from cache err:", err)
		}
	}
	ret_activity := &lottery.Activity{
		Id:        int32(activity.ID),
		CreatTime: activity.CreatedAt.Format("2006-01-02 15:04:05"),
		Name:      activity.Name,
		Picture:   *activity.Picture,
		Desc:      *activity.Des,
		Cost:      *activity.Cost,
		Uid:       *activity.UID,
		Gid:       int32(activity.Gid),
		Start:     activity.Start.Format("2006-01-02 15:04:05"),
		End:       activity.End.Format("2006-01-02 15:04:05"),
		Count:     *activity.Num,
	}
	resp.Activity = ret_activity
	resp.Resp.Code = errmsg.Success
	resp.Resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}
