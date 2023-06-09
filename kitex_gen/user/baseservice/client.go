// Code generated by Kitex v0.5.1. DO NOT EDIT.

package baseservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	user "sign-lottery/kitex_gen/user"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	SendEmail(ctx context.Context, callOptions ...callopt.Option) (r *user.BaseResponse, err error)
	Registe(ctx context.Context, req *user.RegisterRequest, callOptions ...callopt.Option) (r *user.BaseResponse, err error)
	Login(ctx context.Context, req *user.LoginRequest, callOptions ...callopt.Option) (r *user.LoginResponse, err error)
	AdminLogin(ctx context.Context, req *user.AdminLoginRequest, callOptions ...callopt.Option) (r *user.AdminLoginResponse, err error)
	GetUserById(ctx context.Context, req *user.GetUserByIdRequest, callOptions ...callopt.Option) (r *user.UserResponse, err error)
	GetAllUser(ctx context.Context, req *user.GetAllUserRequest, callOptions ...callopt.Option) (r *user.UsersResponse, err error)
	GetUserByGid(ctx context.Context, req *user.GetUserByGidRequest, callOptions ...callopt.Option) (r *user.UsersResponse, err error)
	ChangeUserAvater(ctx context.Context, req *user.ChangeUserAvaterRequest, callOptions ...callopt.Option) (r *user.BaseResponse, err error)
	ChangeUserPassword(ctx context.Context, req *user.ChangePasswordRequest, callOptions ...callopt.Option) (r *user.BaseResponse, err error)
	ChangeUserAddress(ctx context.Context, req *user.ChangeAddressRequest, callOptions ...callopt.Option) (r *user.BaseResponse, err error)
	UserDel(ctx context.Context, req *user.UserDelRequest, callOptions ...callopt.Option) (r *user.BaseResponse, err error)
	CreateGroup(ctx context.Context, req *user.CreateGroupRequest, callOptions ...callopt.Option) (r *user.BaseResponse, err error)
	JoinGroup(ctx context.Context, req *user.JoinGroupRequest, callOptions ...callopt.Option) (r *user.BaseResponse, err error)
	GetGroupById(ctx context.Context, req *user.GetGroupByIdRequest, callOptions ...callopt.Option) (r *user.GroupResponse, err error)
	GetAllGroup(ctx context.Context, req *user.GetAllGroupRequest, callOptions ...callopt.Option) (r *user.GroupsResponse, err error)
	GroupUpdate(ctx context.Context, req *user.GroupUpdateRequest, callOptions ...callopt.Option) (r *user.BaseResponse, err error)
	GroupDel(ctx context.Context, req *user.GroupDelRequest, callOptions ...callopt.Option) (r *user.BaseResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kBaseServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kBaseServiceClient struct {
	*kClient
}

func (p *kBaseServiceClient) SendEmail(ctx context.Context, callOptions ...callopt.Option) (r *user.BaseResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SendEmail(ctx)
}

func (p *kBaseServiceClient) Registe(ctx context.Context, req *user.RegisterRequest, callOptions ...callopt.Option) (r *user.BaseResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Registe(ctx, req)
}

func (p *kBaseServiceClient) Login(ctx context.Context, req *user.LoginRequest, callOptions ...callopt.Option) (r *user.LoginResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Login(ctx, req)
}

func (p *kBaseServiceClient) AdminLogin(ctx context.Context, req *user.AdminLoginRequest, callOptions ...callopt.Option) (r *user.AdminLoginResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AdminLogin(ctx, req)
}

func (p *kBaseServiceClient) GetUserById(ctx context.Context, req *user.GetUserByIdRequest, callOptions ...callopt.Option) (r *user.UserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserById(ctx, req)
}

func (p *kBaseServiceClient) GetAllUser(ctx context.Context, req *user.GetAllUserRequest, callOptions ...callopt.Option) (r *user.UsersResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetAllUser(ctx, req)
}

func (p *kBaseServiceClient) GetUserByGid(ctx context.Context, req *user.GetUserByGidRequest, callOptions ...callopt.Option) (r *user.UsersResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserByGid(ctx, req)
}

func (p *kBaseServiceClient) ChangeUserAvater(ctx context.Context, req *user.ChangeUserAvaterRequest, callOptions ...callopt.Option) (r *user.BaseResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChangeUserAvater(ctx, req)
}

func (p *kBaseServiceClient) ChangeUserPassword(ctx context.Context, req *user.ChangePasswordRequest, callOptions ...callopt.Option) (r *user.BaseResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChangeUserPassword(ctx, req)
}

func (p *kBaseServiceClient) ChangeUserAddress(ctx context.Context, req *user.ChangeAddressRequest, callOptions ...callopt.Option) (r *user.BaseResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChangeUserAddress(ctx, req)
}

func (p *kBaseServiceClient) UserDel(ctx context.Context, req *user.UserDelRequest, callOptions ...callopt.Option) (r *user.BaseResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UserDel(ctx, req)
}

func (p *kBaseServiceClient) CreateGroup(ctx context.Context, req *user.CreateGroupRequest, callOptions ...callopt.Option) (r *user.BaseResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateGroup(ctx, req)
}

func (p *kBaseServiceClient) JoinGroup(ctx context.Context, req *user.JoinGroupRequest, callOptions ...callopt.Option) (r *user.BaseResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.JoinGroup(ctx, req)
}

func (p *kBaseServiceClient) GetGroupById(ctx context.Context, req *user.GetGroupByIdRequest, callOptions ...callopt.Option) (r *user.GroupResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetGroupById(ctx, req)
}

func (p *kBaseServiceClient) GetAllGroup(ctx context.Context, req *user.GetAllGroupRequest, callOptions ...callopt.Option) (r *user.GroupsResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetAllGroup(ctx, req)
}

func (p *kBaseServiceClient) GroupUpdate(ctx context.Context, req *user.GroupUpdateRequest, callOptions ...callopt.Option) (r *user.BaseResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GroupUpdate(ctx, req)
}

func (p *kBaseServiceClient) GroupDel(ctx context.Context, req *user.GroupDelRequest, callOptions ...callopt.Option) (r *user.BaseResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GroupDel(ctx, req)
}
