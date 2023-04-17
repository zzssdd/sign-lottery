package cache

import (
	"context"
	"sign-lottery/dao/db/model"
	"strconv"
	"time"
)

type User struct {
}

const (
	UserPreffix  = "user:"
	LoginPreffix = "login:"
	EmailPreffix = "email:"
)

func UserInfoTag(id int64) string {
	return UserPreffix + strconv.FormatInt(id, 10)
	//user:id
}

func LoginTag(email string) string {
	return LoginPreffix + email
}

func EmailTag(email string) string {
	return EmailPreffix + email
}

func (u *User) StoreCode(ctx context.Context, email string, code string) error {
	return cli.Set(ctx, EmailTag(email), code, time.Minute).Err()
}

func (u *User) ExistCode(ctx context.Context, email string) bool {
	return cli.Exists(ctx, EmailTag(email)).Val() == 1
}

func (u *User) GetCode(ctx context.Context, email string) (code string, err error) {
	code, err = cli.Get(ctx, EmailTag(email)).Result()
	return
}

func (u *User) UserInfoExist(ctx context.Context, id int64) bool {
	return cli.Exists(ctx, UserInfoTag(id)).Val() == 1
}

func (u *User) GetUserById(ctx context.Context, id int64) (*model.User, error) {
	result, err := cli.HGetAll(ctx, UserInfoTag(id)).Result()
	if err != nil {
		return nil, err
	}
	creat_time, _ := time.Parse("2006-01-02 15:04:05", result["created_at"])
	avater := result["avater"]
	user := &model.User{
		ID:        id,
		CreatedAt: &creat_time,
		Email:     result["email"],
		Name:      result["name"],
		Avater:    &avater,
	}
	return user, nil
}

func (u *User) GetUsersById(ctx context.Context, ids []string) ([]*model.User, error) {
	var users []*model.User
	for _, v := range ids {
		result, err := cli.HGetAll(ctx, UserPreffix+v).Result()
		if err != nil {
			return nil, err
		}
		creat_time, _ := time.Parse("2006-01-02 15:04:05", result["created_at"])
		id, _ := strconv.ParseInt(v, 10, 64)
		avater := result["avater"]
		user := &model.User{
			ID:        id,
			CreatedAt: &creat_time,
			Email:     result["email"],
			Name:      result["name"],
			Avater:    &avater,
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *User) StoreUserInfo(ctx context.Context, id int64, info *model.User) error {
	v := map[string]interface{}{
		"create_at": info.CreatedAt.Format("2006-01-02 15:04:05"),
		"email":     info.Email,
		"name":      info.Name,
		"avater":    info.Avater,
	}
	err := cli.HMSet(ctx, UserInfoTag(id), v).Err()
	if err != nil {
		return err
	}
	err = cli.Expire(ctx, UserInfoTag(id), time.Hour).Err()
	if err != nil {
		cli.Del(ctx, UserInfoTag(id))
		return err
	}
	return nil
}

func (u *User) StoreUsersInfo(ctx context.Context, users []*model.User) (err error) {
	for _, v := range users {
		value := map[string]interface{}{
			"create_at": v.CreatedAt.Format("2006-01-02 15:04:05"),
			"email":     v.Email,
			"name":      v.Name,
			"avater":    v.Avater,
		}
		err = cli.HMSet(ctx, UserInfoTag(v.ID), value).Err()
		if err != nil {
			break
		}
		err = cli.Expire(ctx, UserInfoTag(v.ID), time.Hour).Err()
		if err != nil {
			cli.Del(ctx, UserInfoTag(v.ID))
			break
		}
	}
	return err
}

func (u *User) ClearUserInfo(ctx context.Context, id int64) error {
	return cli.Del(ctx, UserInfoTag(id)).Err()
}

func (u *User) ExistsLoginInfo(ctx context.Context, email string) bool {
	return cli.Exists(ctx, LoginTag(email)).Val() == 1
}

func (u *User) GetLoginInfo(ctx context.Context, email string) map[string]string {
	return cli.HGetAll(ctx, LoginTag(email)).Val()
}

func (u *User) StoreLoginInfo(ctx context.Context, id int64, email string, password string) error {
	err := cli.HMSet(ctx, LoginTag(email), "id", id, "password", password).Err()
	if err != nil {
		return err
	}
	err = cli.Expire(ctx, LoginTag(email), time.Hour).Err()
	if err != nil {
		cli.Del(ctx, LoginTag(email))
		return err
	}
	return nil
}

func (u *User) ClearLoginInfo(ctx context.Context, email string) error {
	return cli.Del(ctx, LoginTag(email)).Err()
}
