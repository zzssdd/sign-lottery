package lottery

import (
	"context"
	"math/rand"
	"sign-lottery/dao/db/model"
	"sign-lottery/kitex_gen/lottery"
	"sign-lottery/pkg/errmsg"
	. "sign-lottery/pkg/log"
	"sign-lottery/rabbitmq/consumer"
	model2 "sign-lottery/rabbitmq/model"
	"sign-lottery/rabbitmq/producer"
	"time"
)

// Choose implements the LotteryServiceImpl interface.
func (s *LotteryServiceImpl) Choose(ctx context.Context) (resp *lottery.ChooseResponse, err error) {
	// TODO: Your code here...
	chooseChan := make(chan model2.Choose)
	err = consumer.NewConsumer().Choose.ConsumerChoose(chooseChan)
	if err != nil {
		Log.Fatalln("create consumer err:", err)
	}
	var count int64
	producer := producer.NewProcuer()
	for chooseInfo := range chooseChan {
		uid := chooseInfo.Uid
		aid := chooseInfo.Aid
		if !s.cache.Prize.ExistPrizeByAid(ctx, aid) {
			prizes, err := s.dao.Prize.GetPrizeByAid(ctx, aid)
			if err != nil {
				Log.Errorln("get prizes from db err:", err)
				err := s.cache.Activity.IncrActivityNum(ctx, aid, 1)
				if err != nil {
					Log.Errorln("incr activity num err:", err)
				}
				continue
			}
			s.cache.Prize.StorePrizeByAid(ctx, aid, prizes)
		}
		prizes, err := s.cache.Prize.GetPrizeByAid(ctx, aid)
		if err != nil {
			Log.Errorln("get prize by aid err:", err)
			err := s.cache.Activity.IncrActivityNum(ctx, aid, 1)
			if err != nil {
				Log.Errorln("incr activity num err:", err)
			}
			err = s.cache.HandlerErr.ReturnChooseErr(ctx, uid, aid, errmsg.Error)
			if err != nil {
				Log.Errorln("store return choose code err:", err)
				err := s.cache.Activity.IncrActivityNum(ctx, aid, 1)
				if err != nil {
					Log.Errorln("incr activity num err:", err)
				}
			}
			continue
		}
		for _, v := range prizes {
			count += *v.Num
		}
		rand.Seed(time.Now().Unix())
		rand_num := rand.Int63n(count) + 1
		for _, v := range prizes {
			rand_num -= *v.Num
			if rand_num <= 0 {
				err = s.cache.Prize.IncrPrize(ctx, int32(v.ID))
				if err != nil {
					err = s.cache.HandlerErr.ReturnChooseErr(ctx, uid, aid, errmsg.ChooseNoPrize)
					if err != nil {
						Log.Errorln("store return choose code err:", err)
					}
					continue
				}
				order := &model.UserOrder{
					UID: uid,
					Pid: v.ID,
				}
				err = s.dao.Order.OrderCreate(ctx, order)
				if err != nil {
					msg := &model2.Order{
						Uid: order.UID,
						Pid: int32(order.Pid),
					}
					err = producer.Order.ProducerOrder(msg)
					if err != nil {
						Log.Errorln("producer order err:", err)
					}
					err := s.cache.Activity.IncrActivityNum(ctx, aid, 1)
					if err != nil {
						Log.Errorln("incr activity num err:", err)
					}
					err = s.cache.Activity.IncrActivityNum(ctx, aid, 1)
					if err != nil {
						Log.Errorln("incr activity num err:", err)
					}
					continue
				}
				err = s.cache.HandlerErr.ReturnSignErr(ctx, uid, aid, errmsg.Success)
				if err != nil {
					Log.Errorln("store return choose code err:", err)
					err = s.cache.Activity.IncrActivityNum(ctx, aid, 1)
					if err != nil {
						Log.Errorln("incr activity num err:", err)
					}
				}
				continue
			}
		}
	}
	return
}
