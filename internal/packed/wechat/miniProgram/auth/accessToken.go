package auth

import (
	"api/internal/packed/wechat/kernel"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type AccessToken struct {
	*kernel.AccessToken
	Ctx context.Context
}

func NewAccessToken(ctx context.Context, app *kernel.ApplicationInterface) (*AccessToken, error) {
	kernelToken, err := kernel.NewAccessToken(ctx, app)

	token := &AccessToken{
		AccessToken: kernelToken,
		Ctx:         ctx,
	}
	token.EndpointToGetToken = "https://api.weixin.qq.com/cgi-bin/token"
	token.OverrideGetCredentials()
	return token, err
}

func (accessToken *AccessToken) OverrideGetCredentials() {
	config := (*accessToken.App).GetConfig().Map

	accessToken.GetCredentials = func() *g.MapStrStr {
		return &g.MapStrStr{
			"grant_type": "client_credential",
			"appid":      gconv.String((*config)["appId"]),
			"secret":     gconv.String((*config)["secret"]),
		}
	}
}
