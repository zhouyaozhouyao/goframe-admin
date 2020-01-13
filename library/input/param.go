package input

import (
	"gadmin/library/base"

	"github.com/gogf/gf/os/glog"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/net/ghttp"
)

type ListRequest struct {
	Page    int                    `json:"page"`    // 当前页
	Limit   int                    `json:"limit"`   // 每页显示条数
	Where   map[string]interface{} `json:"where"`   // 查询条件
	OrderBy map[string]interface{} `json:"orderBy"` // 排序
}

// JSONToStruct 解析json参数并转换 pointer传指针地址
func JSONToStruct(r *ghttp.Request, pointer interface{}) {
	j, err := r.GetJson()
	if err != nil {
		base.Error(r, 3002)
		return
	}

	if j == nil {
		return
	}

	if err := j.ToStruct(pointer); err != nil {
		base.Error(r, 3002)
		return
	}
}

// StringToJSON 转换字符串为json对象
func StringToJSON(r *ghttp.Request, pointer interface{}) {
	if ok := r.GetRequest("param"); ok == nil {
		return
	}

	j, err := gjson.DecodeToJson([]byte(gconv.String(r.GetRequest("param"))))
	if err != nil {
		glog.Error(g.Map{"msg": "解析参数异常", "error": err})
		base.Error(r, 3001)
		return
	}

	if err := j.ToStruct(pointer); err != nil {
		base.Error(r, 3003)
		return
	}
}

// ListParams 列表页参数接收处理
func ListParams(r *ghttp.Request) ListRequest {
	var request = ListRequest{}
	JSONToStruct(r, &request)
	return request
}
