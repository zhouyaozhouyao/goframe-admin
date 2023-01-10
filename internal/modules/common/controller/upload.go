package controller

import (
	"api/api/v1/common"
	"api/internal/packed/upload"
	"context"

	"github.com/gogf/gf/v2/util/grand"

	"github.com/gogf/gf/v2/errors/gerror"

	"github.com/gogf/gf/v2/container/garray"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
)

type cUpload struct {
}

var Upload = cUpload{}

func (c *cUpload) Upload(ctx context.Context, req *common.FileUploadReq) (res *common.FileUploadRes, err error) {
	req.Path = gfile.Join(gfile.Pwd(), req.Path)
	ext := gfile.ExtName(req.File.Filename)

	// 检测文件后缀名是否正确
	extArr := garray.NewStrArrayFrom(g.SliceStr{"png", "jpg", "jpeg", "gif", "rar", "zip", "tar", "gz"})
	if !extArr.ContainsI(ext) {
		return &common.FileUploadRes{}, gerror.New("文件格式不正确")
	}

	// 重命名文件名称
	req.File.Filename = grand.Letters(32) + "." + ext
	fileName, _ := req.File.Save(req.Path)
	url, err := upload.Oss.Upload(ctx, req.Path, fileName)
	if err != nil {
		return &common.FileUploadRes{}, err
	}

	// 删除本地文件
	_ = gfile.Remove(req.Path + "/" + fileName)

	return &common.FileUploadRes{
		Url:  url,
		File: fileName,
	}, nil
}
