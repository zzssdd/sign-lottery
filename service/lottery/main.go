package lottery

import (
	"context"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"net"
	lottery "sign-lottery/kitex_gen/lottery/lotteryservice"
	"sign-lottery/pkg/constants"
	. "sign-lottery/pkg/log"
)

func LotteryServer() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdDSN})
	if err != nil {
		Log.Fatalln("get etcd registry err:", err)
	}
	addr, err := net.ResolveTCPAddr("tcp", constants.LotteryTCPAddr)
	if err != nil {
		Log.Fatalln("registry lottery tcp addr err:", err)
	}
	lotteryServer := new(LotteryServiceImpl)
	svr := lottery.NewServer(lotteryServer,
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithMuxTransport(),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.LotteryServiceName}),
	)
	go func() {
		for {
			lotteryServer.Choose(context.Background())
			err := lotteryServer.HandleRabbitOrder(context.Background())
			if err != nil {
				Log.Errorln("save order from rabbitmq into db err", err)
			}
		}
	}()
	err = svr.Run()

	if err != nil {
		Log.Fatalln("lottery service run err:", err)
	}
}
