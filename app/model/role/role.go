package role

import "github.com/casbin/casbin/v2"

type RBAC struct {
	Enforcer *casbin.Enforcer `inject:""`
}
