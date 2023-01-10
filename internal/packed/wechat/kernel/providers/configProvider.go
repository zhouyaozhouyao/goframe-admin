package providers

import (
	"api/internal/packed/wechat/kernel"
)

func RegisterConfigProvider(app kernel.ApplicationInterface) *kernel.Config {
	return kernel.NewConfig(app.GetContainer().GetConfig())
}
