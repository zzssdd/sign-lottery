package cache

import (
	"context"
	"sign-lottery/dao/db/model"
	"strconv"
	"time"
)

type Group struct {
}

const (
	GroupOffsetPreffix = "group:offset:"
	GroupInfoPreffix   = "group:"
)

func GroupOffsetTag(id int32, offset int32, limit int32) string {
	return GroupOffsetPreffix + strconv.FormatInt(int64(id), 10) + ":" + strconv.FormatInt(int64((offset)), 10) + ":" + strconv.FormatInt(int64(limit), 10)
	//group:id:(offset-1)*limit:limit
}

func GroupInfoTag(id int32) string {
	return GroupInfoPreffix + strconv.Itoa(int(id))
}

func (g *Group) GetUserByGid(ctx context.Context, gid int32, offset int32, limit int32) ([]string, int64, error) {
	group_tag := GroupOffsetTag(gid, offset, limit)
	result, err := cli.SMembers(ctx, group_tag).Result()
	if err != nil {
		return nil, 0, err
	}
	count, err := cli.HGet(ctx, GroupInfoTag(gid), "count").Result()
	if err != nil {
		return nil, 0, err
	}
	Scount, err := strconv.ParseInt(count, 10, 64)
	if err != nil {
		return nil, 0, err
	}
	return result, Scount, nil
}

func (g *Group) GroupInfoExist(ctx context.Context, id int32) bool {
	return cli.Exists(ctx, GroupInfoTag(id)).Val() == 1
}

func (g *Group) StoreGroupInfo(ctx context.Context, groupInfo *model.SignGroup) error {
	return cli.HMSet(ctx, GroupInfoTag(int32(groupInfo.ID)), "name", groupInfo.Name, "created_at", groupInfo.CreatedAt, "start", groupInfo.Start, "end", groupInfo.End, "count", groupInfo.Count, "avater", groupInfo.Avater, "owner", groupInfo.Owner).Err()
}

func (g *Group) ClearGroupInfo(ctx context.Context, id int32) error {
	return cli.Del(ctx, GroupInfoTag(id)).Err()
}

func (g *Group) GetGroupInfo(ctx context.Context, id int32) (*model.SignGroup, error) {
	v, err := cli.HGetAll(ctx, GroupInfoTag(id)).Result()
	if err != nil {
		return nil, err
	}
	count, _ := strconv.ParseInt(v["count"], 10, 64)
	oid, _ := strconv.ParseInt(v["owner"], 10, 64)
	start, _ := time.Parse("2006-01-02 15:04:05", v["start"])
	end, _ := time.Parse("2006-01-02 15:04:05", v["end"])
	create, _ := time.Parse("2006-01-02 15:04:05", v["created_at"])
	avater := v["avater"]
	groupInfo := &model.SignGroup{
		ID:        int(id),
		Name:      v["name"],
		Start:     start,
		End:       end,
		Count:     &count,
		Avater:    &avater,
		Owner:     oid,
		CreatedAt: &create,
	}
	return groupInfo, nil
}

func (g *Group) StoreGroupOffset(ctx context.Context, gid, offset, limit int32, groupInfo *model.SignGroup) (err error) {
	script := `if redis.call('exists',KEYS[1])==1 then return redis.call('sadd',KEYS[10],ARGV[9]) else return redis.call('hmset',KEYS[1],KEYS[2],ARGV[1],KEYS[3],ARGV[2],
KEYS[4],ARGV[3],KEYS[5],ARGV[4],KEYS[6],ARGV[5],KEYS[7],ARGV[6],KEYS[8],ARGV[7],KEYS[9],ARGV[8]) and redis.call('sadd',KEYS[10],ARGV[9]) and redis.call('expire',KEYS[1],1000) end`
	err = cli.Eval(ctx, script, []string{GroupInfoTag(gid), "id", "name", "created_at", "start", "end", "count", "avater", "owner", GroupOffsetPreffix},
		gid, groupInfo.Name, groupInfo.CreatedAt, groupInfo.Start, groupInfo.End, groupInfo.Count, groupInfo.Avater, groupInfo.Owner, GroupOffsetTag(gid, offset, limit)).Err()
	return err
}

func (g *Group) GroupOffsetExist(ctx context.Context, gid int32, offset int32, limit int32) bool {
	return cli.Exists(ctx, GroupOffsetTag(gid, offset, limit)).Val() == 1
}
