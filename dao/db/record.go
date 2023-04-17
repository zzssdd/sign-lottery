package db

import (
	"context"
	"gorm.io/gorm"
	"sign-lottery/dao/db/model"
	. "sign-lottery/pkg/log"
)

type Record struct {
}

func (r *Record) DoRecord(ctx context.Context, uid int64, gid int32, month string, bitmap int) error {
	tx := db.Begin()
	err := tx.WithContext(ctx).Create(&model.SignRecord{UID: uid, Gid: gid}).Error
	if err != nil {
		Log.Errorln("create sign record err:", err)
		tx.Rollback()
		return err
	}
	err = tx.WithContext(ctx).Model(&model.SignMonth{}).Where("uid=? AND gid=? AND month=?", uid, gid, month).Update("bitsign", bitmap).Error
	if err != nil {
		Log.Errorln("update month bitsign err:", err)
		tx.Rollback()
		return err
	}
	err = tx.WithContext(ctx).Model(&model.UserGroup{}).Where("uid=? AND gid=?", uid, gid).Updates(map[string]interface{}{"count": gorm.Expr("count+1"), "score": gorm.Expr("score+10")}).Error
	if err != nil {
		Log.Errorln("update sign score err:", err)
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
