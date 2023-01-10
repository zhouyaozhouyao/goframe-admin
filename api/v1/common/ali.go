package common

import "github.com/gogf/gf/v2/frame/g"

// SendMessageReq 发送短信
type SendMessageReq struct {
	g.Meta `path:"/ali/send/message" method:"GET" summary:"发送短信" tags:"公共" dc:"发送短信"`
	Phone  string `json:"phone" v:"required#手机号不能为空" dc:"手机号"`
}

type ResponseRes struct {
	Code int         `json:"code" default:"50"`
	Data interface{} `json:"data"`
}

type AuthReq struct {
	g.Meta `path:"/ali/auth" method:"GET" summary:"微信登录" tags:"公共" dc:"阿里云授权"`
	Code   string `json:"code" v:"required#code不能为空" dc:"code"`
}

type AuthRes struct {
	Data interface{} `json:"data"`
}

type GetPhoneNumberReq struct {
	g.Meta `path:"/ali/getPhoneNumber" method:"GET" summary:"获取手机号" tags:"公共" dc:"阿里云授权"`
	Code   string `json:"code" v:"required#code不能为空" dc:"code"`
}

type GetPhoneNumberRes struct {
	Data interface{} `json:"data"`
}
