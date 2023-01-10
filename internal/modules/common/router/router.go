package router

import (
	"api/internal/modules/common/controller"

	"github.com/gogf/gf/v2/net/ghttp"
)

func CommonBindController(group *ghttp.RouterGroup) {
	// ...
	group.Group("/common", func(group *ghttp.RouterGroup) {
		group.Bind(controller.Message)
		group.Bind(controller.Upload)
	})
}
