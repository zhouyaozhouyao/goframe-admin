package response

type BaseResponse struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type MiniProgramResponse struct {
	BaseResponse

	ResultCode string `json:"resultcode,omitempty"`
	ResultMSG  string `json:"resultmsg,omitempty"`

	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type WorkResponse struct {
	BaseResponse

	ErrCode int    `json:"errcode,omitempty"`
	ErrMsg  string `json:"errmsg,omitempty"`
}
