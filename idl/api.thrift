include "base.thrift"
include "sign.thrift"
include "lottery.thrift"

namespace go api

service SignLotteryApi{
    base.BaseResponse SendEmail(1:base.EmailRequest req)(api.post="/user/email/")
    base.BaseResponse Registe(1:base.RegisterRequest req)(api.post="/user/registe/")
    base.LoginResponse Login(1:base.LoginRequest req)(api.post="/user/login/")
    base.UsersResponse GetUserById(1:base.GetUserByIdRequest req)(api.get="/user/id/")
    base.UsersResponse GetUserByGid(1:base.GetUserByGidRequest req)(api.get="/user/gid/")
    base.BaseResponse ChangeUserAvater(1:base.ChangeUserAvaterRequest req)(api.put="/user/avater/")
    base.BaseResponse ChangeUserPassword(1:base.ChangePasswordRequest req)(api.put="/user/password/")
    base.BaseResponse ChangeUserAddress(1:base.ChangeAddressRequest req)(api.put="/user/address/")

    base.BaseResponse CreateGroup(1:base.CreateGroupRequest req)(api.post="/group/add/")
    base.BaseResponse JoinGroup(1:base.JoinGroupRequest req)(api.post="/group/join/")
    base.GroupResponse GetGroupById(1:base.GetGroupByIdRequest req)(api.get="/group/id/")
    base.GroupsResponse GetAllGroup(1:base.GetAllGroupRequest req)(api.get="/group/list/")
    base.BaseResponse GroupUpdate(1:base.GroupUpdateRequest req)(api.put="/group/put/")
    base.BaseResponse GroupDel(1:base.GroupDelRequest req)(api.delete="/group/del/")

    base.AdminLoginResponse AdminLogin(1:base.AdminLoginRequest req)(api.post="/admin/login/")
    base.BaseResponse UserDel(1:base.UserDelRequest req)(api.delete="/admin/user/")
    base.UsersResponse GetAllUser(1:base.GetAllUserRequest req)(api.get="/admin/userlist/")

    sign.BaseResponse Sign(1:sign.SignRequest req)(api.post="/sign/add/")
    sign.BaseResponse AskLeave(1:sign.AskLeaveRequest req)(api.post="/sign/leave/")
    sign.MonthSignResponse GetMonthSign(1:sign.GetMonthSignRequest req)(api.get="/sign/month/")
    sign.MonthSignsResponse GetMonthSignByGid(1:sign.GetMonthSignsByGid req)(api.get="/sign/gmonth")
    sign.RecordsResponse GetAllRecord(1:sign.GetAllRecordRequest req)(api.get="/sign/recordlist/")
    sign.RecordsResponse GetUserRecord(1:sign.GetUserRecordRequest req)(api.get="/sign/record/")
    sign.BaseResponse SignPosAdd(1:sign.SignPosAddRequest req)(api.post="/sign/pos/")
    sign.BaseResponse SignPosDel(1:sign.SignPosDelRequest req)(api.delete="/sign/pos/")
    sign.GetSignPosResponse GetSignPos(1:sign.GetSignPosRequest req)(api.get="/sign/pos/")

    lottery.BaseResponse ActivityAdd(1:lottery.ActivityAddRequest req)(api.post="/activity/add/")
    lottery.BaseResponse ActivityDel(1:lottery.ActivityDelRequest req)(api.delete="/activity/del/")
    lottery.BaseResponse ActivityUpdate(1:lottery.ActivityUpdateRequest req)(api.put="/activity/update/")
    lottery.ActivitysResponse GetActivityByGid(1:lottery.GetActivityByGidRequest req)(api.get="/activity/gid/")
    lottery.ActivitysResponse GetAllActivity(1:lottery.GetAllActivityRequest req)(api.get="/activity/list/")
    lottery.ActivityResponse GetActivityById(1:lottery.GetActivityByIdRequest req)(api.get="/activity/id/")

    lottery.BaseResponse PrizeAdd(1:lottery.PrizeAddRequest req)(api.post="/prize/add/")
    lottery.BaseResponse PrizeDel(1:lottery.PrizeDelRequest req)(api.delete="/prize/del/")
    lottery.BaseResponse PrizeUpdate(1:lottery.PrizeUpdateRequest req)(api.put="/prize/update/")
    lottery.PrizesResponse GetPrizeByAid(1:lottery.GetPrizeByAidRequest req)(api.get="/prize/aid/")
    lottery.PrizeResponse GetPrizeById(1:lottery.GetPrizeByIdRequest req)(api.get="/prize/id/")
    lottery.ChooseResponse Choose(1:lottery.ChooseRequest req)(api.get="/prize/choose/")

    lottery.OrdersResponse GetUserOrder(1:lottery.GetUserOrderRequest req)(api.get="/order/uid/")
    lottery.OrdersResponse GetAllOrder(1:lottery.GetAllOrderRequest req)(api.get="/order/list/")

}