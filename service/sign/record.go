package sign

import (
	"context"
	"sign-lottery/dao/db/model"
	"sign-lottery/kitex_gen/sign"
	"sign-lottery/pkg/errmsg"
	. "sign-lottery/pkg/log"
	"sign-lottery/rabbitmq/consumer"
	model2 "sign-lottery/rabbitmq/model"
	"strconv"
	"time"
)

// GetMonthSign implements the SignServiceImpl interface.
func (s *SignServiceImpl) GetMonthSign(ctx context.Context, req *sign.GetMonthSignRequest) (resp *sign.MonthSignResponse, err error) {
	// TODO: Your code here...
	resp = new(sign.MonthSignResponse)
	uid, gid, month := req.GetUid(), req.GetGid(), req.GetMonth()
	var bitmap int
	if !s.cache.Sign.ExistUserMonthRecord(ctx, uid, gid, month) {
		bitmap, err = s.dao.Sign.GetMonthSign(ctx, uid, gid, month)
		if err != nil {
			Log.Errorln("get user month record err:", err)
			resp.Resp.Code = errmsg.Error
			resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
			return nil, err
		}
		err = s.cache.Sign.StoreUserMonthRecord(ctx, uid, gid, bitmap, month)
		if err != nil {
			Log.Errorln("store user month record to cache err:", err)
		}
	} else {
		bitmap, err = s.cache.Sign.GetUserMonthRecord(ctx, uid, gid, month)
		if err != nil {
			Log.Errorln("get bitmap from cache err:", err)
			resp.Resp.Code = errmsg.Error
			resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
			return nil, err
		}
	}
	resp.Bitmap = int32(bitmap)
	resp.Resp.Code = errmsg.Success
	resp.Resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

// GetMonthSignByGid implements the SignServiceImpl interface.
func (s *SignServiceImpl) GetMonthSignByGid(ctx context.Context, req *sign.GetMonthSignsByGid) (resp *sign.MonthSignsResponse, err error) {
	// TODO: Your code here...
	resp = new(sign.MonthSignsResponse)
	gid := req.GetGid()
	month := req.GetMonth()
	offset := req.GetOffset()
	limit := req.GetLimit()
	if !s.cache.Group.GroupOffsetExist(ctx, gid, offset, limit) {
		group, err := s.dao.Group.GetGroupById(ctx, gid)
		if err != nil {
			Log.Errorln("get offset groups from db err:", err)
			resp.Resp.Code = errmsg.Error
			resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
			return nil, err
		}
		err = s.cache.Group.StoreGroupOffset(ctx, gid, offset, limit, group)
		if err != nil {
			Log.Errorln("store group offset to cache err:", err)
		}
	}
	ids, count, err := s.cache.Group.GetUserByGid(ctx, gid, offset, limit)
	if err != nil {
		Log.Errorln("get users by gid err:", err)
		resp.Resp.Code = errmsg.Error
		resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
		return nil, err
	}
	bitmaps := []int32{}
	for _, Sid := range ids {
		id, _ := strconv.ParseInt(Sid, 10, 64)
		if !s.cache.Sign.ExistUserMonthRecord(ctx, id, gid, month) {
			bitmap, err := s.dao.Sign.GetMonthSign(ctx, id, gid, month)
			if err != nil {
				Log.Errorln("get user bitmap err:", err)
				resp.Resp.Code = errmsg.Error
				resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
				return nil, err
			}
			s.cache.Sign.StoreUserMonthRecord(ctx, id, gid, bitmap, month)
			bitmaps = append(bitmaps, int32(bitmap))
		}
	}
	resp.Bitmap = bitmaps
	resp.Count = count
	resp.Resp.Code = errmsg.Success
	resp.Resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

// GetAllRecord implements the SignServiceImpl interface.
func (s *SignServiceImpl) GetAllRecord(ctx context.Context, req *sign.GetAllRecordRequest) (resp *sign.RecordsResponse, err error) {
	// TODO: Your code here...
	resp = new(sign.RecordsResponse)
	limit, offset := req.GetLimit(), req.GetOffset()
	records, count, err := s.dao.Sign.GetAllRecord(ctx, int(offset), int(limit))
	if err != nil {
		Log.Errorln("get all record err:", err)
		resp.Resp.Code = errmsg.Error
		resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
		return nil, err
	}
	ret_records := []*sign.Record{}
	for _, v := range records {
		ret_record := &sign.Record{
			Uid:  v.UID,
			Gid:  v.Gid,
			Time: v.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		ret_records = append(ret_records, ret_record)
	}
	resp.Total = count
	resp.Records = ret_records
	resp.Resp.Code = errmsg.Success
	resp.Resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

// GetUserRecord implements the SignServiceImpl interface.
func (s *SignServiceImpl) GetUserRecord(ctx context.Context, req *sign.GetUserRecordRequest) (resp *sign.RecordsResponse, err error) {
	// TODO: Your code here...
	resp = new(sign.RecordsResponse)
	uid, offset, limit := req.GetUid(), req.GetOffset(), req.GetLimit()
	var records []*model.SignRecord
	var count int64
	if !s.cache.Sign.ExistUserRecord(ctx, uid, offset, limit) {
		records, count, err = s.dao.Sign.GetUserRecord(ctx, uid, int(offset), int(limit))
		if err != nil {
			Log.Errorln("get user record from db err:", err)
			resp.Resp.Code = errmsg.Error
			resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
			return nil, err
		}
		err = s.cache.Sign.StoreUserRecord(ctx, uid, limit, offset, records)
		if err != nil {
			Log.Errorln("store record to cache err:", err)
		}
	} else {
		records, err = s.cache.Sign.GetUserRecord(ctx, uid, offset, limit)
		if err != nil {
			Log.Errorln("get records from cache err:", err)
			resp.Resp.Code = errmsg.Error
			resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
			return nil, err
		}
		count = int64(len(records))
	}
	ret_records := []*sign.Record{}
	for _, v := range records {
		ret_record := &sign.Record{
			Uid:  v.UID,
			Gid:  v.Gid,
			Time: v.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		ret_records = append(ret_records, ret_record)
	}
	resp.Records = ret_records
	resp.Total = count
	resp.Resp.Code = errmsg.Success
	resp.Resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

func (s *SignServiceImpl) HandleSuccessSign(ctx context.Context) error {
	signConsumer := consumer.NewConsumer()
	signChan := make(chan model2.Record)
	err := signConsumer.Record.ConsumerRecord(signChan)
	if err != nil {
		Log.Errorln("consumer record from rabbitmq err:", err)
		return err
	}
	for signInfo := range signChan {
		uid := signInfo.Uid
		gid := signInfo.Gid
		month := time.Now().Format("2006-01-02")
		day := time.Now().Day()
		if !s.cache.Sign.ExistUserMonthRecord(ctx, uid, gid, month) {
			bitmap, _ := s.dao.Sign.GetMonthSign(ctx, uid, gid, month)
			s.cache.Sign.StoreUserMonthRecord(ctx, uid, gid, bitmap, month)
		}
		err = s.cache.Sign.UpdateUserMonthRecord(ctx, uid, gid, month, day)
		if err != nil {
			Log.Errorln("update user month record err:", err)
			continue
		}
		err = s.dao.Record.DoRecord(ctx, uid, gid, month, day)
		if err != nil {
			Log.Errorln("store sign to db err:", err)
			continue
		}
	}

	return nil
}
