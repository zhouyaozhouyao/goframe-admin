package miniProgram

import (
	"api/internal/packed/wechat/kernel"
	"api/internal/packed/wechat/kernel/providers"
	"api/internal/packed/wechat/miniProgram/auth"
	"api/internal/packed/wechat/miniProgram/phoneNumber"
	"context"

	"github.com/gogf/gf/v2/util/gconv"
)

type MiniProgram struct {
	*kernel.ServiceContainer
	Config *kernel.Config

	AccessToken *auth.AccessToken
	Auth        *auth.Client
	PhoneNumber *phoneNumber.Client
}

type UserConfig struct {
	AppId        string `json:"appId"`
	Secret       string `json:"secret"`
	RefreshToken string `json:"refreshToken"`
	Cache        kernel.CacheInterface
}

func NewMiniProgram(ctx context.Context, config *UserConfig) (*MiniProgram, error) {
	// 配置文件转成map
	configMap := gconv.Map(config)
	container, _ := kernel.NewServiceContainer(ctx, &configMap)
	container.GetConfig()
	app := &MiniProgram{
		ServiceContainer: container,
	}
	// 注册全局配置中心
	app.Config = providers.RegisterConfigProvider(app)

	// 注册 Token 获取之后并缓存
	app.AccessToken, _ = auth.RegisterProvider(ctx, app)
	// 注册 Token 和 Auth
	app.Auth, _ = auth.RegisterAuthProvider(ctx, app)
	// 手机号
	app.PhoneNumber, _ = phoneNumber.RegisterProvider(ctx, app)
	return app, nil
}

func (app *MiniProgram) GetContainer() *kernel.ServiceContainer {
	return app.ServiceContainer
}

func (app *MiniProgram) GetConfig() *kernel.Config {
	return app.Config
}

func (app *MiniProgram) GetAccessToken() *kernel.AccessToken {
	return app.AccessToken.AccessToken
}
