package settings

import (
	"api/api/v1/admin"
	"api/internal/dao"
	"api/internal/service"
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type sSettings struct {
}

func init() {
	service.RegisterSettings(New())
}

func New() *sSettings {
	return &sSettings{}
}

func (s *sSettings) Update(ctx context.Context, req *admin.SettingUpdateReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.Settings.Ctx(ctx).Where(dao.Settings.Columns().Name, req.Name).Data(req).OmitEmpty().Update()
	})
	return nil
}

func (s *sSettings) Detail(ctx context.Context, req *admin.SettingDetailReq) (res *admin.SettingDetailRes, err error) {
	err = dao.Settings.Ctx(ctx).Where(dao.Settings.Columns().Name, req.Name).Scan(&res)
	return res, err
}
