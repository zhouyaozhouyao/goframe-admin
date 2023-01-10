package response

import "api/internal/packed/wechat/kernel/response"

// SessionCodeResponse 登录凭证校验。
type SessionCodeResponse struct {
	response.MiniProgramResponse

	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
}
