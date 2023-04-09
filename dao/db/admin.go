package db

import "context"

type Admin struct {
}

func (a *Admin) Login(ctx context.Context, name string, password string) bool {
	var count int64
	err := db.WithContext(ctx).Model(&Admin{}).Where("name=? AND password=?", name, password).Count(&count).Error
	return err == nil && count > 0
}
