package base

import (
	"context"
	user "sign-lottery/kitex_gen/user"
)

// BaseServiceImpl implements the last service interface defined in the IDL.
type BaseServiceImpl struct{}

// SendEmail implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) SendEmail(ctx context.Context, req *user.EmailRequest) (resp *user.BaseResponse, err error) {
	// TODO: Your code here...
	return
}

// Registe implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) Registe(ctx context.Context, req *user.RegisterRequest) (resp *user.BaseResponse, err error) {
	// TODO: Your code here...
	return
}

// Login implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	// TODO: Your code here...
	return
}

// AdminLogin implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) AdminLogin(ctx context.Context, req *user.AdminLoginRequest) (resp *user.AdminLoginResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUserById implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) GetUserById(ctx context.Context, req *user.GetUserByIdRequest) (resp *user.UserResponse, err error) {
	// TODO: Your code here...
	return
}

// GetAllUser implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) GetAllUser(ctx context.Context, req *user.GetAllUserRequest) (resp *user.UsersResponse, err error) {
	// TODO: Your code here...
	return
}

// GetUserByGid implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) GetUserByGid(ctx context.Context, req *user.GetUserByGidRequest) (resp *user.UsersResponse, err error) {
	// TODO: Your code here...
	return
}

// ChangeUserAvater implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) ChangeUserAvater(ctx context.Context, req *user.ChangeUserAvaterRequest) (resp *user.BaseResponse, err error) {
	// TODO: Your code here...
	return
}

// ChangeUserPassword implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) ChangeUserPassword(ctx context.Context, req *user.ChangePasswordRequest) (resp *user.BaseResponse, err error) {
	// TODO: Your code here...
	return
}

// ChangeUserAddress implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) ChangeUserAddress(ctx context.Context, req *user.ChangeAddressRequest) (resp *user.BaseResponse, err error) {
	// TODO: Your code here...
	return
}

// UserDel implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) UserDel(ctx context.Context, req *user.UserDelRequest) (resp *user.BaseResponse, err error) {
	// TODO: Your code here...
	return
}

// CreateGroup implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) CreateGroup(ctx context.Context, req *user.CreateGroupRequest) (resp *user.BaseResponse, err error) {
	// TODO: Your code here...
	return
}

// JoinGroup implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) JoinGroup(ctx context.Context, req *user.JoinGroupRequest) (resp *user.BaseResponse, err error) {
	// TODO: Your code here...
	return
}

// GetGroupById implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) GetGroupById(ctx context.Context, req *user.GetGroupByIdRequest) (resp *user.GroupResponse, err error) {
	// TODO: Your code here...
	return
}

// GetAllGroup implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) GetAllGroup(ctx context.Context, req *user.GetAllGroupRequest) (resp *user.GroupsResponse, err error) {
	// TODO: Your code here...
	return
}

// GroupUpdate implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) GroupUpdate(ctx context.Context, req *user.GroupUpdateRequest) (resp *user.BaseResponse, err error) {
	// TODO: Your code here...
	return
}

// GroupDel implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) GroupDel(ctx context.Context, req *user.GroupDelRequest) (resp *user.BaseResponse, err error) {
	// TODO: Your code here...
	return
}
