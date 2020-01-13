/**
 * 权限管控
 * @email <994914376@qq.com>
 * @Author: zhouyao
 * @Date: 2020/1/10 1:19 下午
 */

package rbac

import (
	"fmt"

	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/net/ghttp"

	"github.com/casbin/casbin/v2"
)

type Role struct {
	Enforcer *casbin.Enforcer `inject:""`
}

var Obj *Role

func TestRBAC(r *ghttp.Request) {
	e, _ := casbin.NewEnforcer("config/rbac/rbac.conf", "config/rbac/rbac.csv")
	fmt.Printf("RBAC test start\n")

	// 注册

	sub := "alice" // the user that wants to access a resource.
	obj := "data1" // the resource that is going to be accessed.
	act := "read"  // the operation that the user performs on the resource.

	ok, err := e.Enforce(sub, obj, act)

	if err != nil {
		// handle err
	}

	if ok == true {
		fmt.Println("找到数据了")
	} else {
		fmt.Println("没有找到")
	}

	roles, _ := e.GetRolesForUser(sub)
	g.Dump(roles)
}

func Test(r *ghttp.Request) {
	enforcer, err := casbin.NewEnforcer("config/rbac/rbac.conf", false)
	if err != nil {
		fmt.Println("导入错误")
	}

	// 查询用户信息
	//user, _ := users.Model.FindOne("username = ?", "admin")
	//// 查询中间表信息
	//userRole, _ := users.Model.FindAll("user_id = ?", user.Id)
	//role, := role2.FindAll()

	// 替换为数据库查出来的信息放进来就能使用了
	//_, e2 := enforcer.AddPolicy("zhouyao", "/test", "POST")
	//if e2 != nil {
	//	fmt.Println(e2)
	//	r.Exit()
	//}

	_, _ = enforcer.AddRoleForUser("admin1", "业务主管")
	_, _ = enforcer.AddPermissionForUser("业务主管", "/test", "POST")
	g.Dump(enforcer.GetGroupingPolicy()) // a.Enforcer.GetGroupingPolicy()
	sub := "admin"                       // the user that wants to access a resource.
	obj := r.RequestURI                  // the resource that is going to be accessed.
	act := r.Request.Method              // the operation that the user performs on the resource.

	ok, err := enforcer.Enforce(sub, obj, act)
	g.Dump(ok)
	if err != nil {
		// handle err
	}

	if ok == true {
		fmt.Println("找到数据了")
	} else {
		fmt.Println("没有找到")
		r.Exit()
	}

}

func (a *Role) LoadPolicy(roleId string) error {
	ok, err := a.Enforcer.AddPolicy(roleId, "/test", "GET")
	if err != nil {
		fmt.Println(err)
		return err
	}
	g.Dump(ok)
	return nil
}
