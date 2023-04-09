package errmsg

const (
	Success = 200

	CodeIsExpired        = 301
	CodeNotExist         = 302
	CodeIncorrect        = 303
	SendEmailFailed      = 304
	EmailOrPasswordError = 305

	Error = 500
)

var codeMsg = map[int]string{
	200: "成功！",

	301: "验证码已过期，请重新获取！",
	302: "请先获取验证码",
	303: "验证码不正确",
	304: "发送验证码失败，请重试！",
	305: "邮箱或密码错误，请重新输入",

	500: "错误",
}

func GetMsg(code int) string {
	return codeMsg[code]
}
