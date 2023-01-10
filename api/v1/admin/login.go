package admin

import (
	"api/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta   `path:"/login" tags:"登录/登出" method:"post" summary:"用户登录"`
	Username string `json:"username" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password" v:"required#密码不能为空" dc:"密码"`
}

// LoginRes 登录成功返回数据
type LoginRes struct {
	g.Meta      `mime:"application/json"`
	UserInfo    *model.UserLoginOutput `json:"userInfo" dc:"用户信息"`
	Token       string                 `json:"token" dc:"token"`
	MenuList    []interface{}          `json:"menuList" dc:"菜单列表"`
	Permissions []string               `json:"permissions" dc:"权限列表"`
}

type LogoutReq struct {
	g.Meta        `path:"/logout" tags:"登录/登出" method:"delete" summary:"退出登录"`
	Authorization string `p:"Authorization" in:"header" dc:"Bearer {{token}}" dc:"请求头token"`
}
type LogoutRes struct {
	g.Meta `mime:"application/json"`
}
