package kernel

import (
	"api/internal/library/liberr"
	"context"
	"net/http"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/net/gclient"
)

type BaseClient struct {
	Ctx    context.Context
	App    *ApplicationInterface
	Token  *AccessToken
	client *gclient.Client
}

func NewBaseClient(ctx context.Context, app *ApplicationInterface, token *AccessToken) (*BaseClient, error) {
	// 读取配置
	config := (*app).GetContainer().GetConfig()
	if token == nil {
		token = (*app).GetAccessToken()
	}

	// 注册中间件
	httpRequest := gclient.New().SetPrefix((*config)["base_uri"].(string))
	if (*app).GetAccessToken() != nil {
		httpRequest.Use(func(c *gclient.Client, r *http.Request) (resp *gclient.Response, err error) {
			r, _ = (*app).GetAccessToken().ApplyToRequest(r, config)
			//r.URL.Query().Set("access_token", "123456")
			return c.Next(r)
		})
	}

	client := &BaseClient{
		App:    app,
		Ctx:    ctx,
		Token:  token,
		client: httpRequest,
	}
	return client, nil
}

func (base *BaseClient) HttpGet(url string, data interface{}, result interface{}) (response interface{}, err error) {
	err = g.Try(base.Ctx, func(ctx context.Context) {
		//resp, err := g.Client().SetPrefix("https://api.weixin.qq.com").Get(base.Ctx, url, data)
		//err = base.client.GetVar(base.Ctx, url, data).Scan(&result)
		err = base.client.GetVar(base.Ctx, url, data).Scan(&result)
		liberr.IsNil(base.Ctx, err, "http请求失败")
	})
	return result, err
}

//func (client *BaseClient) HttpGet(url string, data interface{}, result interface{}) (res interface{}, err error) {
//
//	//err = g.Try(b.Ctx, func(ctx context.Context) {
//	//	err = g.Client().GetVar(baseClient.Ctx, url, data).Scan(&result)
//	//	liberr.IsNil(baseClient.Ctx, err, "http请求失败")
//	//})
//	//return result, err
//}

func (base *BaseClient) HttpPost(url string, data interface{}, result interface{}) (res interface{}, err error) {
	//err = base.client.PostVar(base.Ctx, url, data).Scan(&result)
	err = base.client.ContentJson().PostVar(base.Ctx, url, data).Scan(&result)
	liberr.IsNil(base.Ctx, err, "http Post 请求失败")
	return result, err
}
