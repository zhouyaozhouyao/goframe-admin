package admin

import (
	v1 "api/api/v1"
	"api/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// PostListReq 岗位列表
type PostListReq struct {
	g.Meta `path:"/post/list" tags:"岗位管理" method:"GET" summary:"岗位列表" dc:"岗位列表"`
	// 页码
	v1.PageReq
	PostName string `json:"postName" dc:"岗位名称"`
	Status   string `json:"status" dc:"状态"`
}

// PostListRes 岗位列表响应
type PostListRes struct {
	g.Meta `mime:"application/json" dc:"响应"`
	// 总数
	v1.ListRes
	List []*entity.Post `json:"list"`
}

// PostCreateReq 创建岗位
type PostCreateReq struct {
	g.Meta   `path:"/post/create" tags:"岗位管理" method:"POST" summary:"添加岗位" dc:"添加岗位"`
	PostCode string `json:"postCode" v:"required#岗位编码不能为空" dc:"岗位编号"`
	PostName string `json:"postName" v:"required#岗位名称不能为空" dc:"岗位名称"`
	PostSort int    `json:"postSort" v:"required#岗位排序不能为空" dc:"岗位排序"`
	Status   uint   `json:"status" dc:"状态"`
	Remark   string `json:"remark" dc:"备注"`
	v1.CommonActionReq
}

// PostCreateRes 创建岗位响应
type PostCreateRes struct {
}

// PostUpdateReq 编辑岗位
type PostUpdateReq struct {
	g.Meta   `path:"/post/update" tags:"岗位管理" method:"PUT" summary:"编辑岗位" dc:"编辑岗位信息"`
	PostId   int64  `json:"postId" v:"required#岗位ID不能为空" dc:"岗位ID"`
	PostCode string `json:"postCode" v:"required#岗位编码不能为空" dc:"岗位编号"`
	PostName string `json:"postName" v:"required#岗位名称不能为空" dc:"岗位名称"`
	PostSort int    `json:"postSort" v:"required#岗位排序不能为空" dc:"岗位排序"`
	Status   uint   `json:"status" dc:"状态"`
	Remark   string `json:"remark" dc:"备注"`
	v1.CommonActionReq
}

type PostUpdateRes struct {
}

type PostDeleteReq struct {
	g.Meta `path:"/post/delete" tags:"岗位管理" method:"DELETE" summary:"删除岗位" dc:"删除岗位信息"`
	Ids    []int `json:"ids" v:"required#岗位ID不能为空" dc:"岗位ID"`
}

type PostDeleteRes struct {
}
