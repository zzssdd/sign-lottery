package cache

import (
	"sign-lottery/pkg/constants"

	"github.com/redis/go-redis/v9"
)

var cli *redis.Client

type Cache struct {
	User User
}

func NewCache() *Cache {
	if cli == nil {
		cli = redis.NewClient(&redis.Options{
			Addr:     constants.RedisDSN,
			Password: "",
			DB:       0,
		})
	}
	return &Cache{
		User: User{},
	}
}