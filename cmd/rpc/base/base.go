package base

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
	"sign-lottery/kitex_gen/user/baseservice"
	"sign-lottery/pkg/constants"
	. "sign-lottery/pkg/log"
)

var BaseClient baseservice.Client

func newBaseClient() {
	var err error
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdDSN})
	if err != nil {
		Log.Fatalln("get etcd registry err:", err)
	}
	BaseClient, err = baseservice.NewClient(
		constants.BaseServiceName,
		client.WithResolver(r),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.BaseServiceName}),
	)
	if err != nil {
		Log.Fatalln("get base client err:", err)
	}
}

func init() {
	if BaseClient == nil {
		newBaseClient()
	}
}
