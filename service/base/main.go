package base

import (
	"context"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"net"
	user "sign-lottery/kitex_gen/user/baseservice"
	"sign-lottery/pkg/constants"
	. "sign-lottery/pkg/log"
)

func BaseServer() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdDSN})
	if err != nil {
		Log.Fatalln("get etcd registry err:", err)
	}
	addr, err := net.ResolveTCPAddr("tcp", constants.BaseTCPAddress)
	if err != nil {
		Log.Fatalln("registry base tcp addr err:", err)
	}
	baseServer := new(BaseServiceImpl)
	svr := user.NewServer(baseServer,
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithMuxTransport(),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.BaseServiceName}),
	)
	go func() {
		for {
			baseServer.SendEmail(context.Background())
		}
	}()
	err = svr.Run()

	if err != nil {
		Log.Fatalln("base service run err:", err)
	}
}
