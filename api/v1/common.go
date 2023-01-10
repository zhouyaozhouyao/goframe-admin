package v1

type BaseCreateRes struct {
	Code    int         `json:"code" dc:"状态码"`
	Message string      `json:"message" dc:"消息提示"`
	Data    interface{} `json:"data" dc:"数据"`
}

type PageReq struct {
	DateRange []string `json:"dateRange" dc:"日期范围"`
	PageNum   int      `json:"pageNum" d:"1" dc:"当前页码"`
	PageSize  int      `json:"pageSize" d:"10" dc:"每页条数"`
	OrderBy   string   `json:"orderBy" d:"id desc" dc:"排序字段"`
}

type ListRes struct {
	CurrentPage int `json:"currentPage" dc:"当前页码"`
	Total       int `json:"total" dc:"总条数"`
}

// CommonActionReq 公共的创建者与编辑者
type CommonActionReq struct {
	CreateBy uint `json:"createBy" dc:"创建人"`
	UpdateBy uint `json:"updateBy" dc:"更新人"`
}

func (p *PageReq) ConditionOrPaginate() *PageReq {
	if p.PageNum == 0 {
		p.PageNum = 1
	}

	if p.PageSize == 0 {
		p.PageSize = 10
	}

	if p.OrderBy == "" {
		p.OrderBy = "created_at desc"
	}

	return p
}
