package model

import "api/internal/model/entity"

type DeptTreeOutput struct {
	*entity.Dept
	Children []*DeptTreeOutput `json:"children"`
}
