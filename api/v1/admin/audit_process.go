package admin

import (
	v1 "api/api/v1"
	"api/internal/model"
	"api/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type AuditProcessCreateOrUpdateBase struct {
	Authorization     string `json:"Authorization" in:"header" dc:"Bearer {{token}}"`
	Type              uint   `json:"type" d:"2"              dc:"审核类型：1 通用配置 2 部门负责人"`
	PlatformId        uint   `json:"platformId" d:"001"         dc:"平台ID"`
	ServiceName       string `json:"serviceName"        dc:"业务名称"`
	ServiceType       int    `json:"serviceType"        dc:"业务类型 1. 采购审核"`
	AuditDepartmentId int    `json:"auditDepartmentId" dc:"审核部门ID"`
	AuditUserId       string `json:"auditUserId"       dc:"用户ID（可为空，如果为空表明是这个部门下的权限）"`
	Procedure         int    `json:"procedure"           dc:"步骤编号"`
}

type AuditProcessCreateReq struct {
	g.Meta `path:"/audit/processes/create" method:"POST" summary:"创建审核类型" tags:"审核配置" dc:"创建审核类型"`
	AuditProcessCreateOrUpdateBase
}

type AuditProcessUpdateReq struct {
	g.Meta          `path:"/audit/processes/update" method:"PUT" summary:"更新审核类型" tags:"审核配置" dc:"更新审核类型"`
	ServiceType     int                             `json:"serviceType" v:"required#请输入业务类型" dc:"业务类型"`
	ServiceName     string                          `json:"serviceName" v:"required#请输入业务名称" dc:"业务名称"`
	AuditProcessArr []*model.AuditProcessDetailItem `json:"auditProcessArr" v:"required#请输入审核流程" dc:"审核流程数据"`
}

type AuditProcessDeleteReq struct {
	g.Meta      `path:"/audit/processes/delete" method:"DELETE" summary:"删除审核类型" tags:"审核配置" dc:"删除审核类型"`
	ServiceType int `json:"serviceType" v:"required#请输入业务类型" dc:"业务类型"`
	Process     int `json:"process" json:"process" v:"required#请输入步骤编号" dc:"步骤编号"`
}

type AuditProcessDetailReq struct {
	g.Meta      `path:"/audit/processes/detail" method:"GET" summary:"获取审核类型详情" tags:"审核配置" dc:"获取审核类型详情"`
	ServiceType int `json:"serviceType" v:"required#类型不能为空" dc:"审核类型"`
}

type AuditProcessListReq struct {
	g.Meta `path:"/audit/processes/list" method:"GET" summary:"获取审核类型列表" tags:"审核配置" dc:"获取审核类型列表"`
	v1.PageReq
}

type AuditProcessListRes struct {
	g.Meta `min:"application/json" description:"响应数据"`
	v1.ListRes
	List []*entity.ConfigAuditProcess `json:"list" dc:"列表"`
}

type AuditProcessRes struct {
}

type AuditProcessDetailRes struct {
	g.Meta `min:"application/json" description:"响应数据"`
	Lists  []*model.AuditProcessDetailItem `json:"lists" dc:"详情"`
}
