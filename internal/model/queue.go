package model

type PushInput struct {
	ServiceId   string `json:"serviceId"`   // 业务ID
	ServiceType int    `json:"serviceType"` // 业务类型 1：资金业务 2：银行业务 3、其它业务
	ApplyRemark string `json:"applyRemark"` // 申请的备注
	AuditRemark string `json:"auditRemark"` // 审核人员备注
	Result      int    `json:"result"`      // 审核结果 1：审核中 2：审核通过 3：审核拒绝
}
