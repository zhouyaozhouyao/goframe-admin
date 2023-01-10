package admin

import (
	"github.com/gogf/gf/v2/frame/g"

	"api/internal/model"
	"api/internal/model/entity"
)

// DeptListReq 组织架构列表
type DeptListReq struct {
	g.Meta   `path:"/dept/list" method:"get" summary:"组织架构列表" tags:"组织架构" dc:"组织架构列表"`
	DeptName string `json:"deptName" dc:"部门名称"`
	Status   string `json:"status" dc:"状态"`
}

// DeptListRes 组织架构列表
type DeptListRes struct {
	g.Meta `mime:"application/json" dc:"响应"`
	List   []*entity.Dept `json:"list" dc:"列表"`
}

// DeptCreateReq 创建组织架构
type DeptCreateReq struct {
	g.Meta   `path:"/dept/create" method:"POST" summary:"创建组织架构" tags:"组织架构" description:"创建组织架构"`
	DeptName string `json:"deptName" v:"required#组织名称不能为空" dc:"组织名称"`
	ParentId int    `json:"parentId" v:"required#父级不能为空" dc:"父级id"`
	OrderNum int    `json:"orderNum" v:"required#排序不能为空" dc:"排序"`
	Leader   string `json:"leader" dc:"负责人"`
	Phone    string `json:"phone" dc:"联系电话"`
	Email    string `json:"email" v:"email#邮箱格式不正确" dc:"邮箱"`
	Status   int    `json:"status" v:"required#状态不能为空" dc:"状态"`
}

// DeptCreateRes 创建组织架构
type DeptCreateRes struct {
	g.Meta `mime:"application/json" dc:"响应"`
}

// DeptUpdateReq 更新组织架构
type DeptUpdateReq struct {
	g.Meta   `path:"/dept/update" method:"PUT" summary:"更新组织架构" tags:"组织架构" description:"更新组织架构"`
	DeptId   int    `p:"deptId" v:"required#deptId不能为空"`
	DeptName string `json:"deptName" v:"required#组织名称不能为空" dc:"组织名称"`
	ParentId int    `json:"parentId" v:"required#父级不能为空" dc:"父级id"`
	OrderNum int    `json:"orderNum" v:"required#排序不能为空" dc:"排序"`
	Leader   string `json:"leader" dc:"负责人"`
	Phone    string `json:"phone" dc:"联系电话"`
	Email    string `json:"email" v:"email#邮箱格式不正确" dc:"邮箱"`
	Status   int    `json:"status" v:"required#状态不能为空" dc:"状态"`
}

// DeptUpdateRes 更新组织架构
type DeptUpdateRes struct {
	g.Meta `min:"application/json" description:"响应数据"`
}

// DeptDeleteReq 删除组织架构
type DeptDeleteReq struct {
	g.Meta `path:"/dept/delete" method:"DELETE" summary:"删除组织架构" tags:"组织架构" description:"删除组织架构"`
	DeptId int `p:"deptId" v:"required#deptId不能为空"`
}

// DeptDeleteResp 删除组织架构
type DeptDeleteResp struct {
	g.Meta `min:"application/json" description:"响应数据"`
}

// DeptInfoReq 组织架构详情
type DeptInfoReq struct {
	g.Meta `path:"/dept/info" method:"GET" summary:"获取组织架构详情" tags:"组织架构" description:"获取组织架构详情"`
	DeptId int `json:"deptId" v:"required#deptId不能为空"`
}

// DeptInfoRes 组织架构详情
type DeptInfoRes struct {
	g.Meta `min:"application/json" description:"响应数据"`
	Info   *entity.Dept `json:"info" dc:"组织架构详情"`
}

// DeptTreeReq 组织架构树
type DeptTreeReq struct {
	g.Meta `path:"/dept/tree" method:"GET" summary:"获取组织架构树" tags:"组织架构" description:"获取组织架构树"`
}

// DeptTreeRes 组织架构树
type DeptTreeRes struct {
	g.Meta `min:"application/json" description:"响应数据"`
	Tree   []*model.DeptTreeOutput `json:"tree" dc:"组织架构树"`
}

// DeptDeleteRes 删除组织架构
type DeptDeleteRes struct {
	g.Meta `mime:"application/json" dc:"响应"`
}
