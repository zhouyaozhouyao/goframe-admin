package controller

import (
	"api/api/v1/admin"
	"api/internal/model/entity"
	"api/internal/service"
	"context"
)

type cDept struct {
}

var (
	Dept = cDept{}
)

// List 组织架构列表
func (c *cDept) List(ctx context.Context, req *admin.DeptListReq) (res *admin.DeptListRes, err error) {
	res = new(admin.DeptListRes)
	res.List, err = service.Dept().GetList(ctx, req)
	return
}

func (c *cDept) Create(ctx context.Context, req *admin.DeptCreateReq) (res *admin.DeptCreateRes, err error) {
	err = service.Dept().Create(ctx, req)
	return
}

func (c *cDept) Update(ctx context.Context, req *admin.DeptUpdateReq) (res *admin.DeptUpdateRes, err error) {
	err = service.Dept().Update(ctx, req)
	return
}

func (c *cDept) Delete(ctx context.Context, req *admin.DeptDeleteReq) (res *admin.DeptDeleteRes, err error) {
	err = service.Dept().Delete(ctx, req)
	return
}

func (c *cDept) Info(ctx context.Context, req *admin.DeptInfoReq) (res *admin.DeptInfoRes, err error) {
	res = new(admin.DeptInfoRes)
	res.Info, err = service.Dept().Info(ctx, req)
	return
}

func (c *cDept) Tree(ctx context.Context, req *admin.DeptTreeReq) (res *admin.DeptTreeRes, err error) {
	var deptList []*entity.Dept
	deptList, err = service.Dept().GetList(ctx, &admin.DeptListReq{})
	if err != nil {
		return
	}
	res = new(admin.DeptTreeRes)
	res.Tree = service.Dept().Tree(0, deptList)
	return
}
