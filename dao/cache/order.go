package cache

import (
	"context"
	"sign-lottery/dao/db/model"
	"strconv"
	"time"
)

type Order struct {
}

const (
	OrderByUserPreffix = "order:user:"
	OrderAllPreffix    = "order:all:"
	OrderPreffix       = "order:"
)

func OrderByUserTag(id int64, offset int32, limit int32) string {
	return OrderByUserPreffix + strconv.FormatInt(id, 10) + ":" + strconv.Itoa(int(offset)) + ":" + strconv.Itoa(int(limit))
}

func OrderAllTag(offset int32, limit int32) string {
	return OrderAllPreffix + strconv.Itoa(int(offset)) + ":" + strconv.Itoa(int(limit))
}

func OrderTag(id int64) string {
	return OrderPreffix + strconv.FormatInt(id, 10)
}

func (o *Order) ExistOrderByUser(ctx context.Context, id int64, offset int32, limit int32) bool {
	return cli.Exists(ctx, OrderByUserTag(id, offset, limit)).Val() == 1
}

func (o *Order) StoreOrder(ctx context.Context, id int64, order *model.UserOrder) error {
	return cli.HMSet(ctx, OrderTag(id), "id", order.ID, "created_at", order.CreatedAt.Format("2006-01-02 15:04:05"), "uid", order.UID, "pid", order.Pid).Err()
}

func (o *Order) GetOrderById(ctx context.Context, id int64) (*model.UserOrder, error) {
	result, err := cli.HGetAll(ctx, OrderTag(id)).Result()
	if err != nil {
		return nil, err
	}
	oid, _ := strconv.ParseInt(result["id"], 10, 64)
	created_at, _ := time.Parse("2006-01-02 15:04:05", result["created_at"])
	uid, _ := strconv.ParseInt(result["uid"], 10, 64)
	pid, _ := strconv.Atoi(result["pid"])
	order := &model.UserOrder{
		ID:        oid,
		CreatedAt: &created_at,
		UID:       uid,
		Pid:       pid,
	}
	return order, nil
}

func (o *Order) ExistOrder(ctx context.Context, id int64) bool {
	return cli.Exists(ctx, OrderTag(id)).Val() == 1
}

func (o *Order) StoreOrderByUser(ctx context.Context, id int64, offset int32, limit int32, orders []*model.UserOrder) error {
	for _, v := range orders {
		err := cli.SAdd(ctx, OrderByUserTag(id, offset, limit), v.ID).Err()
		if err != nil {
			return err
		}
		if !o.ExistOrder(ctx, v.ID) {
			err = o.StoreOrder(ctx, v.ID, v)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (o *Order) GetOrdersByUser(ctx context.Context, id int64, offset int32, limit int32) ([]*model.UserOrder, error) {
	m, err := cli.SMembers(ctx, OrderByUserTag(id, offset, limit)).Result()
	if err != nil {
		return nil, err
	}
	ret_orders := []*model.UserOrder{}
	for _, v := range m {
		oid, _ := strconv.ParseInt(v, 10, 64)
		order, err := o.GetOrderById(ctx, oid)
		if err != nil {
			return nil, err
		}
		ret_orders = append(ret_orders, order)
	}
	return ret_orders, nil
}
