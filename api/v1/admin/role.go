package admin

import (
	v1 "api/api/v1"
	"api/internal/model"
	"api/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type RoleCreateReq struct {
	g.Meta        `path:"/role/create" method:"POST" summary:"创建角色" tags:"角色管理" description:"创建角色"`
	Authorization string `p:"Authorization" in:"header" dc:"Bearer {{token}}"`
	Name          string `json:"name" v:"required#角色名称不能为空" dc:"角色名称"`
	Status        uint   `json:"status" dc:"角色状态"`
	ListOrder     uint   `json:"listOrder" dc:"排序"`
	Remark        string `json:"remark" dc:"备注"`
	MenuIds       []uint `json:"menuIds" dc:"菜单id列表"`
}

type RoleCreateRes struct {
}

type RoleUpdateReq struct {
	g.Meta        `path:"/role/update" method:"PUT" summary:"更新角色" tags:"角色管理" description:"更新角色"`
	Authorization string `p:"Authorization" in:"header" dc:"Bearer {{token}}"`
	Id            uint   `json:"id" v:"required#角色id不能为空" dc:"角色id"`
	Name          string `json:"name" v:"required#角色名称不能为空" dc:"角色名称"`
	Status        uint   `json:"status" dc:"角色状态"`
	ListOrder     uint   `json:"listOrder" dc:"排序"`
	Remark        string `json:"remark" dc:"备注"`
	MenuIds       []uint `json:"menuIds" dc:"菜单id列表"`
}

type RoleUpdateRes struct {
}

type RoleDeleteReq struct {
	g.Meta        `path:"/role/delete" method:"DELETE" summary:"删除角色" tags:"角色管理" description:"删除角色"`
	Authorization string `p:"Authorization" in:"header" dc:"Bearer {{token}}"`
	Id            uint   `json:"id" v:"required#角色id不能为空" dc:"角色id"`
}

type RoleDeleteRes struct {
}

type RoleListReq struct {
	g.Meta        `path:"/role/list" method:"GET" summary:"获取角色列表" tags:"角色管理" description:"获取角色列表"`
	Authorization string `p:"Authorization" in:"header" dc:"Bearer {{token}}"`
	Name          string `json:"name" dc:"角色名称"`
	v1.PageReq
}

type RoleListRes struct {
	g.Meta `min:"application/json" description:"响应数据"`
	v1.ListRes
	List []*entity.Role `json:"list" dc:"角色列表"`
}

type RoleMenuReq struct {
	g.Meta `path:"/role/menu" method:"GET" summary:"角色菜单列表" tags:"角色管理" description:"获取角色菜单"`
}

type RoleMenuRes struct {
	g.Meta `min:"application/json" description:"响应数据"`
	Menu   []*model.AuthRuleInfoRes `json:"menu" dc:"菜单列表"`
}

type RolePermissionReq struct {
	g.Meta `path:"/role/permission" method:"GET" summary:"当前角色信息" tags:"角色管理" description:"获取角色权限"`
	Id     uint `json:"id" v:"required#角色id不能为空" dc:"角色id"`
}

type RolePermissionRes struct {
	g.Meta  `min:"application/json" description:"响应数据"`
	Role    *model.Info `json:"role" dc:"角色信息"`
	MenuIds []int       `json:"menuIds" dc:"菜单id列表"`
}
