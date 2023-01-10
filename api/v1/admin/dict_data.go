package admin

import (
	v1 "api/api/v1"
	"api/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type DictDataCreateOrUpdateBaseReq struct {
	Authorization string `json:"Authorization" in:"header" dc:"Bearer {{token}}"`
	DictSort      int    `json:"dictSort" v:"required#排序不能为空" dc:"字典排序"`
	DictLabel     string `json:"dictLabel" v:"required#标签不能为空" dc:"字典标签"`
	DictValue     string `json:"dictValue" v:"required#键值不能为空" dc:"字典键值"`
	DictType      string `json:"dictType" v:"required#类型不能为空" dc:"字典类型"`
	IsDefault     int    `json:"isDefault" dc:"默认标识"`
	Status        int    `json:"status" dc:"状态"`
	Remark        string `json:"remark" dc:"备注信息"`
	v1.CommonActionReq
}

type DictDataCreateReq struct {
	g.Meta `path:"/dict/data/create" tags:"字典数据管理" method:"POST" summary:"新增数据" dc:"添加字典数据"`
	DictDataCreateOrUpdateBaseReq
}

type DictDataUpdateReq struct {
	g.Meta   `path:"/dict/data/update" tags:"字典数据管理" method:"PUT" summary:"编辑数据" dc:"编辑字典数据"`
	DictCode int `json:"dict_code" v:"required#字典数据ID不能为空" dc:"字典数据ID"`
	DictDataCreateOrUpdateBaseReq
}

type DictDataDeleteReq struct {
	g.Meta `path:"/dict/data/delete" tags:"字典数据管理" method:"DELETE" summary:"删除数据" dc:"删除字典数据"`
	Ids    []int `json:"ids" v:"required#字典数据ID不能为空" dc:"字典数据ID"`
}

type DictDataListReq struct {
	g.Meta `path:"/dict/data/list" tags:"字典数据管理" method:"GET" summary:"字典数据列表" dc:"字典数据列表"`
	v1.PageReq
	DictType  string `json:"dictType" v:"required#字典类型不能为空" dc:"字典类型"`
	DictLabel string `json:"dictLabel" dc:"字典标签"`
	Status    int    `json:"status" dc:"字典类型"`
}

type DictDataListRes struct {
	g.Meta     `mime:"application/json" dc:"响应"`
	v1.ListRes                    // 总数
	List       []*entity.DictData `json:"list"`
}

type DictDataShowReq struct {
	g.Meta `path:"/dict/data/show" tags:"字典数据管理" method:"GET" summary:"字典数据详情" dc:"字典数据详情"`
	Id     int `json:"id" v:"required#字典数据ID不能为空" dc:"字典数据ID"`
}

type DictDataShowRes struct {
	g.Meta `mime:"application/json" dc:"响应"`
	entity.DictData
}

type DictDataRes struct {
}
