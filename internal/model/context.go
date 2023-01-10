package model

type Context struct {
	Code string // 平台识
	User *ContextUser
}

type ContextUser struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	RealName string `json:"real_name"`
	Avatar   string `json:"avatar"`
}
