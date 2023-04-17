package db

import (
	"context"
	"sign-lottery/dao/db/model"

	"gorm.io/gorm"
)

type Prize struct {
}

func (p *Prize) PrizeAdd(ctx context.Context, prize *model.Prize) error {
	tx := db.Begin()
	err := tx.WithContext(ctx).Create(&prize).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.WithContext(ctx).Model(&Activity{}).Where("id=?", prize.Aid).Updates(map[string]interface{}{"num": gorm.Expr("num+?", prize.Num), "count": gorm.Expr("count+1")}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (p *Prize) GetPrizeAid(ctx context.Context, id int32) (int32, error) {
	var aid int32
	err := db.WithContext(ctx).Select("aid").Where("id=?", id).First(&aid).Error
	return aid, err
}

func (p *Prize) PrizeDel(ctx context.Context, id int32) error {
	var prize *model.Prize
	tx := db.Begin()
	err := tx.WithContext(ctx).Where("id=?", id).First(&prize).Error
	if err != nil {
		return err
	}
	err = tx.WithContext(ctx).Delete(&model.Prize{}, id).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.WithContext(ctx).Model(&Activity{}).Where("id=?", prize.Aid).Updates(map[string]interface{}{"num": gorm.Expr("num-?", prize.Num), "count": gorm.Expr("count-1")}).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (p *Prize) PrizeUpdate(ctx context.Context, id int32, prize *model.Prize) error {
	var Before *model.Prize
	tx := db.Begin()
	err := tx.Where(ctx).Where("id=?", id).First(&Before).Error
	if err != nil {
		return err
	}
	err = tx.WithContext(ctx).Model(&model.Prize{}).Where("id=?", id).Updates(prize).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.WithContext(ctx).Model(&model.Activity{}).Where("aid=?", prize.Aid).Update("num", gorm.Expr("num+?", *prize.Num-*Before.Num)).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (p *Prize) GetPrizeByAid(ctx context.Context, aid int32) (prizes []*model.Prize, err error) {
	err = db.WithContext(ctx).Where("aid=?", aid).Find(&prizes).Error
	return
}

func (p *Prize) GetPrizeById(ctx context.Context, id int32) (prize *model.Prize, err error) {
	err = db.Find(&prize, id).Error
	return
}
