// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AuditRecordDao is the data access object for table log_audit_record.
type AuditRecordDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns AuditRecordColumns // columns contains all the column names of Table for convenient usage.
}

// AuditRecordColumns defines and stores column names for table log_audit_record.
type AuditRecordColumns struct {
	Id          string // ID编号
	AuditId     string // 审核ID
	ServiceType string // 业务类型 1. 采购审核
	ServiceId   string // 业务ID
	AuditorId   string // 审核人ID
	AuditRemark string // 审核备注
	AuditStep   string // 审核流程
	Action      string // 动作 1 通过 0 拒绝
	CreatedAt   string // 生成时间
}

// auditRecordColumns holds the columns for table log_audit_record.
var auditRecordColumns = AuditRecordColumns{
	Id:          "id",
	AuditId:     "audit_id",
	ServiceType: "service_type",
	ServiceId:   "service_id",
	AuditorId:   "auditor_id",
	AuditRemark: "audit_remark",
	AuditStep:   "audit_step",
	Action:      "action",
	CreatedAt:   "created_at",
}

// NewAuditRecordDao creates and returns a new DAO object for table data access.
func NewAuditRecordDao() *AuditRecordDao {
	return &AuditRecordDao{
		group:   "default",
		table:   "log_audit_record",
		columns: auditRecordColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AuditRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AuditRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AuditRecordDao) Columns() AuditRecordColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AuditRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AuditRecordDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AuditRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
