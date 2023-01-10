package service

import (
	"api/internal/packed/wechat"
	"api/internal/packed/wechat/kernel"
	"api/internal/packed/wechat/miniProgram"
	"context"
)

var MiniProgramApp *miniProgram.MiniProgram

func NewMiniProgramService(ctx context.Context, conf *wechat.Configuration) (*miniProgram.MiniProgram, error) {
	var cache kernel.CacheInterface
	cache = kernel.NewRedisClient()
	// 进行初始化
	app, err := miniProgram.NewMiniProgram(ctx, &miniProgram.UserConfig{
		AppId:  conf.MiniProgram.AppID,
		Secret: conf.MiniProgram.Secret,
		Cache:  cache,
	})
	return app, err
}
