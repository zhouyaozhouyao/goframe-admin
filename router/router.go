package router

import (
	"gadmin/app/api/admin/login"
	"gadmin/app/api/admin/user"
	"gadmin/app/middleware/permission"
	"gadmin/app/middleware/token"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()

	s.Group("/v1", func(group *ghttp.RouterGroup) {
		group.POST("/admin/loginSubmit", login.GfJWTMiddleware.LoginHandler)
		group.Group("/admin", func(group *ghttp.RouterGroup) {
			// 中间件检测token是否有效
			group.Middleware(token.Validator, permission.CasBinMiddleware)
			// 刷新token令牌
			group.GET("/refresh", login.GfJWTMiddleware.RefreshHandler)

			// 用户管理组
			group.ALL("/user", new(user.Controller))

		})
	})
}
