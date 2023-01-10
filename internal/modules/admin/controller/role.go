package controller

import (
	"context"

	"api/api/v1/admin"
	"api/internal/service"
)

type cRole struct {
}

var (
	Role = cRole{}
)

// Create 创建
func (c *cRole) Create(ctx context.Context, req *admin.RoleCreateReq) (res *admin.RoleCreateRes, err error) {
	err = service.Role().Create(ctx, req)
	return
}

func (c *cRole) Update(ctx context.Context, req *admin.RoleUpdateReq) (res *admin.RoleCreateRes, err error) {
	err = service.Role().Update(ctx, req)
	return
}

func (c *cRole) Delete(ctx context.Context, req *admin.RoleDeleteReq) (res *admin.RoleDeleteRes, err error) {
	err = service.Role().Delete(ctx, req.Id)
	return
}

func (c *cRole) List(ctx context.Context, req *admin.RoleListReq) (res *admin.RoleListRes, err error) {
	res, err = service.Role().GetRoleListSearch(ctx, req)
	return
}

func (c *cRole) Menu(ctx context.Context, req *admin.RoleMenuReq) (res *admin.RoleMenuRes, err error) {
	res = new(admin.RoleMenuRes)
	res.Menu, err = service.Menu().GetMenuList(ctx)
	return
}

func (c *cRole) Permission(ctx context.Context, req *admin.RolePermissionReq) (res *admin.RolePermissionRes, err error) {
	res = new(admin.RolePermissionRes)
	// 获取角色信息
	res.Role, err = service.Role().Info(ctx, req.Id)
	// 获取角色权限
	res.MenuIds, err = service.Role().GetFilteredNamedPolicy(ctx, req.Id)
	return
}
