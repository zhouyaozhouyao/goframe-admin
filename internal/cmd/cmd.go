package cmd

import (
	commonRouter "api/internal/modules/common/router"
	"context"

	"api/internal/modules/admin/router"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Use()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				// group.Middleware(middleware.go.MiddlewareErrorHandler)
				router.AdminBindController(group)
				commonRouter.CommonBindController(group)
			})

			s.Run()
			return nil
		},
	}
)
