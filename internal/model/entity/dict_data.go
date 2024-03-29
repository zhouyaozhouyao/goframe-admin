// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DictData is the golang structure for table dict_data.
type DictData struct {
	DictCode  int64       `json:"dict_code"  description:"字典编码"`
	DictSort  int         `json:"dict_sort"  description:"字典排序"`
	DictLabel string      `json:"dict_label" description:"字典标签"`
	DictValue string      `json:"dict_value" description:"字典键值"`
	DictType  string      `json:"dict_type"  description:"字典类型"`
	CssClass  string      `json:"css_class"  description:"样式属性（其他样式扩展）"`
	ListClass string      `json:"list_class" description:"表格回显样式"`
	IsDefault int         `json:"is_default" description:"是否默认（1是 0否）"`
	Status    int         `json:"status"     description:"状态（0正常 1停用）"`
	CreateBy  uint64      `json:"create_by"  description:"创建者"`
	UpdateBy  uint64      `json:"update_by"  description:"更新者"`
	Remark    string      `json:"remark"     description:"备注"`
	CreatedAt *gtime.Time `json:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updated_at" description:"修改时间"`
}
