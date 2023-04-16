package common

import (
	"github.com/jinzhu/copier"
)

func BindRpcOption(req interface{}, rpcReq interface{}) error {
	if rpcReq != nil {
		err := copier.Copy(rpcReq, req)
		if err != nil {
			return err
		}
	}
	return nil
}
