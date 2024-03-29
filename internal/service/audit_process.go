// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"api/api/v1/admin"
	"context"
)

type (
	IAuditProcess interface {
		Create(ctx context.Context, req *admin.AuditProcessCreateReq) (err error)
		Update(ctx context.Context, req *admin.AuditProcessUpdateReq) (err error)
		Delete(ctx context.Context, req *admin.AuditProcessDeleteReq) (err error)
		Detail(ctx context.Context, req *admin.AuditProcessDetailReq) (res *admin.AuditProcessDetailRes, err error)
		List(ctx context.Context, req *admin.AuditProcessListReq) (res *admin.AuditProcessListRes, err error)
	}
)

var (
	localAuditProcess IAuditProcess
)

func AuditProcess() IAuditProcess {
	if localAuditProcess == nil {
		panic("implement not found for interface IAuditProcess, forgot register?")
	}
	return localAuditProcess
}

func RegisterAuditProcess(i IAuditProcess) {
	localAuditProcess = i
}
