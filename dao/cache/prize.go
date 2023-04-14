package cache

import (
	"context"
	"sign-lottery/dao/db/model"
	"strconv"
	"time"
)

type Prize struct {
}

const (
	PrizeInfoPreffix  = "prize:"
	PrizeByAidPreffix = "prize:aid:"
)

func PrizeTag(id int32) string {
	return PrizeInfoPreffix + strconv.Itoa(int(id))
}

func PrizeByAidTag(id int32) string {
	return PrizeByAidPreffix + strconv.Itoa(int(id))
}

func (p *Prize) ExistPrize(ctx context.Context, id int32) bool {
	return cli.Exists(ctx, PrizeTag(id)).Val() == 1
}

func (p *Prize) ClearPrize(ctx context.Context, id int32) error {
	return cli.Del(ctx, PrizeTag(id)).Err()
}

func (p *Prize) GetPrize(ctx context.Context, id int32) (*model.Prize, error) {
	result, err := cli.HGetAll(ctx, PrizeTag(id)).Result()
	if err != nil {
		return nil, err
	}
	created_at, _ := time.Parse("2006-01-02 15:04:05", result["created_at"])
	num, _ := strconv.ParseInt(result["num"], 10, 64)
	picture := result["picture"]
	aid, _ := strconv.Atoi(result["aid"])
	prize := &model.Prize{
		ID:        int(id),
		CreatedAt: &created_at,
		Name:      result["name"],
		Num:       &num,
		Picture:   &picture,
		Aid:       aid,
	}
	return prize, err
}

func (p *Prize) StorePrize(ctx context.Context, id int32, prize *model.Prize) error {
	return cli.HMSet(ctx, PrizeTag(id), "create_at", prize.CreatedAt.Format("2006-01-02 15:04:05"), "name", prize.Name, "num", prize.Num, "picture", prize.Picture, "aid", prize.Aid).Err()
}

func (p *Prize) ExistPrizeByAid(ctx context.Context, id int32) bool {
	return cli.Exists(ctx, PrizeByAidTag(id)).Val() == 1
}

func (p *Prize) StorePrizeByAid(ctx context.Context, id int32, prizes []*model.Prize) error {
	for _, v := range prizes {
		err := cli.SAdd(ctx, PrizeByAidTag(id), v.ID).Err()
		if err != nil {
			return err
		}
		err = p.StorePrize(ctx, int32(v.ID), v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Prize) GetPrizeByAid(ctx context.Context, id int32) ([]*model.Prize, error) {
	m, err := cli.SMembers(ctx, PrizeByAidTag(id)).Result()
	if err != nil {
		return nil, err
	}
	ret_prizes := []*model.Prize{}
	for _, v := range m {
		pid, _ := strconv.Atoi(v)
		result, err := cli.HGetAll(ctx, PrizeTag(id)).Result()
		if err != nil {
			return nil, err
		}
		created_at, _ := time.Parse("2006-01-02 15:04:05", result["create_at"])
		num, _ := strconv.ParseInt(result["num"], 10, 64)
		picture := result["picture"]
		aid, _ := strconv.Atoi(result["aid"])
		ret_prize := &model.Prize{
			ID:        pid,
			CreatedAt: &created_at,
			Name:      result["name"],
			Num:       &num,
			Picture:   &picture,
			Aid:       aid,
		}
		ret_prizes = append(ret_prizes, ret_prize)
	}
	return ret_prizes, err
}
