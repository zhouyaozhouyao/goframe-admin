package controller

import (
	"api/api/v1/admin"
	"api/internal/service"
	"context"
)

type cDictType struct {
}

var DictType = cDictType{}

// List 字典类型列表
func (c *cDictType) List(ctx context.Context, req *admin.DictTypeListReq) (res *admin.DictTypeListRes, err error) {
	res, err = service.DictType().List(ctx, req)
	return
}

func (c *cDictType) Create(ctx context.Context, req *admin.DictTypeCreateReq) (res *admin.DictTypeRes, err error) {
	err = service.DictType().Create(ctx, req)
	return
}

// Update 更新字典类型信息
func (c *cDictType) Update(ctx context.Context, req *admin.DictTypeUpdateReq) (res *admin.DictTypeRes, err error) {
	err = service.DictType().Update(ctx, req)
	return
}

// Delete 删除字典类型
func (c *cDictType) Delete(ctx context.Context, req *admin.DictTypeDeleteReq) (res *admin.DictTypeRes, err error) {
	err = service.DictType().Delete(ctx, req)
	return
}
