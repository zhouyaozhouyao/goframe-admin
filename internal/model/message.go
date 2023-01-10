package model

// SmsMessageInput 短信参数
type SmsMessageInput struct {
	Phone        string `json:"phone" dc:"手机号"`
	TemplateCode string `json:"templateCode" dc:"模板编码"`
	Code         string `json:"code" dc:"验证码"`
}
