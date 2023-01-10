package router

import (
	"api/internal/modules/admin/controller"
	"api/internal/modules/admin/middleware"
	adminService "api/internal/modules/admin/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

func AdminBindController(group *ghttp.RouterGroup) {
	// ...
	group.Group("/admin", func(group *ghttp.RouterGroup) {
		group.Bind(controller.Login)
		// 登录验证 token 是否有效
		_ = adminService.GfToken().Middleware(group)
		// 初始化上下文登录者信息
		group.Middleware(middleware.NewMiddleware().Ctx)
		group.Bind(
			// 菜单管理
			controller.Menu,
			// 角色管理
			controller.Role,
			// 组织架构
			controller.Dept,
			// 岗位管理
			controller.Post,
			// 字典类型管理
			controller.DictType,
			// 字典数据管理
			controller.DictData,
			// 用户管理
			controller.User,
			// 系统配置
			controller.Settings,
			// 审核类型配置
			controller.AuditProcess,
			// 审核流程列表
			controller.Audit,
		)
	})
}
