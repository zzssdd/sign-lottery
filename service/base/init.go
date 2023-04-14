package base

import (
	"sign-lottery/dao/cache"
	"sign-lottery/dao/db"
)

// BaseServiceImpl implements the last service interface defined in the IDL.
type BaseServiceImpl struct {
	dao   *db.Dao
	cache *cache.Cache
}

func NewService() *BaseServiceImpl {
	return &BaseServiceImpl{
		dao:   db.NewDao(),
		cache: cache.NewCache(),
	}
}
