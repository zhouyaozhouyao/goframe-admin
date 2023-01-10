package model

// LoginLogInput 登录日志写入参数
type LoginLogInput struct {
	Status    int
	Username  string
	Ip        string
	UserAgent string
	Msg       string
	Module    string
}
