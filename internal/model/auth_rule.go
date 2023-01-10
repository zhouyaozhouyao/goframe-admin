package model

import "api/internal/model/entity"

type AuthRuleInfoRes struct {
	*entity.AuthRule
}

type AuthRuleTreeOutput struct {
	*AuthRuleInfoRes
	Children []*AuthRuleTreeOutput `json:"children"`
}
