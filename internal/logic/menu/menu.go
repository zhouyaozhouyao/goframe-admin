package menu

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/util/gconv"

	"api/api/v1/admin"
	"api/internal/dao"
	"api/internal/library/casbin"
	"api/internal/library/liberr"
	"api/internal/model"
	"api/internal/model/do"
	"api/internal/modules/admin/consts"
	adminService "api/internal/modules/admin/service"
	"api/internal/service"

	"github.com/gogf/gf/v2/database/gdb"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/errors/gerror"
)

type sMenu struct {
}

func init() {
	service.RegisterMenu(New())
}

func New() *sMenu {
	return &sMenu{}
}

// Create 添加菜单
func (s *sMenu) Create(ctx context.Context, in *admin.MenuCreateReq) (err error) {
	// 检测菜单是否存在
	if s.menuNameExists(ctx, in.Title, 0) {
		err = gerror.New("菜单名称已存在")
		return
	}
	// 开启事务
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			// 菜单数据
			data := do.AuthRule{
				Pid:       in.Pid,
				Name:      in.Name,
				Title:     in.Title,
				Icon:      in.Icon,
				Condition: in.Condition,
				Remark:    in.Remark,
				MenuType:  in.MenuType,
				Weigh:     in.Weigh,
				IsHide:    in.IsHide,
				Path:      in.Path,
				IsLink:    in.IsLink,
				IsIframe:  in.IsIframe,
				IsCached:  in.IsCached,
				Redirect:  in.Redirect,
				IsAffix:   in.IsAffix,
				LinkUrl:   in.LinkUrl,
			}
			ruleId, e := dao.AuthRule.Ctx(ctx).TX(tx).Data(data).InsertAndGetId()
			liberr.IsNil(ctx, e, "添加菜单失败")
			// 菜单绑定角色 暂无看到场景
			e = s.BindRoleRule(ctx, ruleId, in.Roles)
			liberr.IsNil(ctx, e, "添加菜单失败-绑定角色") // if len(in.Roles) > 0 {
		})
		return err
	})
	if err == nil {
		// 删除相关缓存
		adminService.Cache().Remove(ctx, consts.CacheSysAuthMenu)
	}
	return
}

// BindRoleRule 菜单绑定角色
func (s *sMenu) BindRoleRule(ctx context.Context, ruleId interface{}, roleIds []uint) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		enforcer, e := casbin.CabinEnforcer(ctx)
		liberr.IsNil(ctx, e)
		for _, roleId := range roleIds {
			_, err = enforcer.AddPolicy(fmt.Sprintf("%d", roleId), fmt.Sprintf("%d", ruleId), "All")
			liberr.IsNil(ctx, err)
		}
	})
	return
}

// Update 更新菜单信息
func (s *sMenu) Update(ctx context.Context, in *admin.MenuUpdateReq) (err error) {
	if s.menuNameExists(ctx, in.Title, in.Id) {
		err = gerror.New("菜单名称已存在")
		return
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			// 菜单数据
			data := do.AuthRule{
				Pid:       in.Pid,
				Name:      in.Name,
				Title:     in.Title,
				Icon:      in.Icon,
				Condition: in.Condition,
				Remark:    in.Remark,
				MenuType:  in.MenuType,
				Weigh:     in.Weigh,
				IsHide:    in.IsHide,
				Path:      in.Path,
				IsLink:    in.IsLink,
				IsIframe:  in.IsIframe,
				IsCached:  in.IsCached,
				Redirect:  in.Redirect,
				IsAffix:   in.IsAffix,
				LinkUrl:   in.LinkUrl,
			}
			_, e := dao.AuthRule.Ctx(ctx).TX(tx).Data(data).WherePri(in.Id).Update()
			liberr.IsNil(ctx, err, "更新菜单失败")
			e = s.UpdateRoleRule(ctx, in.Id, in.Roles)
			liberr.IsNil(ctx, e, "添加菜单失败")
		})
		return err
	})
	if err == nil {
		adminService.Cache().Remove(ctx, consts.CacheSysAuthMenu)
	}
	return
}

// UpdateRoleRule 更新菜单绑定角色
func (s *sMenu) UpdateRoleRule(ctx context.Context, ruleId uint, roleIds []uint) (err error) {
	return g.Try(ctx, func(ctx context.Context) {
		enforcer, e := casbin.CabinEnforcer(ctx)
		liberr.IsNil(ctx, e)
		// 删除旧权限
		_, e = enforcer.RemoveFilteredPolicy(1, gconv.String(ruleId))
		liberr.IsNil(ctx, e)
		// 添加新权限
		roleIdsStrArr := gconv.Strings(roleIds)
		for _, v := range roleIdsStrArr {
			_, e = enforcer.AddPolicy(v, gconv.String(ruleId), "All")
			liberr.IsNil(ctx, e)
		}
	})
}

