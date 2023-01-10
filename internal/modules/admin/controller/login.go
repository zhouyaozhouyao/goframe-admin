package controller

import (
	v1 "api/api/v1/admin"
	"api/internal/library/libUtils"
	"api/internal/model"
	adminService "api/internal/modules/admin/service"
	"api/internal/service"
	"context"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/frame/g"
)

type cLogin struct {
}

var (
	Login = cLogin{}
)

// Login 用户登录
func (c *cLogin) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	var (
		user  *model.UserLoginOutput
		token string
	)
	// 获取当前登录者IP地址
	ip := libUtils.GetClientIp(ctx)
	userAgent := libUtils.GetUserAgent(ctx)

	// 检测账号密码是否正确
	user, err = service.User().GetAdminUserByUsernamePassword(ctx, req)
	if err != nil {
		// 保存登录失败的日志信息
		service.LoginLog().Invoke(ctx, &model.LoginLogInput{
			Status:    0,
			Username:  req.Username,
			Ip:        ip,
			UserAgent: userAgent,
			Msg:       err.Error(),
			Module:    "系统后台",
		})
		return
	}
	// 更新登录记录
	err = service.User().UpdateLoginInfo(ctx, user.Id, ip)
	// 保存登录成功日志信息
	service.LoginLog().Invoke(ctx, &model.LoginLogInput{
		Status:    1,
		Username:  req.Username,
		Ip:        ip,
		UserAgent: userAgent,
		Msg:       "登录成功",
		Module:    "系统后台",
	})
	key := gconv.String(user.Id) + "-" + gmd5.MustEncryptString(user.UserName) + gmd5.MustEncryptString(user.UserPassword)
	if g.Cfg("token").MustGet(ctx, "gfToken.multiLogin").Bool() {
		key = gconv.String(user.Id) + "-" + gmd5.MustEncryptString(user.UserName) + gmd5.MustEncryptString(user.UserPassword+ip+userAgent)
	}
	user.UserPassword = ""

	token, err = adminService.GfToken().GenerateToken(ctx, key, user)
	if err != nil {
		return
	}
	res = &v1.LoginRes{
		Token:    token,
		UserInfo: user,
	}
	// TODO 获取用户菜单数据
	return
}

// Logout 用户登出
func (c *cLogin) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	// 获取配置
	var options *model.TokenOptions
	_ = g.Cfg("token").MustGet(ctx, "gfToken").Struct(&options)
	err = service.GfToken(options).RemoveToken(ctx, service.GfToken(options).GetRequestToken(g.RequestFromCtx(ctx)))
	return
}
