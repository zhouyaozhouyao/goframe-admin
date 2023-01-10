package dto

type SmsResponse struct {
	Status    SmsStatus `json:"status"`
	RequestId string    `json:"requestId"`
}

type SmsStatus struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Phone   string `json:"phone"`
}

type Options struct {
	AccessKeyId     string
	AccessKeySecret string
}

const (
	AliYun         = "aliyun"
	WechatTemplate = "wechat_template"
)
