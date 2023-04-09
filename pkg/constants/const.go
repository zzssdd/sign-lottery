package constants

const (
	BaseServiceName    = "base"
	SignServiceName    = "sign"
	LotteryServiceName = "lottery"
	BaseTCPAddress     = "127.0.0.1:8889"
	SignTCPAddr        = "127.0.0.1:8890"
	LotteryTCPAddr     = "127.0.0.1:8891"
	MysqlDSN           = "yogen:yogen123@tcp(127.0.0.1:9910)/lottery?charset=utf8mb4&parseTime=True&loc=Local"
	RedisDSN           = "127.0.0.1:9911"
	EtcdDSN            = "127.0.0.1:9912"
	RabbitMqDSN        = "amqp://yogen:yogen123@127.0.0.1:5672/lottery"
	MaxUsernameLen     = 30
	MinPasswordLen     = 6
	MaxPasswordLen     = 30
	Version            = 1.0
	Mode               = "debug"
	SmtpAddr           = "smtp.qq.com"
	SmtpPort           = ":587"
	SmtpAuthKey        = "afhncxjpsukidgif"
	SmtpFromEmail      = "1654622146@qq.com"
)
