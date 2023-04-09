package db

import (
	"context"
	"fmt"
	"sign-lottery/dao/db/model"

	"gorm.io/gorm"
)

type Sign struct {
}

func (s *Sign) SignPosAdd(ctx context.Context, gPos *model.SignGroupPos) error {
	return db.WithContext(ctx).Create(gPos).Error
}

func (s *Sign) SignPosDel(ctx context.Context, gid int32, name string) error {
	return db.WithContext(ctx).Model(&model.SignGroupPos{}).Delete("gid=? AND name=?", gid, name).Error
}

func (s *Sign) SignPosGet(ctx context.Context, gid int32, offset int, limit int) (Signpos []*model.SignGroupPos, count int64, err error) {
	err = db.WithContext(ctx).Model(&model.SignGroupPos{}).Where("gid=?", gid).Count(&count).Error
	if err != nil {
		return
	}
	err = db.WithContext(ctx).Select("name,longtitude,latitude").Where("gid=?", gid).Offset((offset - 1) * limit).Limit(limit).Find(&Signpos).Error
	return
}

func (s *Sign) CheckUserGroup(ctx context.Context, uid int64, gid int32) bool {
	var count int64
	err := db.WithContext(ctx).Model(&model.UserGroup{}).Where("uid=? AND gid=?", uid, gid).Count(&count).Error
	return err == nil && count > 0
}

func (s *Sign) SignAdd(ctx context.Context, uid int64, gid int32) (err error) {
	if !s.CheckUserGroup(ctx, uid, gid) {
		return fmt.Errorf("用户未加入该组")
	}
	tx := db.Begin()
	err = tx.WithContext(ctx).Model(&model.UserGroup{}).Where("uid=? AND gid=?", uid, gid).Updates(map[string]interface{}{"score": gorm.Expr("score+10"), "count": gorm.Expr("count+1")}).Error
	if err != nil {
		tx.Rollback()
		return
	}
	err = tx.Create(&model.SignRecord{
		Gid: gid,
		UID: uid,
	}).Error
	if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}

func (s *Sign) UpdateMonthSign(ctx context.Context, uid int64, gid int32, month string, bitmap int) error {
	var err error
	var count int64
	tx := db.Begin()
	err = tx.WithContext(ctx).Model(&model.SignMonth{}).Where("uid=? AND gid=? AND month=?", uid, gid, month).Count(&count).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	if count > 0 {
		err = tx.WithContext(ctx).Model(&model.SignMonth{}).Where("uid=? AND gid=? AND month=?", uid, gid, month).Update("bitmap", bitmap).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	} else {
		err = tx.WithContext(ctx).Create(&model.SignMonth{
			Month:   month,
			Gid:     gid,
			BitSign: &bitmap,
			UID:     uid,
		}).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}

func (s *Sign) GetMonthSign(ctx context.Context, uid int64, gid int32, month string) (bitmap int, err error) {
	err = db.WithContext(ctx).Select("bit_sign").Where("uid=? AND gid=? AND month=?", uid, gid, month).First(&bitmap).Error
	return
}

func (s *Sign) AskLeave(ctx context.Context, leaveInfo *model.AskLeave) error {
	return db.WithContext(ctx).Create(&leaveInfo).Error
}

func (s *Sign) GetAllRecord(ctx context.Context, offset int, limit int) (records []*model.SignRecord, count int64, err error) {
	err = db.WithContext(ctx).Model(&model.SignRecord{}).Count(&count).Error
	if err != nil {
		return
	}
	err = db.WithContext(ctx).Model(&model.SignRecord{}).Offset((offset - 1) * limit).Limit(limit).Find(&records).Error
	return
}

func (s *Sign) GetUserRecord(ctx context.Context, uid int64, offset int, limit int) (records []*model.SignRecord, count int64, err error) {
	err = db.WithContext(ctx).Model(&model.SignRecord{}).Where("uid=?", uid).Count(&count).Error
	if err != nil {
		return
	}
	err = db.WithContext(ctx).Model(&model.SignRecord{}).Where("uid=?", uid).Offset((offset - 1) * limit).Limit(limit).Find(&records).Error
	return
}

func (s *Sign) GetMonthSignByGid(ctx context.Context, gid int32, month string, offset int, limit int) (signs []*model.SignMonth, count int64, err error) {
	err = db.WithContext(ctx).Model(&model.SignMonth{}).Where("gid=? AND month=?", gid, month).Count(&count).Error
	if err != nil {
		return
	}
	err = db.WithContext(ctx).Model(&model.SignMonth{}).Where("gid=? AND month=?", gid, month).Offset((offset - 1) * limit).Limit(limit).Find(&signs).Error
	return
}
