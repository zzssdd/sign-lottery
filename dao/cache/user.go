package cache

import (
	"context"
	"sign-lottery/dao/db/model"
	"strconv"
	"time"
)

type User struct {
}

func UserInfoTag(id int64) string {
	return "user:" + strconv.FormatInt(id, 10)
}

func (u *User) StoreCode(ctx context.Context, email string, code string) error {
	return cli.Set(ctx, email, code, time.Minute).Err()
}

func (u *User) GetCode(ctx context.Context, email string) (code string, err error) {
	code, err = cli.Get(ctx, email).Result()
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

func (u *User) StoreUserInfo(ctx context.Context, id int64, info *model.User) error {
	v := map[string]interface{}{
		"create_at": info.CreatedAt.Format("2006-01-02 15:04:05"),
		"email":     info.Email,
		"name":      info.Name,
		"avater":    info.Avater,
	}
	return cli.HSet(ctx, UserInfoTag(id), v).Err()
}

func (u *User) GetUserByGid(ctx context.Context, gid int32) {

}
