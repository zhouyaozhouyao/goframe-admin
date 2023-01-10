package middleware

import (
	"api/internal/model"
	"api/internal/modules/admin/consts"
	adminService "api/internal/modules/admin/service"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/net/ghttp"
)

type IMiddleware interface {
	Ctx(r *ghttp.Request)
	Auth(r *ghttp.Request)
}

type middlewareImpl struct {
}

var (
	middleService = middlewareImpl{}
)

func NewMiddleware() IMiddleware {
	return &middleService
}

// Ctx 自定义上下文对象
func (s *middlewareImpl) Ctx(r *ghttp.Request) {
	ctx := r.GetCtx()
	// 初始化登录用户信息
	data, err := adminService.GfToken().ParseToken(r)
	if err != nil {
		// 执行下一步请求逻辑
		r.Middleware.Next()
	}
	if data != nil {
		context := new(model.Context)
		// 后续根据登录者的信息来进行存储
		context.Code = consts.PlatFormCode
		err = gconv.Struct(data.Data, &context.User)
		if err != nil {
			g.Log("exception").Error(ctx, err)
			// 执行下一步请求逻辑
			r.Middleware.Next()
		}

		// 绑定参数添加默认操作人信息
		if r.Method == "POST" {
			r.SetParam("createBy", context.User.Id)
		}
		if r.Method == "PUT" {
			r.SetParam("updateBy", context.User.Id)
		}
		adminService.Context().Init(r, context)
	}
	// 执行下一步请求逻辑
	r.Middleware.Next()
}

func (s *middlewareImpl) Auth(r *ghttp.Request) {
	r.Middleware.Next()
}
