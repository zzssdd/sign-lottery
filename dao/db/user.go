package db

import (
	"context"
	"sign-lottery/dao/db/model"
)

type User struct {
}

func (u *User) Registe(ctx context.Context, user *model.User) (error, int64) {
	return db.WithContext(ctx).Create(user).Error, user.ID
}

func (u *User) GetPasswordAndIdByEmail(ctx context.Context, email string) (*model.User, error) {
	var user *model.User
	err := db.WithContext(ctx).Model(&model.User{}).Select("id,password").Where("email=?", email).First(&user).Error
	return user, err
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

func (u *User) ChangePassword(ctx context.Context, id int64, oldPassword string, newPassword string) (email *string, err error) {
	err = db.WithContext(ctx).Model(&User{}).Select("email").Where("id=? AND password=?", id, oldPassword).First(&email).Error
	if email == nil {
		return
	}
	return email, db.WithContext(ctx).Model(&User{}).Where("id=?", id).Update("password", newPassword).Error
}

func (u *User) ChangeAddress(ctx context.Context, id int64, address string) error {
	return db.WithContext(ctx).Model(&User{}).Where("id=?", id).Update("address", address).Error
}

func (u *User) UserDel(ctx context.Context, id int64) error {
	return db.WithContext(ctx).Delete(&User{}, id).Error
}

func (u *User) GetUserByGid(ctx context.Context, gid int32, offset int, limit int) (users []*model.User, count int64, err error) {
	err = db.WithContext(ctx).Model(&model.UserGroup{}).Where("gid=?", gid).Count(&count).Error
	if err != nil {
		return
	}
	err = db.WithContext(ctx).Model(&model.User{}).Select("user.id,user.email,user.name,user.avater").Joins("left join user_group on user_group.uid=user.id").Where("user_group.gid=?", gid).Offset((offset - 1) * limit).Limit(limit).Find(&users).Error
	return
}

func (u *User) CheckUserIsExist(ctx context.Context, email string) (int64, bool) {
	var count int64
	var id int64
	err := db.WithContext(ctx).Model(&model.User{}).Select("id").Where("email=?", email).First(&id).Count(&count).Error
	return id, err != nil || count > 0
}
