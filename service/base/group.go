package base

import (
	"context"
	"sign-lottery/dao/db/model"
	"sign-lottery/kitex_gen/user"
	"sign-lottery/pkg/errmsg"
	. "sign-lottery/pkg/log"
	"time"
)

// CreateGroup implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) CreateGroup(ctx context.Context, req *user.CreateGroupRequest) (resp *user.BaseResponse, err error) {
	// TODO: Your code here...
	resp = new(user.BaseResponse)
	name := req.GetName()
	start, _ := time.Parse("2006-01-02 15:04:05", req.GetStart())
	end, _ := time.Parse("2006-01-02 15:04:05", req.GetEnd())
	avater := req.GetAvater()
	owner := req.GetOwner()
	group := &model.SignGroup{
		Name:   name,
		Start:  start,
		End:    end,
		Avater: &avater,
		Owner:  owner,
	}
	err = s.dao.Group.CreateGroup(ctx, group)
	if err != nil {
		Log.Errorln("Create group to db err:", err)
		resp.Code = errmsg.Error
		resp.Msg = errmsg.GetMsg(errmsg.Error)
		return nil, err
	}
	err = s.cache.Group.StoreGroupInfo(ctx, group)
	if err != nil {
		Log.Errorln("store group info to cache err:", err)
	}
	resp.Code = errmsg.Success
	resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

// JoinGroup implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) JoinGroup(ctx context.Context, req *user.JoinGroupRequest) (resp *user.BaseResponse, err error) {
	// TODO: Your code here...
	resp = new(user.BaseResponse)
	uid, gid := req.GetUid(), req.GetGid()
	err = s.dao.Group.JoinGroup(ctx, uid, gid)
	if err != nil {
		Log.Errorln("Join group err:", err)
		resp.Code = errmsg.Error
		resp.Msg = errmsg.GetMsg(errmsg.Error)
		return nil, err
	}
	resp.Code = errmsg.Success
	resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

// GetGroupById implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) GetGroupById(ctx context.Context, req *user.GetGroupByIdRequest) (resp *user.GroupResponse, err error) {
	// TODO: Your code here...
	resp = new(user.GroupResponse)
	id := req.GetId()
	if !s.cache.Group.GroupInfoExist(ctx, id) {
		group, err := s.dao.Group.GetGroupById(ctx, id)
		if err != nil {
			Log.Errorln("get group info by id from db err:", err)
			resp.Resp.Code = errmsg.Error
			resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
			return nil, err
		}
		err = s.cache.Group.StoreGroupInfo(ctx, group)
		if err != nil {
			Log.Errorln("store group info to cache err:", err)
		}
	}
	group, err := s.cache.Group.GetGroupInfo(ctx, id)
	if err != nil {
		Log.Errorln("get group info by id from cache err:", err)
		resp.Resp.Code = errmsg.Error
		resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
		return nil, err
	}
	ret_group := &user.GroupInfo{
		Id:         id,
		Name:       group.Name,
		Start:      group.Start.Format("2006-01-02 15:04:05"),
		End:        group.End.Format("2006-01-02 15:04:05"),
		Count:      *group.Count,
		Avater:     *group.Avater,
		Owner:      group.Owner,
		CreateTime: group.CreatedAt.Format("2006-01-02 15:04:05"),
	}
	resp.Group = ret_group
	resp.Resp.Code = errmsg.Success
	resp.Resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

// GetAllGroup implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) GetAllGroup(ctx context.Context, req *user.GetAllGroupRequest) (resp *user.GroupsResponse, err error) {
	// TODO: Your code here...
	resp = new(user.GroupsResponse)
	offset, limit := req.GetOffset(), req.GetLimit()
	groups, count, err := s.dao.Group.GetAllGroup(ctx, int(offset), int(limit))
	if err != nil {
		Log.Errorln("get all groups err:", err)
		resp.Resp.Code = errmsg.Error
		resp.Resp.Msg = errmsg.GetMsg(errmsg.Error)
		return nil, err
	}
	ret_groups := []*user.GroupInfo{}
	for _, v := range groups {
		ret_group := &user.GroupInfo{
			Id:         int32(v.ID),
			Name:       v.Name,
			Start:      v.Start.Format("2006-01-02 15:04:05"),
			End:        v.End.Format("2006-01-02 15:04:05"),
			Count:      *v.Count,
			Avater:     *v.Avater,
			Owner:      v.Owner,
			CreateTime: v.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		ret_groups = append(ret_groups, ret_group)
	}
	resp.Resp.Code = errmsg.Success
	resp.Resp.Msg = errmsg.GetMsg(errmsg.Success)
	resp.Groups = ret_groups
	resp.Total = count
	return
}

// GroupUpdate implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) GroupUpdate(ctx context.Context, req *user.GroupUpdateRequest) (resp *user.BaseResponse, err error) {
	// TODO: Your code here...
	resp = new(user.BaseResponse)
	id := req.GetId()
	name := req.GetName()
	start, _ := time.Parse("2006-01-02 15:04:05", req.GetStart())
	end, _ := time.Parse("2006-01-02 15:04:05", req.GetEnd())
	avater := req.GetAvater()
	owner := req.GetOwner()
	uid := req.GetUid()
	previlege := s.dao.Group.CheckGroupPrevilege(ctx, uid, id)
	if !previlege {
		resp.Code = errmsg.NoPreviledge
		resp.Msg = errmsg.GetMsg(errmsg.NoPreviledge)
		return nil, err
	}
	group := &model.SignGroup{
		ID:     int(id),
		Name:   name,
		Start:  start,
		End:    end,
		Avater: &avater,
		Owner:  owner,
	}
	err = s.dao.Group.GroupUpdate(ctx, id, group)
	if err != nil {
		Log.Errorln("update group to db error:", err)
		resp.Code = errmsg.Error
		resp.Msg = errmsg.GetMsg(errmsg.Error)
		return nil, err
	}
	err = s.cache.Group.ClearGroupInfo(ctx, id)
	if err != nil {
		Log.Errorln("clear group info from cache err:", err)
	}
	go func() {
		time.Sleep(time.Millisecond)
		s.cache.Group.ClearGroupInfo(ctx, id)
	}()
	resp.Code = errmsg.Success
	resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}

// GroupDel implements the BaseServiceImpl interface.
func (s *BaseServiceImpl) GroupDel(ctx context.Context, req *user.GroupDelRequest) (resp *user.BaseResponse, err error) {
	resp = new(user.BaseResponse)
	id := req.GetId()
	uid := req.GetUid()
	err = s.dao.Group.GroupDel(ctx, id)
	if err != nil {
		Log.Errorln("delete from db err:", err)
		resp.Code = errmsg.Error
		resp.Msg = errmsg.GetMsg(errmsg.Error)
		return nil, err
	}
	previlege := s.dao.Group.CheckGroupPrevilege(ctx, uid, id)
	if !previlege {
		resp.Code = errmsg.NoPreviledge
		resp.Msg = errmsg.GetMsg(errmsg.NoPreviledge)
		return nil, err
	}
	if s.cache.Group.GroupInfoExist(ctx, id) {
		err = s.cache.Group.ClearGroupInfo(ctx, id)
		if err != nil {
			Log.Errorln("clear group from cache err:", err)
		}
		go func() {
			time.Sleep(time.Millisecond)
			s.cache.Group.ClearGroupInfo(ctx, id)
		}()
	}
	resp.Code = errmsg.Success
	resp.Msg = errmsg.GetMsg(errmsg.Success)
	return
}
