/**
 * 权限初始化设置
 * @email <994914376@qq.com>
 * @Author: zhouyao
 * @Date: 2020/1/11 1:54 下午
 */

package inject

import (
	"gadmin/app/model/rbac"
	"gadmin/library/base"
	"gadmin/library/e"

	"github.com/casbin/casbin/v2"
	"github.com/facebookgo/inject"
	"github.com/gogf/gf/net/ghttp"
)

// CasBinObj 注入对象
type CasBinObj struct {
	Common   *rbac.Common
	Enforcer *casbin.Enforcer
}

// Obj CasBin 实例变量
var Obj *CasBinObj

// 初始化 CasBinRBAC 配置
func init() {
	i := new(inject.Graph)
	var r *ghttp.Request
	var path = "config/rbac/rbac.conf"

	enforcer, err := casbin.NewEnforcer(path, false)
	if err != nil {
		base.Error(r, e.Error)
	}
	_ = i.Provide(&inject.Object{Value: enforcer})

	common := new(rbac.Common)
	_ = i.Provide(&inject.Object{Value: common})

	if err := i.Populate(); err != nil {
		base.Error(r, e.Error)
	}

	Obj = &CasBinObj{
		Enforcer: enforcer,
		Common:   common,
	}
	return
}

// LoadCasBinPolicyData 加载用户角色、角色权限等数据
func LoadCasBinPolicyData() error {
	m := Obj.Common

	// 注入用户与角色信息
	err := m.LoadPolicyData(1, "admin")
	if err != nil {
		return err
	}

	return nil
}
