package sign

import (
	"sign-lottery/dao/cache"
	"sign-lottery/dao/db"
)

// SignServiceImpl implements the last service interface defined in the IDL.
type SignServiceImpl struct {
	dao   *db.Dao
	cache *cache.Cache
}

func NewService() *SignServiceImpl {
	return &SignServiceImpl{
		dao:   db.NewDao(),
		cache: cache.NewCache(),
	}
}
