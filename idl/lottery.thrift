namespace go lottery

struct Activity{
    1:required i32 id
    2:required string creat_time
    3:required string name
    4:required string picture
    5:required string desc
    6:required i64 cost
    7:required i64 uid
    8:required i32 gid
    9:required string start
    10:required string end
}

struct Prize{
    1:required i32 id
    2:required string create_time
    3:required string name
    4:required i64 num
    5:required string picture
}

struct BaseResponse{
    1:required string msg
    2:required i32 code
}

struct ActivityAddRequest{
    1:required string name
    2:optional string picture
    3:required string dec
    4:required i32 cost
    5:optional i64 uid
    6:required i32 gid
    7:required string start
    8:required string end
}

struct ActivityUpdateRequest{
    1:required i32 id
    2:required string name
    3:required string picture
    4:required string desc
    5:required i32 cost
    6:required i32 uid
    7:required i32 gid
    8:required string start
    9:required string end
}

struct ActivityDelRequest{
    1:required i32 id
}

struct GetActivityByGidRequest{
    1:required i32 gid
    2:required i64 offset
    3:required i64 limit
}

struct GetAllActivityRequest{
    1:required i64 offset
    2:required i64 limit
}

struct ActivitysResponse{
    1:required BaseResponse resp
    2:required i64 total
    3:required list<Activity> activity
}

struct GetActivityByIdRequest{
    1:required i32 id
}

struct ActivityResponse{
    1:required BaseResponse resp
    2:required Activity activity
}

struct PrizeAddRequest{
    1:required string name
    2:required i64 num
    3:required string picture
    4:required i32 aid
}

struct PrizeDelRequest{
    1:required i32 id
}

struct PrizeUpdateRequest{
    1:required i32 id
    2:required string name
    3:required i64 num
    4:required string picture
    5:required i32 aid
}

struct GetPrizeByAidRequest{
    1:required i32 aid
}

struct GetPrizeByIdRequest{
    1:required i32 id
}

struct PrizeResponse{
    1:required BaseResponse resp
    2:required Prize prize
}

struct PrizesResponse{
    1:required BaseResponse resp
    2:required i64 total
    3:required list<Prize> prizes
}

struct ChooseRequest{
    1:optional i64 uid
    2:required i32 aid
}

struct ChooseResponse{
    1:required BaseResponse resp
    2:required string name
}

struct GetUserOrderRequest{
    1:optional i64 uid
    2:required i64 offset
    3:required i64 limit
}

struct GetAllOrderRequest{
    1:required i64 offset
    2:required i64 limit
}

struct Order{
    1:required i64 id
    2:required string create_time
    3:required i64 uid
    4:required i32 pid
}

struct OrdersResponse{
    1:required BaseResponse resp
    2:required i64 total
    3:required list<Order> order
}

service LotteryService{
    BaseResponse ActivityAdd(1:ActivityAddRequest req)
    BaseResponse ActivityDel(1:ActivityDelRequest req)
    BaseResponse ActivityUpdate(1:ActivityUpdateRequest req)
    ActivitysResponse GetActivityByGid(1:GetActivityByGidRequest req)
    ActivitysResponse GetAllActivity(1:GetAllActivityRequest req)
    ActivityResponse GetActivityById(1:GetActivityByIdRequest req)
    BaseResponse PrizeAdd(1:PrizeAddRequest req)
    BaseResponse PrizeDel(1:PrizeDelRequest req)
    BaseResponse PrizeUpdate(1:PrizeUpdateRequest req)
    PrizesResponse GetPrizeByAid(1:GetPrizeByAidRequest req)
    PrizeResponse GetPrizeById(1:GetPrizeByIdRequest req)
    ChooseResponse Choose(1:ChooseRequest req)
    OrdersResponse GetUserOrder(1:GetUserOrderRequest req)
    OrdersResponse GetAllOrder(1:GetAllOrderRequest req)
}