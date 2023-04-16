package lottery

import (
	"context"
	"sign-lottery/dao/db/model"
	"sign-lottery/kitex_gen/lottery"
	"sign-lottery/pkg/errmsg"
	. "sign-lottery/pkg/log"
	"sign-lottery/rabbitmq/consumer"
	model2 "sign-lottery/rabbitmq/model"
)

// GetUserOrder implements the LotteryServiceImpl interface.
func (s *LotteryServiceImpl) GetUserOrder(ctx context.Context, req *lottery.GetUserOrderRequest) (resp *lottery.OrdersResponse, err error) {
	// TODO: Your code here...
	resp = new(lottery.OrdersResponse)
	uid, offset, limit := req.GetUid(), req.GetOffset(), req.GetLimit()
	var orders []*model.UserOrder
	var count int64
	if !s.cache.Order.ExistOrderByUser(ctx, uid, offset, limit) {
		orders, count, err = s.dao.Order.GetUserOrder(ctx, uid, int(offset), int(limit))
		if err != nil {
			Log.Error("get user order from db err:", err)
			resp.Resp.Code = errmsg.Error
			resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
			return nil, err
		}
		err = s.cache.Order.StoreOrderByUser(ctx, uid, offset, limit, orders)
		if err != nil {
			Log.Errorln("store user order to cache err:", err)
		}
	} else {
		orders, err = s.cache.Order.GetOrdersByUser(ctx, uid, offset, limit)
		if err != nil {
			Log.Errorln("get user order from cache err:", err)
			resp.Resp.Code = errmsg.Error
			resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
			return nil, err
		}
		count = int64(len(orders))
	}
	ret_orders := []*lottery.Order{}
	for _, v := range orders {
		ret_order := &lottery.Order{
			Id:         v.ID,
			CreateTime: v.CreatedAt.Format("2006-01-02 15:04:05"),
			Uid:        v.UID,
			Pid:        int32(v.Pid),
		}
		ret_orders = append(ret_orders, ret_order)
	}
	resp.Order = ret_orders
	resp.Total = count
	resp.Resp.Code = errmsg.Success
	resp.Resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

// GetAllOrder implements the LotteryServiceImpl interface.
func (s *LotteryServiceImpl) GetAllOrder(ctx context.Context, req *lottery.GetAllOrderRequest) (resp *lottery.OrdersResponse, err error) {
	// TODO: Your code here...
	resp = new(lottery.OrdersResponse)
	offset, limit := req.GetOffset(), req.GetLimit()
	orders, count, err := s.dao.Order.GetAllOrder(ctx, int(offset), int(limit))
	if err != nil {
		resp.Resp.Code = errmsg.Error
		resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
		return nil, err
	}
	ret_orders := []*lottery.Order{}
	for _, v := range orders {
		ret_order := &lottery.Order{
			Id:         v.ID,
			CreateTime: v.CreatedAt.Format("2006-01-02 15:04:05"),
			Uid:        v.UID,
			Pid:        int32(v.Pid),
		}
		ret_orders = append(ret_orders, ret_order)
	}
	resp.Order = ret_orders
	resp.Total = count
	resp.Resp.Code = errmsg.Success
	resp.Resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

func (s *LotteryServiceImpl) HandleRabbitOrder(ctx context.Context) error {
	orderChan := make(chan model2.Order)
	err := consumer.NewConsumer().Order.ProducerOrder(orderChan)
	if err != nil {
		return err
	}
	for msg := range orderChan {
		uid := msg.Uid
		pid := msg.Pid
		uOrder := &model.UserOrder{
			UID: uid,
			Pid: int(pid),
		}
		err = s.dao.Order.OrderCreate(ctx, uOrder)
		if err != nil {
			return err
		}
	}
	return nil
}
