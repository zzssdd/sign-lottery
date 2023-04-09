package base

import (
	"context"
	"sign-lottery/kitex_gen/user"
)

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