// DeleteMenuByIds 删除菜单
func (s *sMenu) DeleteMenuByIds(ctx context.Context, ids []int) (err error) {
	var list []*model.AuthRuleInfoRes
	list, err = s.GetMenuList(ctx)
	if err != nil {
		return
	}
	childrenIds := make([]int, 0, len(list))
	for _, id := range ids {
		// 递归查询
		rules := s.FindSonParentId(list, gconv.Uint(id))
		for _, child := range rules {
			childrenIds = append(childrenIds, gconv.Int(child.Id))
		}
	}
	ids = append(ids, childrenIds...)
	// 开启事务
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		return g.Try(ctx, func(ctx context.Context) {
			// 删除所选菜单
			_, err = dao.AuthRule.Ctx(ctx).TX(tx).WhereIn(dao.AuthRule.Columns().Id, ids).Delete()
			liberr.IsNil(ctx, err, "删除菜单失败")
			// 删除权限
			enforce, e := casbin.CabinEnforcer(ctx)
			liberr.IsNil(ctx, e)
			for _, v := range ids {
				_, e = enforce.RemoveFilteredPolicy(1, gconv.String(v))
				liberr.IsNil(ctx, e)
			}
			// 删除相关缓存
			adminService.Cache().Remove(ctx, consts.CacheSysAuthMenu)
		})
	})
	return
}

// GetMenuList 获取所有菜单
func (s *sMenu) GetMenuList(ctx context.Context) (list []*model.AuthRuleInfoRes, err error) {
	cache := adminService.Cache()
	// 从缓存中获取
	iList := cache.GetOrSetFuncLock(ctx, consts.CacheSysAuthMenu, s.getMenuListFromDb, 0, consts.CacheSysAuthTag)
	if iList != nil {
		err = gconv.Struct(iList, &list)
		liberr.IsNil(ctx, err)
	}
	return
}

// FindSonParentId 递归查询子菜单
func (s *sMenu) FindSonParentId(list []*model.AuthRuleInfoRes, pid uint) []*model.AuthRuleInfoRes {
	children := make([]*model.AuthRuleInfoRes, 0, len(list))
	for _, v := range list {
		if v.Pid == pid {
			children = append(children, v)
			fChildren := s.FindSonParentId(list, v.Id)
			children = append(children, fChildren...)
		}
	}
	return children
}

// GetMenuSearchList 获取菜单搜索列表
func (s *sMenu) GetMenuSearchList(ctx context.Context, in *admin.MenuListReq) (out []*model.AuthRuleInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.AuthRule.Ctx(ctx).Where(dao.AuthRule.Columns().IsHide, in.IsHide)
		if in.Title != "" {
			m = m.WhereLike(dao.AuthRule.Columns().Title, in.Title)
		}
		err = m.Order("weigh desc, id asc").Scan(&out)
	})
	return
}

// GetMenuListTree 获取菜单转换成树形结构
func (s *sMenu) GetMenuListTree(pid uint, list []*model.AuthRuleInfoRes) []*model.AuthRuleTreeOutput {
	// 初始化保存变量
	tree := make([]*model.AuthRuleTreeOutput, 0, len(list))
	for _, menu := range list {
		if pid == menu.Pid {
			t := &model.AuthRuleTreeOutput{
				AuthRuleInfoRes: menu,
			}
			child := s.GetMenuListTree(menu.Id, list)
			if child != nil {
				t.Children = child
			}
			tree = append(tree, t)
		}
	}
	return tree
}

// getMenuListFromDb 从数据库中获取菜单列表
func (s *sMenu) getMenuListFromDb(ctx context.Context) (value interface{}, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		var v []*model.AuthRuleInfoRes
		err = dao.AuthRule.Ctx(ctx).Fields(model.AuthRuleInfoRes{}).Order("weigh desc, id asc").Scan(&v)
		liberr.IsNil(ctx, err, "获取菜单数据错误")
		value = v
	})
	return
}

// menuNameExists 检测菜单名称是否存在
func (s *sMenu) menuNameExists(ctx context.Context, title string, id uint) bool {
	m := dao.AuthRule.Ctx(ctx).Where(do.AuthRule{Title: title})
	if id != 0 {
		m = m.WhereNot(dao.AuthRule.Columns().Id, id)
	}
	c, err := m.Fields(dao.AuthRule.Columns().Id).One()
	if err != nil {
		g.Log("exception").Error(ctx, err)
		return false
	}

	return !c.IsEmpty()
}
