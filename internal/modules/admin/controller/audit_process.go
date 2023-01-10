package controller

import (
	"api/api/v1/admin"
	"api/internal/service"
	"context"
)

type cAuditProcess struct {
}

var AuditProcess = cAuditProcess{}

func (c *cAuditProcess) Create(ctx context.Context, req *admin.AuditProcessCreateReq) (res *admin.AuditProcessRes, err error) {
	err = service.AuditProcess().Create(ctx, req)
	return
}

func (c *cAuditProcess) Update(ctx context.Context, req *admin.AuditProcessUpdateReq) (res *admin.AuditProcessRes, err error) {
	err = service.AuditProcess().Update(ctx, req)
	return
}

// Delete 删除审核类型
func (c *cAuditProcess) Delete(ctx context.Context, req *admin.AuditProcessDeleteReq) (res *admin.AuditProcessRes, err error) {
	err = service.AuditProcess().Delete(ctx, req)
	return
}

// Detail 审核类型详情
func (c *cAuditProcess) Detail(ctx context.Context, req *admin.AuditProcessDetailReq) (res *admin.AuditProcessDetailRes, err error) {
	res, err = service.AuditProcess().Detail(ctx, req)
	return
}

// List 审核类型列表
func (c *cAuditProcess) List(ctx context.Context, req *admin.AuditProcessListReq) (res *admin.AuditProcessListRes, err error) {
	res, err = service.AuditProcess().List(ctx, req)
	return
}
