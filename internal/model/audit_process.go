package model

type AuditProcessDetailItem struct {
	DepartmentId int   `json:"departmentId" dc:"部门ID"`
	Procedure    int   `json:"procedure" dc:"步骤"`
	UserId       []int `json:"userId" dc:"用户ID组"`
}
