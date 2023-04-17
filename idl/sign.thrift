namespace go sign

struct BaseResponse{
    1:required string msg
    2:required i32 code
}

struct SignRequest{
    1:optional i64 uid
    2:required i32 gid
    3:required double latitude
    4:required double longtitude
    5:optional string ip
}

struct AskLeaveRequest{
    1:optional i64 uid
    2:required i32 gid
    3:optional string time
    4:required string issue
}

struct GetMonthSignRequest{
    1:optional i64 uid
    2:required i32 gid
    3:required string month
}

struct MonthSignResponse{
    1:required BaseResponse resp
    2:required i32 bitmap
}

struct GetMonthSignsByGid{
    1:required i32 gid
    2:required string month
    3:required i32 offset
    4:required i32 limit
}

struct MonthSignsResponse{
    1:required BaseResponse resp
    2:required list<i32>bitmap
    3:required i64 count
}

struct GetAllRecordRequest{
    1:required i32 offset
    2:required i32 limit
}

struct Record{
    1:required i64 uid
    2:required i32 gid
    3:required string time
}

struct RecordsResponse{
    1:required BaseResponse resp
    2:required i64 total
    3:required list<Record> records
}

struct GetUserRecordRequest{
    1:optional i64 uid
    2:required i32 offset
    3:required i32 limit
}

struct SignPosAddRequest{
    1:required i32 gid
    2:required string name
    3:required double longtitude
    4:required double latitude
    5:required i64 uid
}

struct SignPosDelRequest{
    1:required i32 gid
    2:required string name
    3:required i64 uid
}

struct GetSignPosRequest{
    1:required i32 gid
    2:required i32 offset
    3:required i32 limit
}

struct PosInfo{
    1:required string name
    2:required double longtitude
    3:required double latitude
}

struct GetSignPosResponse{
    1:required BaseResponse resp
    2:required i64 total
    3:required list<PosInfo> pos
}

service SignService{
     BaseResponse Sign()
     BaseResponse AskLeave(1:AskLeaveRequest req)
     MonthSignResponse GetMonthSign(1:GetMonthSignRequest req)
     MonthSignsResponse GetMonthSignByGid(1:GetMonthSignsByGid req)
     RecordsResponse GetAllRecord(1:GetAllRecordRequest req)
     RecordsResponse GetUserRecord(1:GetUserRecordRequest req)
     BaseResponse SignPosAdd(1:SignPosAddRequest req)
     BaseResponse SignPosDel(1:SignPosDelRequest req)
     GetSignPosResponse GetSignPos(1:GetSignPosRequest req)
}