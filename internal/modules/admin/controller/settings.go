package controller

import (
	"api/api/v1/admin"
	"api/internal/service"
	"context"
)

type cSettings struct {
}

var Settings = new(cSettings)

func (c *cSettings) Update(ctx context.Context, req *admin.SettingUpdateReq) (res *admin.SettingUpdateRes, err error) {
	err = service.Settings().Update(ctx, req)
	return
}

func (c *cSettings) Detail(ctx context.Context, req *admin.SettingDetailReq) (res *admin.SettingDetailRes, err error) {
	res, err = service.Settings().Detail(ctx, req)
	return
}
