package cache

import (
	"context"
	"strconv"
	"time"
)

type IpLimit struct {
}

const IpLimitPreffix = "ip:count:"

func IpLimitTag(ip string) string {
	return IpLimitPreffix + ip
}

func (i *IpLimit) IncrIpCount(ctx context.Context, ip string) error {
	err := cli.Incr(ctx, IpLimitTag(ip)).Err()
	if err != nil {
		return err
	}
	return cli.Expire(ctx, IpLimitTag(ip), 2*time.Minute).Err()
}

func (i *IpLimit) GetIpCount(ctx context.Context, ip string) (int, error) {
	result, err := cli.Get(ctx, IpLimitTag(ip)).Result()
	if err != nil {
		return 0, err
	}
	count, err := strconv.Atoi(result)
	if err != nil {
		return 0, err
	}
	return count, nil
}
