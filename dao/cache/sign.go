package cache

import (
	"context"
	"sign-lottery/dao/db/model"
	"sign-lottery/pkg/constants"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type Sign struct {
}

const (
	SignPosPreffix            = "sign:location:"
	SignIpPreffix             = "sign:ip:"
	SignUserMonthPreffix      = "sign:month:"
	SignUserGroupStartPreffix = "sign:user:start:"
	SignUserGroupEndPreffix   = "sign:user:end:"
	RecordUserOffsetPreffix   = "record:offset:"
	RecordUserPreffix         = "record:"
)

func SignPosTag(id int32) string {
	return SignPosPreffix + strconv.FormatInt(int64(id), 10)
}

func SignIpTag(ip string) string {
	return SignIpPreffix + ip
}

func SignUserMonthTag(uid int64, gid int32, month string) string {
	return SignUserMonthPreffix + strconv.FormatInt(int64(gid), 10) + ":" + strconv.FormatInt(uid, 10) + ":" + month
}

func SignGroupStartTag(gid int32) string {
	return SignUserGroupStartPreffix + strconv.FormatInt(int64(gid), 10)
}

func SignGroupEndTag(gid int32) string {
	return SignUserGroupEndPreffix + strconv.FormatInt(int64(gid), 10)
}

func RecordUserTag(id int64) string {
	return RecordUserPreffix + strconv.FormatInt(id, 10)
}

func RecordUserOffsetTag(uid int64, offset int32, limit int32) string {
	return RecordUserOffsetPreffix + strconv.FormatInt(uid, 10) + ":" + strconv.FormatInt(int64(offset), 10) + ":" + strconv.FormatInt(int64(limit), 10)
}

func (s *Sign) PosLimit(ctx context.Context, id int32, latitude float64, longtitude float64) bool {
	result, err := cli.GeoRadius(ctx, SignPosTag(id), latitude, longtitude, &redis.GeoRadiusQuery{
		Radius:    constants.SignDist,
		Unit:      "km",
		WithCoord: true,
	}).Result()
	if err != nil {
		return false
	}
	return len(result) > 0
}

func (s *Sign) IpLimit(ctx context.Context, ip string, uid int64) bool {
	if cli.Exists(ctx, SignIpTag(ip)).Val() == 0 {
		return true
	}
	result, err := cli.Get(ctx, SignIpTag(ip)).Result()
	if err != nil {
		logrus.Errorln("ip relate to id err:", err)
		return false
	}
	ip_id, _ := strconv.ParseInt(result, 10, 64)
	return ip_id == uid
}

func (s *Sign) IpLimitAdd(ctx context.Context, ip string, uid int64) error {
	return cli.Set(ctx, SignIpTag(ip), uid, 24*time.Hour).Err()
}

func (s *Sign) ExistSignStart(ctx context.Context, uid int64, gid int32) bool {
	return cli.Exists(ctx, SignGroupStartTag(gid), strconv.FormatInt(uid, 10)).Val() == 1
}

func (s *Sign) UserSignStart(ctx context.Context, uid int64, gid int32) error {
	return cli.SAdd(ctx, SignGroupStartTag(gid), uid).Err()
}

func (s *Sign) UserSignEnd(ctx context.Context, uid int64, gid int32) error {
	return cli.SAdd(ctx, SignGroupEndTag(gid), uid).Err()
}

func (s *Sign) ExistPos(ctx context.Context, gid int32) bool {
	return cli.Exists(ctx, SignPosTag(gid)).Val() == 1
}

func (s *Sign) SignPosAdd(ctx context.Context, gid int32, pos *model.SignGroupPos) error {
	rpos := &redis.GeoLocation{
		Name:      pos.Name,
		Longitude: pos.Longtitude,
		Latitude:  pos.Latitude,
	}
	return cli.GeoAdd(ctx, SignPosTag(gid), rpos).Err()
}

func (s *Sign) SignPosDel(ctx context.Context, gid int32, name string) error {
	return cli.ZRem(ctx, SignPosTag(gid), name).Err()
}

func (s *Sign) ExistUserMonthRecord(ctx context.Context, uid int64, gid int32, month string) bool {
	return cli.Exists(ctx, SignUserMonthTag(uid, gid, month)).Val() == 1
}

func (s *Sign) GetUserMonthRecord(ctx context.Context, uid int64, gid int32, month string) (int, error) {
	result, err := cli.Get(ctx, SignUserMonthTag(uid, gid, month)).Result()
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(result)
}

func (s *Sign) StoreUserMonthRecord(ctx context.Context, uid int64, gid int32, bitmap int, month string) error {
	return cli.Set(ctx, SignUserMonthTag(uid, gid, month), bitmap, -1).Err()
}

func (s *Sign) UpdateUserMonthRecord(ctx context.Context, uid int64, gid int32, month string, day int) error {
	return cli.SetBit(ctx, SignUserMonthTag(uid, gid, month), int64(day), 1).Err()
}

func (s *Sign) ExistUserRecord(ctx context.Context, uid int64, limit int32, offset int32) bool {
	return cli.SIsMember(ctx, RecordUserTag(uid), RecordUserOffsetTag(uid, offset, limit)).Val()
}

func (s *Sign) StoreUserRecord(ctx context.Context, uid int64, limit int32, offset int32, records []*model.SignRecord) error {
	err := cli.SAdd(ctx, RecordUserTag(uid), RecordUserOffsetTag(uid, offset, limit)).Err()
	if err != nil {
		return err
	}
	for _, v := range records {
		err = cli.SAdd(ctx, RecordUserOffsetTag(uid, offset, limit), v.ID).Err()
		if err != nil {
			return err
		}
		if cli.Exists(ctx, RecordUserTag(v.ID)).Val() != 1 {
			err = cli.HMSet(ctx, RecordUserTag(v.ID), "gid", v.Gid, "time", v.CreatedAt.Format("2006-01-02 15:04:05")).Err()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *Sign) GetUserRecord(ctx context.Context, uid int64, limit int32, offset int32) ([]*model.SignRecord, error) {
	result, err := cli.SMembers(ctx, RecordUserOffsetTag(uid, offset, limit)).Result()
	if err != nil {
		return nil, err
	}
	ret_records := []*model.SignRecord{}
	for _, v := range result {
		record, err := cli.HGetAll(ctx, v).Result()
		if err != nil {
			return nil, err
		}
		created_at, _ := time.Parse("2006-01-02 15:04:05", record["time"])
		gid, _ := strconv.Atoi(record["gid"])
		ret_record := &model.SignRecord{
			CreatedAt: &created_at,
			Gid:       int32(gid),
		}
		ret_records = append(ret_records, ret_record)
	}
	return ret_records, err
}
