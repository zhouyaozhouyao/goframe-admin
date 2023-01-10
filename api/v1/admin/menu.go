package admin

import (
	"api/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

// MenuCreateReq 创建菜单请求参数
type MenuCreateReq struct {
	g.Meta        `path:"/menu/create" method:"POST" summary:"创建菜单" tags:"菜单管理" description:"创建菜单"`
	Authorization string `json:"Authorization" in:"header" dc:"Bearer {{ token }}"`
	MenuType      uint   `json:"menuType" v:"min:0|max:2#菜单类型最小值为:min|菜单类型最大值为:max"`
	Pid           uint   `json:"pid" v:"min:0"`
	Name          string `json:"name" v:"required#请填写规则名称"`
	Title         string `json:"title" v:"required|length:1,100#请填写标题|标题长度在:min到:max位"`
	Icon          string `json:"icon"`
	Weigh         int    `json:"weigh"`
	Condition     string `json:"condition"`
	Remark        string `json:"remark"`
	IsHide        uint   `json:"isHide"`
	Path          string `json:"path"`
	Redirect      string `json:"redirect"` // 路由重定向
	Roles         []uint `json:"roles"`    // 角色ids
	IsLink        uint   `json:"isLink"`
	IsIframe      uint   `json:"isIframe"`
	IsCached      uint   `json:"isCached"`
	IsAffix       uint   `json:"isAffix"`
	LinkUrl       string `json:"linkUrl"`
}

// MenuCreateRes 创建菜单响应
type MenuCreateRes struct {
}

// MenuUpdateReq 更新菜单请求参数
type MenuUpdateReq struct {
	g.Meta        `path:"/menu/update" method:"POST" summary:"更新菜单" tags:"菜单管理" description:"更新菜单"`
	Authorization string `json:"Authorization" in:"header" dc:"Bearer {{ token }}"`
	Id            uint   `json:"id" v:"required|min:1#请填写id|id最小值为:min"`
	MenuType      uint   `json:"menuType" v:"min:0|max:2#菜单类型最小值为:min|菜单类型最大值为:max"`
	Pid           uint   `json:"pid" v:"min:0"`
	Name          string `json:"name" v:"required#请填写规则名称"`
	Title         string `json:"title" v:"required|length:1,100#请填写标题|标题长度在:min到:max位"`
	Icon          string `json:"icon"`
	Weigh         int    `json:"weigh"`
	Condition     string `json:"condition"`
	Remark        string `json:"remark"`
	IsHide        uint   `json:"isHide"`
	Path          string `json:"path"`
	Redirect      string `json:"redirect"` // 路由重定向
	Roles         []uint `json:"roles"`    // 角色ids
	IsLink        uint   `json:"isLink"`
	IsIframe      uint   `json:"isIframe"`
	IsCached      uint   `json:"isCached"`
	IsAffix       uint   `json:"isAffix"`
	LinkUrl       string `json:"linkUrl"`
}

// MenuUpdateRes 更新菜单响应
type MenuUpdateRes struct {
}

// MenuDeleteReq 删除菜单请求参数
type MenuDeleteReq struct {
	g.Meta        `path:"/menu/delete" method:"DELETE" summary:"删除菜单" tags:"菜单管理" description:"删除菜单"`
	Authorization string `p:"Authorization" in:"header" dc:"Bearer {{token}}"`
	Ids           []int  `json:"ids" v:"required#请填写id"`
}

// MenuDeleteRes 删除菜单响应
type MenuDeleteRes struct {
}

// MenuListReq 菜单列表请求参数
type MenuListReq struct {
	g.Meta        `path:"/menu/list" method:"GET" summary:"菜单列表" tags:"菜单管理" description:"菜单列表"`
	Authorization string `p:"Authorization" in:"header" dc:"Bearer {{token}}"`
	Title         string `json:"title" dc:"菜单标题"`
	IsHide        uint   `json:"isHide" dc:"是否显示隐藏"`
}

type MenuListRes struct {
	g.Meta `mime:"application/json"`
	Menu   []*model.AuthRuleTreeOutput `json:"menu" dc:"菜单列表层级"`
}
