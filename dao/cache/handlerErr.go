package cache

import (
	"context"
	"strconv"
	"time"
)

type HandlerErr struct {
}

const (
	EmailErrPreffix  = "email:err:"
	SignErrPreffix   = "sign:err:"
	ChooseErrPreffix = "choose:err:"
)

func EmailErrTag(email string) string {
	return EmailErrPreffix + email
}

func SignErrTag(uid int64, gid int32) string {
	return SignErrPreffix + strconv.FormatInt(uid, 10) + strconv.Itoa(int(gid))
}

func ChooseErrTag(uid int64, gid int32) string {
	return ChooseErrPreffix + strconv.FormatInt(uid, 10) + strconv.Itoa(int(gid))
}

func (r *HandlerErr) ExistEmailErr(ctx context.Context, email string) bool {
	return cli.Exists(ctx, EmailErrTag(email)).Val() == 1
}

func (r *HandlerErr) ExistSignErr(ctx context.Context, uid int64, gid int32) bool {
	return cli.Exists(ctx, SignErrTag(uid, gid)).Val() == 1
}

func (r *HandlerErr) ExistChooseErr(ctx context.Context, uid int64, gid int32) bool {
	return cli.Exists(ctx, ChooseErrTag(uid, gid)).Val() == 1
}

func (r *HandlerErr) ReturnEmailErr(ctx context.Context, email string, code int) error {
	return cli.Set(ctx, EmailErrTag(email), code, time.Hour).Err()
}

func (r *HandlerErr) ReturnSignErr(ctx context.Context, uid int64, gid int32, code int) error {
	return cli.Set(ctx, SignErrTag(uid, gid), code, time.Hour).Err()
}

func (r *HandlerErr) ReturnChooseErr(ctx context.Context, uid int64, aid int32, prizeName string) error {
	return cli.Set(ctx, ChooseErrTag(uid, aid), prizeName, time.Hour).Err()
}

func (r *HandlerErr) GetEmailErr(ctx context.Context, email string) (int, error) {
	result, err := cli.Get(ctx, EmailErrTag(email)).Result()
	if err != nil {
		return 0, err
	}
	err = cli.Del(ctx, EmailErrTag(email)).Err()
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(result)
}

func (r *HandlerErr) GetSignErr(ctx context.Context, uid int64, gid int32) (int, error) {
	result, err := cli.Get(ctx, SignErrTag(uid, gid)).Result()
	if err != nil {
		return 0, err
	}
	err = cli.Del(ctx, SignErrTag(uid, gid)).Err()
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(result)
}
func (r *HandlerErr) GetChooseErr(ctx context.Context, uid int64, aid int32) (string, error) {
	result, err := cli.Get(ctx, ChooseErrTag(uid, aid)).Result()
	if err != nil {
		return "", err
	}
	err = cli.Del(ctx, ChooseErrTag(uid, aid)).Err()
	if err != nil {
		return "", err
	}
	return result, nil
}
