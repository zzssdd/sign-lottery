package db

import (
	"context"
	"sign-lottery/dao/db/model"

	"gorm.io/gorm"
)

type Group struct {
}

func (g *Group) CreateGroup(ctx context.Context, group *model.SignGroup) error {
	return db.WithContext(ctx).Create(&group).Error
}

func (g *Group) GetAllGroup(ctx context.Context, offset int, limit int) (groups []*model.SignGroup, count int64, err error) {
	err = db.WithContext(ctx).Model(&model.SignGroup{}).Count(&count).Error
	if err != nil {
		return
	}
	err = db.WithContext(ctx).Offset((offset - 1) * limit).Limit(limit).Find(&groups).Error
	return
}

func (g *Group) GetGroupById(ctx context.Context, id int32) (group *model.SignGroup, err error) {
	err = db.WithContext(ctx).First(&group, id).Error
	return
}

func (g *Group) GroupUpdate(ctx context.Context, id int32, group *model.SignGroup) error {
	return db.WithContext(ctx).Model(&model.SignGroup{}).Where("id=?", id).Updates(group).Error
}

func (g *Group) CheckGroupPrevilege(ctx context.Context, uid int64, gid int32) bool {
	var count int64
	err := db.WithContext(ctx).Model(&model.SignGroup{}).Where("id=? AND owner=?", gid, uid).Count(&count).Error
	return count > 0 && err == nil
}

func (g *Group) GroupDel(ctx context.Context, id int32) error {
	return db.WithContext(ctx).Delete(&model.SignGroup{}, id).Error
}

func (g *Group) JoinGroup(ctx context.Context, uid int64, gid int32) error {
	var err error
	user_group := &model.UserGroup{
		UID: uid,
		Gid: gid,
	}
	tx := db.Begin()
	err = tx.WithContext(ctx).Create(&user_group).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.WithContext(ctx).Model(&model.SignGroup{}).Where("id=?", gid).Update("count", gorm.Expr("count+1")).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
