package response

import "api/internal/packed/wechat/kernel/response"

type Watermark struct {
	Timestamp int    `json:"timestamp"`
	AppID     string `json:"appid,omitempty"`
}

type PhoneInfo struct {
	PhoneNumber     string     `json:"phoneNumber,omitempty"`
	PurePhoneNumber string     `json:"purePhoneNumber,omitempty"`
	CountryCode     string     `json:"countryCode,omitempty"`
	Watermark       *Watermark `json:"watermark,omitempty"`
}

type PhoneNumberResponse struct {
	*response.MiniProgramResponse

	ErrCode   int        `json:"errcode"`
	ErrMsg    string     `json:"errmsg"`
	PhoneInfo *PhoneInfo `json:"phone_info,omitempty"`
}
