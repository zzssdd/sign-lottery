package db

import (
	"context"
	"sign-lottery/dao/db/model"
)

type Activity struct {
}

func (a *Activity) ActivityAdd(ctx context.Context, activity *model.Activity) (error, int) {
	return db.WithContext(ctx).Create(&activity).Error, activity.ID
}

func (a *Activity) ActivityDel(ctx context.Context, id int32) error {
	tx := db.Begin()
	err := tx.WithContext(ctx).Delete(&model.Activity{}, id).Error
	if err != nil {
		return err
	}
	err = tx.WithContext(ctx).Where("aid=?", id).Delete(&model.Prize{}).Error
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (a *Activity) CheckActivityPrevilege(ctx context.Context, uid int64, aid int32) bool {
	var count int64
	err := db.WithContext(ctx).Model(&model.Activity{}).Where("uid=? AND id=?", uid, aid).Count(&count).Error
	return count > 0 && err == nil
}

func (a *Activity) ActivityUpdate(ctx context.Context, id int32, activity *model.Activity) error {
	return db.WithContext(ctx).Where("id=?", id).Updates(activity).Error
}

func (a *Activity) GetActivityByGid(ctx context.Context, gid int32, offset int, limit int) (activitys []*model.Activity, count int64, err error) {
	err = db.WithContext(ctx).Model(&model.Activity{}).Where("gid=?", gid).Count(&count).Error
	if err != nil {
		return
	}
	err = db.WithContext(ctx).Where("gid=?", gid).Offset((offset - 1) * limit).Limit(limit).Find(activitys).Error
	return
}

func (a *Activity) GetAllActivity(ctx context.Context, offset int, limit int) (activitys []*model.Activity, count int64, err error) {
	err = db.WithContext(ctx).Model(&model.Activity{}).Count(&count).Error
	if err != nil {
		return
	}
	err = db.WithContext(ctx).Offset((offset - 1) * limit).Limit(limit).Find(activitys).Error
	return
}

func (a *Activity) GetActivityById(ctx context.Context, id int32) (activity *model.Activity, err error) {
	err = db.WithContext(ctx).First(&activity, id).Error
	return
}
