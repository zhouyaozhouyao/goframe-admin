package admin

import (
	v1 "api/api/v1"
	"api/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type AuditBaseReq struct {
	Authorization string `json:"Authorization" in:"header" dc:"Bearer {{token}}"`
	PlatformId    string `json:"platformId" d:"001"    dc:"平台ID"`
	ServiceName   string `json:"serviceName"   dc:"业务名称"`
	ServiceType   int    `json:"serviceType"   dc:"业务类型 1. 采购审核"`
	ServiceId     string `json:"serviceId"     dc:"业务ID"`
	FinalStep     int    `json:"finalStep"     dc:"最终步骤"`
	CurrentStep   int    `json:"currentStep"   dc:"当前步骤"`
	IndexColumn   string `json:"indexColumn"   dc:"步骤与业务索引"`
	ApplyId       string `json:"applyId"       dc:"申请人ID"`
	AuditorId     string `json:"auditorId"     dc:"审核人ID"`
	AuditRemark   string `json:"auditRemark"   dc:"最后一次审核备注"`
	Remark        string `json:"remark"         dc:"申请审核备注"`
	Status        uint   `json:"status"         dc:"审核状态 1 审核中 2 已通过 3 已拒绝"`
}

type AuditCreateReq struct {
	g.Meta `path:"/audit/create" method:"POST" summary:"创建审核" tags:"审核" description:"创建审核记录"`
	AuditBaseReq
}

type AuditUpdateReq struct {
	g.Meta      `path:"/audit/update" method:"PUT" summary:"更新审核" tags:"审核" description:"审核操作，同意或拒绝"`
	Id          uint   `json:"id" dc:"ID编号"`
	Action      int    `json:"action" v:"required#审核操作不能为空" dc:"操作 2 同意 3 拒绝"`
	AuditRemark string `json:"auditRemark" v:"required#审核备注不能为空" dc:"审核备注"`
	AuditorId   int    `json:"auditorId" dc:"审核人ID"`
}

type AuditDeleteReq struct {
	g.Meta `path:"/audit/delete" method:"DELETE" summary:"删除审核" tags:"审核" description:"删除审核记录"`
	Ids    []uint `json:"ids" v:"required#ID不能为空" dc:"审核ID"`
}

type AuditListReq struct {
	g.Meta `path:"/audit/list" method:"GET" summary:"审核列表" tags:"审核" description:"审核列表"`
	v1.PageReq
	AuditUserId int `json:"auditUserId"`
}

type AuditDetailReq struct {
	g.Meta `path:"/audit/detail" method:"GET" summary:"审核详情" tags:"审核" description:"审核详情"`
	Id     int `json:"id" v:"required#ID不能为空" dc:"审核ID"`
}

type AuditListRes struct {
	g.Meta `min:"application/json" description:"响应数据"`
	v1.ListRes
	List []*entity.Audit `json:"list" dc:"列表"`
}

type AuditRes struct {
}

type AuditDetailRes struct {
	*entity.Audit
	AuditRecord interface{} `json:"auditRecord" dc:"审核记录"`
}
