package controller

import (
	"context"

	"api/internal/model"

	"api/api/v1/admin"
	"api/internal/service"
)

type cMenu struct {
}

var (
	Menu = cMenu{}
)

// Create 创建
func (c *cMenu) Create(ctx context.Context, req *admin.MenuCreateReq) (res *admin.MenuCreateRes, err error) {
	err = service.Menu().Create(ctx, req)
	return
}

// Update 更新
func (c *cMenu) Update(ctx context.Context, req *admin.MenuUpdateReq) (res *admin.MenuUpdateRes, err error) {
	err = service.Menu().Update(ctx, req)
	return
}

// List 列表
func (c *cMenu) List(ctx context.Context, req *admin.MenuListReq) (res *admin.MenuListRes, err error) {
	var list []*model.AuthRuleInfoRes
	list, err = service.Menu().GetMenuSearchList(ctx, req)
	// 递归生成层级菜单
	menu := service.Menu().GetMenuListTree(0, list)
	return &admin.MenuListRes{
		Menu: menu,
	}, err
}

// Delete 删除
func (c *cMenu) Delete(ctx context.Context, req *admin.MenuDeleteReq) (res *admin.MenuDeleteRes, err error) {
	err = service.Menu().DeleteMenuByIds(ctx, req.Ids)
	return
}
