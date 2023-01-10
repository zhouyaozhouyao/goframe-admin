package common

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type FileUploadReq struct {
	g.Meta `path:"/upload" method:"POST" summary:"文件上传" tags:"公共" dc:"文件上传"`
	File   *ghttp.UploadFile `json:"file" v:"required#文件不能为空" dc:"文件"`
	Path   string            `json:"path" d:"runtime/temp" dc:"文件路径"`
}

type FileUploadRes struct {
	Url  string `json:"url" dc:"文件地址"`
	File string `json:"file" dc:"文件名称"`
}
