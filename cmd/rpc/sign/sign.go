package sign

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
	"sign-lottery/kitex_gen/sign/signservice"
	"sign-lottery/pkg/constants"
	. "sign-lottery/pkg/log"
)

var SignClient signservice.Client

func newSignClient() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdDSN})
	if err != nil {
		Log.Fatalln("get etcd registry err:", err)
	}
	SignClient, err = signservice.NewClient(
		constants.SignServiceName,
		client.WithResolver(r),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.SignServiceName}),
	)
	if err != nil {
		Log.Fatalln("get sign client err:", err)
	}
}

func init() {
	if SignClient == nil {
		newSignClient()
	}
}
