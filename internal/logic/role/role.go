package role

import (
	"context"

	"api/api/v1/admin"
	"api/internal/dao"
	"api/internal/library/casbin"
	"api/internal/library/liberr"
	"api/internal/model"
	"api/internal/model/do"
	"api/internal/modules/admin/consts"
	adminService "api/internal/modules/admin/service"
	"api/internal/service"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/frame/g"
)

type sRole struct {
}

var Role = new(sRole)

func init() {
	service.RegisterRole(New())
}

func New() *sRole {
	return &sRole{}
}

// GetRoleListSearch 角色列表
func (s *sRole) GetRoleListSearch(ctx context.Context, in *admin.RoleListReq) (res *admin.RoleListRes, err error) {
	res = new(admin.RoleListRes)
	_ = g.Try(ctx, func(ctx context.Context) {
		m := dao.Role.Ctx(ctx)
		if in.Name != "" {
			m = m.WhereLike(dao.Role.Columns().Name, in.Name)
		}
		order := "list_order asc, id asc"
		if in.OrderBy != "" {
			order = in.OrderBy
		}

		res.Total, err = m.Count()
		liberr.IsNil(ctx, err, "获取角色数据失败")
		res.CurrentPage = in.PageNum
		if in.PageSize == 0 {
			in.PageSize = consts.PageSize
		}
		err = m.Page(res.CurrentPage, in.PageSize).Order(order).Scan(&res.List)
		liberr.IsNil(ctx, err, "获取数据失败")
	})
	return res, err
}

// Create 添加角色
func (s *sRole) Create(ctx context.Context, in *admin.RoleCreateReq) (err error) {
	// 检测角色名称是否存在
	if s.CheckByNameExists(ctx, in.Name, 0) {
		err = gerror.New("角色名称已存在")
		return
	}

	// 开启事务
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			roleId, e := dao.Role.Ctx(ctx).Data(in).InsertAndGetId()
			liberr.IsNil(ctx, e, "添加角色失败")
			// 添加角色权限
			e = s.AddRoleRule(ctx, in.MenuIds, roleId)
		})
		return err
	})

	return
}

// Update 更新角色与权限信息
func (s *sRole) Update(ctx context.Context, in *admin.RoleUpdateReq) (err error) {
	// 检测角色名称是否存在
	if s.CheckByNameExists(ctx, in.Name, in.Id) {
		err = gerror.New("角色名称已存在")
		return
	}

	// 开启事务
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			// 更新角色信息
			_, e := dao.Role.Ctx(ctx).Data(in).WherePri(in.Id).Update()
			liberr.IsNil(ctx, e, "更新失败")
			// 删除角色关联权限
			e = s.DelRoleRule(ctx, in.Id)
			liberr.IsNil(ctx, e)
			// 添加角色关联权限
			e = s.AddRoleRule(ctx, in.MenuIds, gconv.Int64(in.Id))
			liberr.IsNil(ctx, e)
			// 清除缓存
			adminService.Cache().Remove(ctx, consts.CacheSysRole)
		})
		return err
	})
	return
}

// Delete 删除角色
func (s *sRole) Delete(ctx context.Context, id uint) (err error) {
	// 开启事务
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			// 删除角色
			_, e := dao.Role.Ctx(ctx).WherePri(id).Delete()
			liberr.IsNil(ctx, e, "删除失败")
			// 删除角色关联权限
			e = s.DelRoleRule(ctx, id)
			liberr.IsNil(ctx, e)
			// 清除缓存
			adminService.Cache().Remove(ctx, consts.CacheSysRole)
		})
		return err
	})
	return
}

// CheckByNameExists 检测数据是否存在
func (s *sRole) CheckByNameExists(ctx context.Context, name string, id uint) bool {
	m := dao.Role.Ctx(ctx).Where(do.Role{Name: name})
	if id > 0 {
		m = m.WhereNot(dao.Role.Columns().Id, id)
	}
	c, err := m.One()
	if err != nil {
		g.Log("exception").Error(ctx, err)
		return false
	}
	return !c.IsEmpty()
}

// AddRoleRule 添加角色与权限关联
func (s *sRole) AddRoleRule(ctx context.Context, ruleIds []uint, roleId int64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		enforcer, e := casbin.CabinEnforcer(ctx)
		liberr.IsNil(ctx, e, "获取casbin失败")
		// 转成字符串数组
		ruleIdsStr := gconv.Strings(ruleIds)
		for _, v := range ruleIdsStr {
			// 添加角色权限
			_, err = enforcer.AddPolicy(gconv.String(roleId), v, "All")
			liberr.IsNil(ctx, err)
			// 清除缓存
			adminService.Cache().Remove(ctx, consts.CacheSysRole)
		}
	})
	return
}

// DelRoleRule 删除指定角色关联的权限
func (s *sRole) DelRoleRule(ctx context.Context, roleId uint) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		enforcer, e := casbin.CabinEnforcer(ctx)
		liberr.IsNil(ctx, e, "获取casbin失败")
		// 删除角色权限
		_, err = enforcer.RemoveFilteredPolicy(0, gconv.String(roleId))
		liberr.IsNil(ctx, err)
	})
	return
}

func (s *sRole) Info(ctx context.Context, id uint) (output *model.Info, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.Role.Ctx(ctx).WherePri(id).Scan(&output)
		liberr.IsNil(ctx, err, "获取信息失败")
	})
	return
}

func (s *sRole) GetFilteredNamedPolicy(ctx context.Context, id uint) (gSlice []int, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		enforcer, e := casbin.CabinEnforcer(ctx)
		liberr.IsNil(ctx, e, "获取casbin失败")
		// 获取角色权限
		gp := enforcer.GetFilteredNamedPolicy("p", 0, gconv.String(id))
		gSlice = make([]int, len(gp))
		for k, v := range gp {
			gSlice[k] = gconv.Int(v[1])
		}
	})
	return
}
