package lottery

import (
	"sign-lottery/dao/cache"
	"sign-lottery/dao/db"
)

// LotteryServiceImpl implements the last service interface defined in the IDL.
type LotteryServiceImpl struct {
	dao   *db.Dao
	cache *cache.Cache
}

func NewService() *LotteryServiceImpl {
	return &LotteryServiceImpl{
		dao:   db.NewDao(),
		cache: cache.NewCache(),
	}
}
