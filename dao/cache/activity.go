package cache

import (
	"context"
	"sign-lottery/dao/db/model"
	"strconv"
	"time"
)

type Activity struct {
}

const (
	ActivityInfoPreffix        = "activity:"
	ActivityOffsetByGidPreffix = "activityOffset:"
	ActivityGroupPreffix       = "activity:group:"
)

func ActivityInfoTag(id int32) string {
	return ActivityInfoPreffix + strconv.Itoa(int(id))
}

func GidOffsetTag(gid int32, offset int32, limit int32) string {
	return ActivityGroupPreffix + strconv.Itoa(int(gid)) + ":" + strconv.Itoa(int(offset)) + ":" + strconv.Itoa(int(limit))
}

func ActivityOffsetTag(gid int32) string {
	return ActivityOffsetByGidPreffix + strconv.Itoa(int(gid))
}

func (a *Activity) ExistActivityInfo(ctx context.Context, id int32) bool {
	return cli.Exists(ctx, ActivityInfoTag(id)).Val() == 1
}

func (a *Activity) StoreActivityInfo(ctx context.Context, id int32, activity *model.Activity) error {
	return cli.HMSet(ctx, ActivityInfoTag(id), "create_at", activity.CreatedAt.Format("2006-01-02 15:04:05"), "name", activity.Name, "desc", activity.Des, "picture", activity.Picture, "cost", activity.Cost,
		"uid", activity.UID, "gid", activity.Gid, "start", activity.Start, "end", activity.End, "num", activity.Num).Err()
}

func (a *Activity) ClearActivityInfo(ctx context.Context, id int32) error {
	return cli.Del(ctx, ActivityInfoTag(id)).Err()
}

func (a *Activity) ExistActivityOffset(ctx context.Context, id int32, offset int32, limit int32) bool {
	return cli.SIsMember(ctx, ActivityOffsetTag(id), GidOffsetTag(id, offset, limit)).Val()
}

func (a *Activity) StoreActivityOffset(ctx context.Context, id int32, offset int32, limit int32, acticitys []*model.Activity) error {
	err := cli.SAdd(ctx, ActivityOffsetTag(id), GroupOffsetTag(id, offset, limit)).Err()
	if err != nil {
		return err
	}
	for _, v := range acticitys {
		if !a.ExistActivityInfo(ctx, int32(v.ID)) {
			err = a.StoreActivityInfo(ctx, int32(v.ID), v)
			if err != nil {
				return err
			}
			err = cli.SAdd(ctx, GidOffsetTag(id, offset, limit), v.ID).Err()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (a *Activity) GetActivityById(ctx context.Context, id int32) (*model.Activity, error) {
	result, err := cli.HGetAll(ctx, ActivityInfoTag(id)).Result()
	if err != nil {
		return nil, err
	}
	create_at, _ := time.Parse("2006-01-02 15:04:05", result["create_at"])
	start, _ := time.Parse("2006-01-02 15:04:05", result["start"])
	end, _ := time.Parse("2006-01-02 15:04:05", result["end"])
	desc := result["desc"]
	picture := result["picture"]
	tmpcost, _ := strconv.Atoi(result["cost"])
	cost := int32(tmpcost)
	uid, _ := strconv.ParseInt(result["uid"], 10, 64)
	gid, _ := strconv.Atoi(result["gid"])
	num, _ := strconv.ParseInt(result["num"], 10, 64)
	activity := &model.Activity{
		ID:        int(id),
		CreatedAt: &create_at,
		Name:      result["name"],
		Des:       &desc,
		Picture:   &picture,
		Cost:      &cost,
		UID:       &uid,
		Gid:       gid,
		Start:     &start,
		End:       &end,
		Num:       &num,
	}
	return activity, nil
}

func (a *Activity) GetActivityOffset(ctx context.Context, id int32, offset int32, limit int32) ([]*model.Activity, error) {
	result, err := cli.SMembers(ctx, GidOffsetTag(id, offset, limit)).Result()
	if err != nil {
		return nil, err
	}
	var activitys []*model.Activity
	for _, v := range result {
		id, _ := strconv.Atoi(v)
		activity, err := a.GetActivityById(ctx, int32(id))
		if err != nil {
			return nil, err
		}
		activitys = append(activitys, activity)
	}
	return activitys, nil
}

func (a *Activity) IncrActivityNum(ctx context.Context, id int32, base int64) error {
	return cli.HIncrBy(ctx, ActivityInfoTag(id), "num", base).Err()
}

func (a *Activity) CheckActivityNum(ctx context.Context, id int32) bool {
	result, err := cli.HGetAll(ctx, ActivityInfoTag(id)).Result()
	if err != nil {
		return false
	}
	num, err := strconv.Atoi(result["num"])
	if num <= 0 && err != nil {
		return false
	}
	return true
}

func (a *Activity) CheckActivityTime(ctx context.Context, id int32, t time.Time) bool {
	result, err := cli.HGetAll(ctx, ActivityInfoTag(id)).Result()
	if err != nil {
		return false
	}
	start, err := time.Parse("2006-01-02 15:04:05", result["start"])
	if err != nil {
		return false
	}
	end, err := time.Parse("2006-01-02 15:04:05", result["end"])
	if err != nil {
		return false
	}
	return !(t.Before(start) || t.After(end))
}
