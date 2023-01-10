package admin

import (
	"api/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type SettingUpdateReq struct {
	g.Meta `path:"/settings/update" method:"PUT" summary:"更新系统设置" tags:"系统设置" dc:"更新系统设置"`
	Name   string      `json:"name" dc:"名称"`
	Value  interface{} `json:"value" dc:"值"`
}

type SettingDetailReq struct {
	g.Meta `path:"/settings/detail" method:"GET" summary:"获取系统设置详情" tags:"系统设置" dc:"获取系统设置详情"`
	Name   string `json:"name" v:"required#名称不能为空" dc:"名称"`
}

type SettingUpdateRes struct {
}

type SettingDetailRes struct {
	g.Meta `min:"application/json" description:"响应数据"`
	*entity.Settings
}
