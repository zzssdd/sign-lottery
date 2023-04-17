package lottery

import (
	"context"
	"sign-lottery/dao/db/model"
	"sign-lottery/kitex_gen/lottery"
	"sign-lottery/pkg/errmsg"
	. "sign-lottery/pkg/log"
	"time"
)

// PrizeAdd implements the LotteryServiceImpl interface.
func (s *LotteryServiceImpl) PrizeAdd(ctx context.Context, req *lottery.PrizeAddRequest) (resp *lottery.BaseResponse, err error) {
	// TODO: Your code here...
	resp = new(lottery.BaseResponse)
	name := req.GetName()
	num := req.GetNum()
	picture := req.GetPicture()
	aid := req.GetAid()
	prize := &model.Prize{
		Name:    name,
		Num:     &num,
		Picture: &picture,
		Aid:     int(aid),
	}
	err = s.dao.Prize.PrizeAdd(ctx, prize)
	if err != nil {
		Log.Errorln("add prize into db err:", err)
		resp.Code = errmsg.Error
		resp.Msg = errmsg.GetMsg(errmsg.Error)
		return nil, err
	}
	err = s.cache.Prize.StorePrizeByAid(ctx, aid, []*model.Prize{prize})
	if err != nil {
		Log.Errorln("add prize into cache err:", err)
	}
	resp.Code = errmsg.Success
	resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

// PrizeDel implements the LotteryServiceImpl interface.
func (s *LotteryServiceImpl) PrizeDel(ctx context.Context, req *lottery.PrizeDelRequest) (resp *lottery.BaseResponse, err error) {
	// TODO: Your code here...
	resp = new(lottery.BaseResponse)
	id := req.GetId()
	uid := req.GetUid()
	aid, err := s.dao.Prize.GetPrizeAid(ctx, id)
	if err != nil {
		Log.Errorln("get prize aid from db err:", err)
		resp.Code = errmsg.Error
		resp.Msg = errmsg.GetMsg(errmsg.Error)
		return
	}
	previlege := s.dao.Activity.CheckActivityPrevilege(ctx, uid, aid)
	if !previlege {
		resp.Code = errmsg.NoPreviledge
		resp.Msg = errmsg.GetMsg(errmsg.NoPreviledge)
		return
	}
	err = s.dao.Prize.PrizeDel(ctx, id)
	if err != nil {
		Log.Errorln("delete prize from db err:", err)
		resp.Code = errmsg.Error
		resp.Msg = errmsg.GetMsg(errmsg.Error)
		return nil, err
	}
	if s.cache.Prize.ExistPrize(ctx, id) {
		err = s.cache.Prize.ClearPrize(ctx, id)
		if err != nil {
			Log.Errorln("clear prize cache err:", err)
		}
		go func() {
			time.Sleep(time.Millisecond)
			s.cache.Prize.ClearPrize(ctx, id)
		}()
	}
	resp.Code = errmsg.Success
	resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

// PrizeUpdate implements the LotteryServiceImpl interface.
func (s *LotteryServiceImpl) PrizeUpdate(ctx context.Context, req *lottery.PrizeUpdateRequest) (resp *lottery.BaseResponse, err error) {
	// TODO: Your code here...
	resp = new(lottery.BaseResponse)
	id := req.GetId()
	name := req.GetName()
	num := req.GetNum()
	picture := req.GetPicture()
	aid := req.GetAid()
	uid := req.GetUid()
	previlege := s.dao.Activity.CheckActivityPrevilege(ctx, uid, aid)
	if !previlege {
		resp.Code = errmsg.NoPreviledge
		resp.Msg = errmsg.GetMsg(errmsg.NoPreviledge)
		return
	}
	prize := &model.Prize{
		ID:      int(id),
		Name:    name,
		Num:     &num,
		Picture: &picture,
		Aid:     int(aid),
	}
	err = s.dao.Prize.PrizeUpdate(ctx, id, prize)
	if err != nil {
		Log.Errorln("update prize from db err:", err)
		resp.Code = errmsg.Error
		resp.Msg = errmsg.GetMsg(errmsg.Error)
		return nil, err
	}
	if s.cache.Prize.ExistPrize(ctx, id) {
		err = s.cache.Prize.ClearPrize(ctx, id)
		if err != nil {
			Log.Errorln("clear prize from cache err:", err)
		}
		go func() {
			time.Sleep(time.Millisecond)
			s.cache.Prize.ClearPrize(ctx, id)
		}()
	}
	resp.Code = errmsg.Success
	resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

// GetPrizeByAid implements the LotteryServiceImpl interface.
func (s *LotteryServiceImpl) GetPrizeByAid(ctx context.Context, req *lottery.GetPrizeByAidRequest) (resp *lottery.PrizesResponse, err error) {
	// TODO: Your code here...
	resp = new(lottery.PrizesResponse)
	aid := req.GetAid()
	var prizes []*model.Prize
	if !s.cache.Prize.ExistPrizeByAid(ctx, aid) {
		prizes, err = s.dao.Prize.GetPrizeByAid(ctx, aid)
		if err != nil {
			Log.Errorln("get prize from db err:", err)
			resp.Resp.Code = errmsg.Error
			resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
			return nil, err
		}
		err = s.cache.Prize.StorePrizeByAid(ctx, aid, prizes)
		if err != nil {
			Log.Errorln("store prize from cache err:", err)
		}
	} else {
		prizes, err = s.cache.Prize.GetPrizeByAid(ctx, aid)
		if err != nil {
			Log.Errorln("get prize from cache err:", err)
			resp.Resp.Code = errmsg.Error
			resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
			return nil, err
		}
	}
	ret_prizes := []*lottery.Prize{}
	for _, v := range prizes {
		ret_prize := &lottery.Prize{
			Id:        int32(v.ID),
			CreatedAt: v.CreatedAt.Format("2006-01-02 15:04:05"),
			Name:      v.Name,
			Num:       *v.Num,
			Picture:   *v.Picture,
		}
		ret_prizes = append(ret_prizes, ret_prize)
	}
	resp.Prizes = ret_prizes
	resp.Resp.Code = errmsg.Success
	resp.Resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

// GetPrizeById implements the LotteryServiceImpl interface.
func (s *LotteryServiceImpl) GetPrizeById(ctx context.Context, req *lottery.GetPrizeByIdRequest) (resp *lottery.PrizeResponse, err error) {
	// TODO: Your code here...
	resp = new(lottery.PrizeResponse)
	id := req.GetId()
	var prize *model.Prize
	if !s.cache.Prize.ExistPrize(ctx, id) {
		prize, err = s.dao.Prize.GetPrizeById(ctx, id)
		if err != nil {
			Log.Errorln("get prize from db,err:", err)
			resp.Resp.Code = errmsg.Error
			resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
			return nil, err
		}
		err = s.cache.Prize.StorePrize(ctx, id, prize)
		if err != nil {
			Log.Errorln("store prize into cache err:", err)
		}
	} else {
		prize, err = s.cache.Prize.GetPrize(ctx, id)
		if err != nil {
			Log.Errorln("get prize from cache err:", err)
			resp.Resp.Code = errmsg.Error
			resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
			return nil, err
		}
	}
	ret_prize := &lottery.Prize{
		Id:        int32(prize.ID),
		CreatedAt: prize.CreatedAt.Format("2006-01-02 15:04:05"),
		Name:      prize.Name,
		Num:       *prize.Num,
		Picture:   *prize.Picture,
	}
	resp.Resp.Code = errmsg.Success
	resp.Resp.Msg = errmsg.GetMsg(errmsg.Success)
	resp.Prize = ret_prize
	return
}
