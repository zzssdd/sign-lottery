package lottery

import (
	"context"
	"sign-lottery/dao/cache"
	"sign-lottery/dao/db"
	lottery "sign-lottery/kitex_gen/lottery"
)

// LotteryServiceImpl implements the last service interface defined in the IDL.
type LotteryServiceImpl struct {
	dao   *db.Dao
	cache *cache.Cache
}

// ActivityAdd implements the LotteryServiceImpl interface.
func (s *LotteryServiceImpl) ActivityAdd(ctx context.Context, req *lottery.ActivityAddRequest) (resp *lottery.BaseResponse, err error) {
	// TODO: Your code here...
	return
}

// ActivityDel implements the LotteryServiceImpl interface.
func (s *LotteryServiceImpl) ActivityDel(ctx context.Context, req *lottery.ActivityDelRequest) (resp *lottery.BaseResponse, err error) {
	// TODO: Your code here...
	return
}

// ActivityUpdate implements the LotteryServiceImpl interface.
func (s *LotteryServiceImpl) ActivityUpdate(ctx context.Context, req *lottery.ActivityUpdateRequest) (resp *lottery.BaseResponse, err error) {
	// TODO: Your code here...
	return
}

// GetActivityByGid implements the LotteryServiceImpl interface.
func (s *LotteryServiceImpl) GetActivityByGid(ctx context.Context, req *lottery.GetActivityByGidRequest) (resp *lottery.ActivitysResponse, err error) {
	// TODO: Your code here...
	return
}

// GetAllActivity implements the LotteryServiceImpl interface.
func (s *LotteryServiceImpl) GetAllActivity(ctx context.Context, req *lottery.GetAllActivityRequest) (resp *lottery.ActivitysResponse, err error) {
	// TODO: Your code here...
	return
}

// GetActivityById implements the LotteryServiceImpl interface.
func (s *LotteryServiceImpl) GetActivityById(ctx context.Context, req *lottery.GetActivityByIdRequest) (resp *lottery.ActivityResponse, err error) {
	// TODO: Your code here...
	return
}

// PrizeAdd implements the LotteryServiceImpl interface.
func (s *LotteryServiceImpl) PrizeAdd(ctx context.Context, req *lottery.PrizeAddRequest) (resp *lottery.BaseResponse, err error) {
	// TODO: Your code here...
	return
}

// PrizeDel implements the LotteryServiceImpl interface.
func (s *LotteryServiceImpl) PrizeDel(ctx context.Context, req *lottery.PrizeDelRequest) (resp *lottery.BaseResponse, err error) {
	// TODO: Your code here...
	return
}

// PrizeUpdate implements the LotteryServiceImpl interface.
func (s *LotteryServiceImpl) PrizeUpdate(ctx context.Context, req *lottery.PrizeUpdateRequest) (resp *lottery.BaseResponse, err error) {
	// TODO: Your code here...
	return
}

// GetPrizeByAid implements the LotteryServiceImpl interface.
func (s *LotteryServiceImpl) GetPrizeByAid(ctx context.Context, req *lottery.GetPrizeByAidRequest) (resp *lottery.PrizesResponse, err error) {
	// TODO: Your code here...
	return
}

// GetPrizeById implements the LotteryServiceImpl interface.
func (s *LotteryServiceImpl) GetPrizeById(ctx context.Context, req *lottery.GetPrizeByIdRequest) (resp *lottery.PrizeResponse, err error) {
	// TODO: Your code here...
	return
}

// Choose implements the LotteryServiceImpl interface.
func (s *LotteryServiceImpl) Choose(ctx context.Context, req *lottery.ChooseRequest) (resp *lottery.ChooseResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUserOrder implements the LotteryServiceImpl interface.
func (s *LotteryServiceImpl) GetUserOrder(ctx context.Context, req *lottery.GetUserOrderRequest) (resp *lottery.OrdersResponse, err error) {
	// TODO: Your code here...
	return
}

// GetAllOrder implements the LotteryServiceImpl interface.
func (s *LotteryServiceImpl) GetAllOrder(ctx context.Context, req *lottery.GetAllOrderRequest) (resp *lottery.OrdersResponse, err error) {
	// TODO: Your code here...
	return
}
