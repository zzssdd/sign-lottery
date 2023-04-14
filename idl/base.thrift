namespace go user

struct UserInfo{
    1:required i64 id
    2:required string creat_time
    3:required string email
    4:required string name
    5:required string avater
}

struct GroupInfo{
    1:required i32 id
    2:required string name
    3:required string start
    4:required string end
    5:required i64 count
    6:required string avater
    7:required i64 owner
    8:required string create_time
}

struct EmailRequest{
    1:required string email
}

struct RegisterRequest{
    1:required string email
    2:required string password
    3:required string name
    4:required string code
}

struct BaseResponse{
    1:required string msg
    2:required i32 code
}

struct LoginRequest{
    1:required string email
    2:required string password
}

struct LoginResponse{
    1:required BaseResponse resp
    2:required i64 id
    3:required string token
}

struct GetUserByIdRequest{
    1:required i64 id
}


struct UserResponse{
    1:required BaseResponse resp
    2:required UserInfo user
}

struct UsersResponse{
    1:required BaseResponse resp
    2:required i64 total
    3:required list<UserInfo> users
}


struct GetAllUserRequest{
    1:required i32 offset
    2:required i32 limit
}

struct AdminLoginRequest{
    1:required string name
    2:required string password
}

struct AdminLoginResponse{
    1:required BaseResponse resp
    2:required string token
    3:required i64 id
}


struct CreateGroupRequest{
    1:required string name
    2:required string start
    3:required string end
    4:optional string avater
    5:required i64 owner
}

struct JoinGroupRequest{
    1:optional i64 uid
    2:required i32 gid
}

struct GetUserByGidRequest{
    1:required i32 gid
    2:required i32 offset
    3:required i32 limit
}

struct GetAllGroupRequest{
    1:required i32 offset
    2:required i32 limit
}

struct GroupsResponse{
    1:required BaseResponse resp
    2:required i64 total
    3:required list<GroupInfo> groups
}

struct GetGroupByIdRequest{
    1:required i32 id
}

struct GroupResponse{
    1:required GroupInfo group
    2:required BaseResponse resp
}

struct ChangeUserAvaterRequest{
    1:optional i64 id
    2:required string avater
}

struct ChangePasswordRequest{
    1:optional i64 id
    2:required string old
    3:required string new
}

struct ChangeAddressRequest{
    1:optional i64 id
    2:required string address
}

struct GroupUpdateRequest{
    1:required i32 id
    2:required string name
    3:required string start
    4:required string end
    5:required string avater
    6:required i64 owner
}

struct UserDelRequest{
    1:required i64 id
}

struct GroupDelRequest{
    1:required i32 id
}

service BaseService{
     BaseResponse SendEmail(1:EmailRequest req)
     BaseResponse Registe(1:RegisterRequest req)
     LoginResponse Login(1:LoginRequest req)
     AdminLoginResponse AdminLogin(1:AdminLoginRequest req)
     UserResponse GetUserById(1:GetUserByIdRequest req)
     UsersResponse GetAllUser(1:GetAllUserRequest req)
     UsersResponse GetUserByGid(1:GetUserByGidRequest req)
     BaseResponse ChangeUserAvater(1:ChangeUserAvaterRequest req)
     BaseResponse ChangeUserPassword(1:ChangePasswordRequest req)
     BaseResponse ChangeUserAddress(1:ChangeAddressRequest req)
     BaseResponse UserDel(1:UserDelRequest req)
     BaseResponse CreateGroup(1:CreateGroupRequest req)
     BaseResponse JoinGroup(1:JoinGroupRequest req)
     GroupResponse GetGroupById(1:GetGroupByIdRequest req)
     GroupsResponse GetAllGroup(1:GetAllGroupRequest req)
     BaseResponse GroupUpdate(1:GroupUpdateRequest req)
     BaseResponse GroupDel(1:GroupDelRequest req)
     //GetGroupByUid
}