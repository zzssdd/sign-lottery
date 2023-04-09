package db

import (
	"context"
	"sign-lottery/dao/db/model"
)

type User struct {
}

func (u *User) Registe(ctx context.Context, user *model.User) error {
	return db.WithContext(ctx).Create(user).Error
}

func (u *User) Login(ctx context.Context, email string, password string) bool {
	var count int64
	err := db.WithContext(ctx).Model(&User{}).Where("email=? AND password=?", email, password).Count(&count).Error
	return err == nil && count > 0
}

func (u *User) GetUserById(ctx context.Context, id int64) (user *model.User, err error) {
	err = db.WithContext(ctx).First(&user, id).Error
	return
}

func (u *User) GetAllUser(ctx context.Context, offset int, limit int) (users []*model.User, count int64, err error) {
	err = db.WithContext(ctx).Model(&model.User{}).Count(&count).Error
	if err != nil {
		return
	}
	err = db.WithContext(ctx).Offset((offset - 1) * limit).Limit(limit).Find(&users).Error
	return
}

func (u *User) ChangeUserAvater(ctx context.Context, id int64, avater string) error {
	return db.WithContext(ctx).Model(&User{}).Where("id=?", id).Update("avater", avater).Error
}

func (u *User) ChangePassword(ctx context.Context, id int64, password string) error {
	return db.WithContext(ctx).Model(&User{}).Where("id=?", id).Update("password", password).Error
}

func (u *User) ChangeAddress(ctx context.Context, id int64, address string) error {
	return db.WithContext(ctx).Model(&User{}).Where("id=?", id).Update("address", address).Error
}

func (u *User) UserDel(ctx context.Context, id int64) error {
	return db.WithContext(ctx).Delete(&User{}, id).Error
}

func (u *User) GetUserByGid(ctx context.Context, gid int32, offset int, limit int) (users []*model.User, err error) {
	err = db.WithContext(ctx).Model(&model.User{}).Select("user.id,user.email,user.name,user.avater").Joins("left join user_group on user_group.uid=user.id").Where("user_group.gid=?", gid).Offset((offset - 1) * limit).Limit(limit).Find(&users).Error
	return
}
