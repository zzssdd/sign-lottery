// Code generated by Kitex v0.5.1. DO NOT EDIT.
package baseservice

import (
	server "github.com/cloudwego/kitex/server"
	user "sign-lottery/kitex_gen/user"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler user.BaseService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
