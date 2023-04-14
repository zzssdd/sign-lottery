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

	SignNotInPos  = 401
	SignIpUsed    = 402
	SignNotInTime = 403

	Error = 500
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

	401: "未在签到范围内签到",
	402: "每个手机只供一个用户签到",
	403: "不在签到时间段内",

	500: "错误",
}

func GetMsg(code int) string {
	return codeMsg[code]
}
