package lottery

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
	"sign-lottery/kitex_gen/lottery/lotteryservice"
	"sign-lottery/pkg/constants"
	. "sign-lottery/pkg/log"
)

var LotteryClient lotteryservice.Client

func newLotteryClient() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdDSN})
	if err != nil {
		Log.Fatalln("get etcd registry err:", err)
	}
	LotteryClient, err = lotteryservice.NewClient(
		constants.LotteryServiceName,
		client.WithResolver(r),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.LotteryServiceName}),
	)
	if err != nil {
		Log.Fatalln("get lottery client err:", err)
	}
}

func init() {
	if LotteryClient == nil {
		newLotteryClient()
	}
}
