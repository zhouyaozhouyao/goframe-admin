package controller

import (
	"api/api/v1/admin"
	"api/internal/service"
	"context"
)

type cDictData struct {
}

var DictData = cDictData{}

func (c *cDictData) List(ctx context.Context, req *admin.DictDataListReq) (res *admin.DictDataListRes, err error) {
	res, err = service.DictData().List(ctx, req)
	return
}

func (c *cDictData) Create(ctx context.Context, req *admin.DictDataCreateReq) (res *admin.DictDataRes, err error) {
	err = service.DictData().Create(ctx, req)
	return res, err
}

func (c *cDictData) Update(ctx context.Context, req *admin.DictDataUpdateReq) (res *admin.DictDataRes, err error) {
	err = service.DictData().Update(ctx, req)
	return
}

func (c *cDictData) Delete(ctx context.Context, req *admin.DictDataDeleteReq) (res *admin.DictDataRes, err error) {
	err = service.DictData().Delete(ctx, req)
	return
}

func (c *cDictData) Show(ctx context.Context, req *admin.DictDataShowReq) (res *admin.DictDataShowRes, err error) {
	res, err = service.DictData().Show(ctx, req)
	return
}
