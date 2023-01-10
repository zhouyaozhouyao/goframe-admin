package consts

const (
	PageSize     = 10
	PlatFormCode = "001"
	SmsOk        = "OK"
	SmsFail      = "FAIL"

	// ------------------------------- 审核状态 --------------------------------------------

	AuditStatusPending = 1 // 审核中
	AuditStatusPass    = 2 // 审核通过
	AuditStatusRefuse  = 3 // 审核拒绝

	// ------------------------------- 员工状态 --------------------------------------------

	UserStatusDisable = 0 // 员工禁用
	UserStatusEnable  = 1 // 员工正常
)
