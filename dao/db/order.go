package db

import (
	"context"
	"sign-lottery/dao/db/model"
)

type Order struct {
}

func (o *Order) GetUserOrder(ctx context.Context, uid int64, offset int, limit int) (orders []*model.UserOrder, count int64, err error) {
	err = db.WithContext(ctx).Model(&orders).Where("uid=?", uid).Count(&count).Error
	if err != nil {
		return
	}
	err = db.WithContext(ctx).Where("uid=?", uid).Offset((offset - 1) * limit).Limit(limit).Find(&orders).Error
	return
}

func (o *Order) GetAllOrder(ctx context.Context, offset int, limit int) (orders []*model.UserOrder, count int64, err error) {
	err = db.WithContext(ctx).Model(&orders).Count(&count).Error
	if err != nil {
		return
	}
	err = db.WithContext(ctx).Offset((offset - 1) * limit).Limit(limit).Find(&orders).Error
	return
}

func (o *Order) OrderCreate(ctx context.Context, uOrder *model.UserOrder) error {
	return db.WithContext(ctx).Create(&uOrder).Error
}
