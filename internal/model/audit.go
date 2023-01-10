package model

import "api/internal/model/entity"

type AuditCreateInput struct {
	*entity.Audit
}

type AuditUpdateInput struct {
	Action int
	*entity.Audit
}

type AuditDetailOutput struct {
	*entity.Audit
	AuditRecord interface{} `json:"auditRecord"`
}
