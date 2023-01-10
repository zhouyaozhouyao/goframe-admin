package admin

import (
	v1 "api/api/v1"
	"api/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type CreateOrUpdateBaseReq struct {
	Authorization string `p:"Authorization" in:"header" dc:"Bearer {{token}}"`
	DictName      string `json:"dictName" v:"required#名称不能为空" dc:"字典名称"`
	DictType      string `json:"dictType" v:"required#类型不能为空" dc:"字典类型"`
	Status        uint   `json:"status"  v:"required|in:0,1#状态不能为空|状态只能为0或1" dc:"状态"`
	Remark        string `json:"remark" dc:"备注信息"`
	*v1.CommonActionReq
}

type DictTypeCreateReq struct {
	g.Meta `path:"/dict/type/create" tags:"字典类型管理" method:"POST" summary:"新增" dc:"添加类型"`
	CreateOrUpdateBaseReq
}

type DictTypeUpdateReq struct {
	g.Meta `path:"/dict/type/update" tags:"字典类型管理" method:"PUT" summary:"编辑" dc:"编辑类型信息"`
	Id     int `json:"id" `
}

type DictTypeDeleteReq struct {
	g.Meta `path:"/dict/type/delete" tags:"字典类型管理" method:"DELETE" summary:"删除" dc:"删除字典类型"`
	Ids    []int `json:"ids" v:"required#字典类型ID不能为空" dc:"岗位ID"`
}

type DictTypeListReq struct {
	g.Meta `path:"/dict/type/list" tags:"字典类型管理" method:"GET" summary:"字典列表" dc:"字典类型列表"`
	v1.PageReq
	DictName string `json:"dictName"` //字典名称
	DictType string `json:"dictType"` //字典类型
	Status   string `json:"status"`   //字典状态
}

type DictTypeListRes struct {
	g.Meta `mime:"application/json" dc:"响应"`
	// 总数
	v1.ListRes
	List []*entity.DictType `json:"list"`
}

type DictTypeRes struct {
}
