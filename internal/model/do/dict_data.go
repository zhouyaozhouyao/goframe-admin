// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DictData is the golang structure of table data_dict_data for DAO operations like Where/Data.
type DictData struct {
	g.Meta    `orm:"table:data_dict_data, do:true"`
	DictCode  interface{} // 字典编码
	DictSort  interface{} // 字典排序
	DictLabel interface{} // 字典标签
	DictValue interface{} // 字典键值
	DictType  interface{} // 字典类型
	CssClass  interface{} // 样式属性（其他样式扩展）
	ListClass interface{} // 表格回显样式
	IsDefault interface{} // 是否默认（1是 0否）
	Status    interface{} // 状态（0正常 1停用）
	CreateBy  interface{} // 创建者
	UpdateBy  interface{} // 更新者
	Remark    interface{} // 备注
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 修改时间
}
