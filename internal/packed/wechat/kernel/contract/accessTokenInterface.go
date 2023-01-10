package contract

import (
	response2 "api/internal/packed/wechat/kernel/response"
	"net/http"

	"github.com/gogf/gf/v2/frame/g"
)

type AccessTokenInterface interface {
	// GetToken 获取token
	GetToken(refresh bool) (resToken *response2.TokenResponse, err error)
	Refresh() AccessTokenInterface
	ApplyToRequest(request *http.Request, requestOptions *g.Map) (*http.Request, error)
}
