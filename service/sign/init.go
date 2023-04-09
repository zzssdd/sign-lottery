package sign

import (
	"context"
	"sign-lottery/dao/cache"
	"sign-lottery/dao/db"
	sign "sign-lottery/kitex_gen/sign"
)

// SignServiceImpl implements the last service interface defined in the IDL.
type SignServiceImpl struct {
	dao   *db.Dao
	cache *cache.Cache
}

// Sign implements the SignServiceImpl interface.
func (s *SignServiceImpl) Sign(ctx context.Context, req *sign.SignRequest) (resp *sign.BaseResponse, err error) {
	// TODO: Your code here...
	return
}

// AskLeave implements the SignServiceImpl interface.
func (s *SignServiceImpl) AskLeave(ctx context.Context, req *sign.AskLeaveRequest) (resp *sign.BaseResponse, err error) {
	// TODO: Your code here...
	return
}

// GetMonthSign implements the SignServiceImpl interface.
func (s *SignServiceImpl) GetMonthSign(ctx context.Context, req *sign.GetMonthSignRequest) (resp *sign.MonthSignResponse, err error) {
	// TODO: Your code here...
	return
}

// GetMonthSignByGid implements the SignServiceImpl interface.
func (s *SignServiceImpl) GetMonthSignByGid(ctx context.Context, req *sign.GetMonthSignsByGid) (resp *sign.MonthSignsResponse, err error) {
	// TODO: Your code here...
	return
}

// GetAllRecord implements the SignServiceImpl interface.
func (s *SignServiceImpl) GetAllRecord(ctx context.Context, req *sign.GetAllRecordRequest) (resp *sign.RecordsResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUserRecord implements the SignServiceImpl interface.
func (s *SignServiceImpl) GetUserRecord(ctx context.Context, req *sign.GetUserRecordRequest) (resp *sign.RecordsResponse, err error) {
	// TODO: Your code here...
	return
}

// SignPosAdd implements the SignServiceImpl interface.
func (s *SignServiceImpl) SignPosAdd(ctx context.Context, req *sign.SignPosAddRequest) (resp *sign.BaseResponse, err error) {
	// TODO: Your code here...
	return
}

// SignPosDel implements the SignServiceImpl interface.
func (s *SignServiceImpl) SignPosDel(ctx context.Context, req *sign.SignPosDelRequest) (resp *sign.BaseResponse, err error) {
	// TODO: Your code here...
	return
}

// GetSignPos implements the SignServiceImpl interface.
func (s *SignServiceImpl) GetSignPos(ctx context.Context, req *sign.GetSignPosRequest) (resp *sign.GetSignPosResponse, err error) {
	// TODO: Your code here...
	return
}
