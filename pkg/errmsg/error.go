package errmsg

const (
	Success = 200

	CodeIsExpired       = 301
	CodeNotExist        = 302
	CodeIncorrect       = 303
	SendEmailFailed     = 304
	EmailHasExist       = 305
	PasswordError       = 306
	EmailNotExist       = 307
	NewEqualOld         = 308
	OldPassworError     = 309
	NameOrPasswordError = 310
	PasswordIsTooShort  = 311
	PasswordIsTooLong   = 312

	SignNotInPos  = 401
	SignIpUsed    = 402
	SignNotInTime = 403

	Error         = 500
	TokenNotExist = 501
	TokenIsError  = 502
	TokenExpired  = 503

	ChooseNoPrize     = 600
	PrizeIsNull       = 601
	NotInActivityTime = 602

	NoPreviledge = 700
)

var codeMsg = map[int]string{
	200: "成功！",

	301: "验证码已过期，请重新获取！",
	302: "请先获取验证码",
	303: "验证码不正确",
	304: "发送验证码失败，请重试！",
	305: "邮箱已被注册！",
	306: "邮箱或密码错误，请重新输入",
	307: "邮箱还未被注册",
	308: "新密码和旧密码相等",
	309: "旧密码输入错误",
	310: "管理员姓名或密码错误",
	311: "密码长度太短",
	312: "密码长度过长",

	401: "未在签到范围内签到",
	402: "每个手机只供一个用户签到",
	403: "不在签到时间段内",

	500: "错误",
	501: "请先登陆",
	502: "非法登陆",
	503: "请重新登陆",

	600: "很遗憾，您未中奖",
	601: "您来晚了，活动奖品已被抽完",
	602: "对不起，不在活动时间内",

	700: "对不起，您没有权限进行此操作",
}

func GetMsg(code int) string {
	return codeMsg[code]
}
