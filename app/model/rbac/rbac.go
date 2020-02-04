/**
 *
 * @email <994914376@qq.com>
 * @Author: zhouyao
 * @Date: 2020/1/13 3:24 下午
 */

package rbac

import (
	"gadmin/app/model/role"
	"gadmin/app/model/users"

	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/os/glog"

	"github.com/casbin/casbin/v2"
)

// Common 定义全局对象
type Common struct {
	Enforcer *casbin.Enforcer `inject:""`
}

// LoadPolicyData 注入权限策略
func (a *Common) LoadPolicyData(id int, username string) error {
	// 查询用户、角色数据
	userResult, err := users.Model.M.As("u").
		InnerJoin("user_role as ur", "u.id = ur.user_id").
		LeftJoin("role as r", "ur.role_id = r.id").
		Fields("u.username, ur.role_id, r.name").
		FindAll("u.id = ?", 1)

	if err != nil {
		glog.Error("查询用户、角色数据错误", err)
		return err
	}

	// 先清除掉已有的权限数据
	_, _ = a.Enforcer.DeleteRolesForUser(username)

	// roleID 保存角色id切片
	var roleID []uint

	// 注册用户、角色 到CasBin配置中
	for _, v := range userResult.List() {
		roleID = append(roleID, gconv.Uint(v["role_id"]))
		_, _ = a.Enforcer.AddRoleForUser(gconv.String(v["username"]), gconv.String(v["name"]))
	}

	// 查询角色与路由对应关系
	roleMenuResult, err := role.Model.M.As("r").
		InnerJoin("role_menu as rm", "r.id = rm.role_id").
		LeftJoin("menu as m", "rm.menu_id = m.id").
		Fields("r.name as role_name, m.name, m.path, m.method").
		FindAll("r.id in (?)", roleID)

	if err != nil {
		glog.Error("查询角色与路由错错", err)
		return err
	}
	for _, v := range roleMenuResult.List() {
		if v["path"] == "" || v["method"] == "" {
			continue
		}
		_, _ = a.Enforcer.AddPermissionForUser(gconv.String(v["role_name"]), gconv.String(v["path"]), gconv.String(v["method"]))
	}

	return nil
}
