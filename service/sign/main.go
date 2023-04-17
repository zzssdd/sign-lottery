package sign

import (
	"context"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"net"
	. "sign-lottery/pkg/log"

	sign "sign-lottery/kitex_gen/sign/signservice"
	"sign-lottery/pkg/constants"
)

func SignServer() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdDSN})
	if err != nil {
		Log.Fatalln("get etcd registry err:", err)
	}
	addr, err := net.ResolveTCPAddr("tcp", constants.SignTCPAddr)
	if err != nil {
		Log.Fatalln("registry sign tcp addr err:", err)
	}
	signServer := new(SignServiceImpl)
	svr := sign.NewServer(signServer,
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithMuxTransport(),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.SignServiceName}),
	)
	err = svr.Run()
	go func() {
		for {
			signServer.Sign(context.Background())
			signServer.HandleSuccessSign(context.Background())
		}
	}()
	if err != nil {
		Log.Fatalln("sign server run err:", err)
	}
}
