package controller

import (
	"api/api/v1/admin"
	"api/internal/model"
	adminService "api/internal/modules/admin/service"
	"api/internal/service"
	"context"

	"github.com/gogf/gf/v2/util/gconv"
)

type cAudit struct {
}

var Audit = cAudit{}

func (c *cAudit) Create(ctx context.Context, req *admin.AuditCreateReq) (res *admin.AuditRes, err error) {
	var in model.AuditCreateInput
	if err := gconv.Struct(req, &in); err != nil {
		return res, err
	}

	// 获取当前登录人ID
	in.ApplyId = gconv.String(adminService.Context().GetUserId(ctx))
	err = service.Audit().Create(ctx, &in)
	return
}

func (c *cAudit) Update(ctx context.Context, req *admin.AuditUpdateReq) (res *admin.AuditRes, err error) {
	req.AuditorId = gconv.Int(adminService.Context().GetUserId(ctx))
	err = service.Audit().Update(ctx, req)
	return
}

func (c *cAudit) Delete(ctx context.Context, req *admin.AuditDeleteReq) (res *admin.AuditRes, err error) {
	return
}

func (c *cAudit) List(ctx context.Context, req *admin.AuditListReq) (res *admin.AuditListRes, err error) {
	req.AuditUserId = gconv.Int(adminService.Context().GetUserId(ctx))
	res, err = service.Audit().List(ctx, req)
	return
}

func (c *cAudit) Detail(ctx context.Context, req *admin.AuditDetailReq) (res *admin.AuditDetailRes, err error) {
	res, err = service.Audit().Detail(ctx, req)

	return
}
