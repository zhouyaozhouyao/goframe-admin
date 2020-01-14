/**
 * 校验当前用户是否有访问由路权限
 * @email <994914376@qq.com>
 * @Author: zhouyao
 * @Date: 2020/1/14 1:25 下午
 */

package permission

import (
	"gadmin/app/middleware/inject"
	"gadmin/library/base"
	"gadmin/library/e"

	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/net/ghttp"
)

// CasBinMiddleware 检测当前用户是否具有访问权限
func CasBinMiddleware(r *ghttp.Request) {
	var username = "admin2"
	g.Dump(r.RequestURI)
	if ok, err := inject.Obj.Enforcer.Enforce(username, r.RequestURI, r.Method); err != nil {
		base.Error(r, e.Error)
	} else if !ok {
		base.Error(r, e.Forbidden)
	}

	r.Middleware.Next()
}
