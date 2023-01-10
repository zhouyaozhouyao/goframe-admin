package auth

import (
	"api/internal/packed/wechat/kernel"
	"api/internal/packed/wechat/miniProgram/auth/response"
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/frame/g"
)

type Client struct {
	*kernel.BaseClient
	Ctx context.Context
}

func NewClient(ctx context.Context, app *kernel.ApplicationInterface) (*Client, error) {
	baseClient, _ := kernel.NewBaseClient(ctx, app, nil)
	return &Client{
		BaseClient: baseClient,
		Ctx:        ctx,
	}, nil
}

// Session 登录凭证校验。
func (c *Client) Session(code string) (*response.SessionCodeResponse, error) {
	result := &response.SessionCodeResponse{}
	// 读取配置
	config := (*c.App).GetConfig().Map
	param := g.Map{
		"appid":      gconv.String((*config)["appId"]),
		"secret":     gconv.String((*config)["secret"]),
		"js_code":    code,
		"grant_type": "authorization_code",
	}
	_, _ = c.HttpGet("/sns/jscode2session", param, result)

	return result, nil
}
