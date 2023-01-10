package controller

import (
	"api/api/v1/common"
	"api/internal/packed/wechat"
	wechatService "api/internal/packed/wechat/service"
	"api/internal/service"
	"context"

	"github.com/gogf/gf/v2/os/gcache"

	"github.com/gogf/gf/v2/frame/g"
)

type cMessage struct {
	cache *gcache.Cache
}

var Message = cMessage{}

func (c *cMessage) SendMessage(ctx context.Context, req *common.SendMessageReq) (res *common.ResponseRes, err error) {
	_, _ = service.Message().SendMessage(ctx, req)
	return res, err
}

func (c *cMessage) Auth(ctx context.Context, req *common.AuthReq) (res *common.AuthRes, err error) {

	wechatService.MiniProgramApp, _ = wechatService.NewMiniProgramService(ctx, &wechat.Configuration{
		MiniProgram: wechat.MiniProgramCfg{
			AppID:  g.Cfg().MustGet(ctx, "wechat.mini_program.default.appid").String(),
			Secret: g.Cfg().MustGet(ctx, "wechat.mini_program.default.secret").String(),
		},
	})

	res = new(common.AuthRes)
	res.Data, _ = wechatService.MiniProgramApp.Auth.Session(req.Code)
	return res, nil
}

func (c *cMessage) GetPhoneNumber(ctx context.Context, req *common.GetPhoneNumberReq) (res *common.GetPhoneNumberRes, err error) {
	wechatService.MiniProgramApp, _ = wechatService.NewMiniProgramService(ctx, &wechat.Configuration{
		MiniProgram: wechat.MiniProgramCfg{
			AppID:  g.Cfg().MustGet(ctx, "wechat.mini_program.default.appid").String(),
			Secret: g.Cfg().MustGet(ctx, "wechat.mini_program.default.secret").String(),
		},
	})
	res = new(common.GetPhoneNumberRes)
	res.Data, _ = wechatService.MiniProgramApp.PhoneNumber.GetUserPhoneNumber(req.Code)
	return res, nil
}
