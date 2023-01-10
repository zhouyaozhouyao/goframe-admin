package dict_data

import (
	"api/api/v1/admin"
	"api/internal/dao"
	"api/internal/library/liberr"
	"api/internal/service"
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type sDictData struct {
}

func init() {
	service.RegisterDictData(New())
}

func New() *sDictData {
	return &sDictData{}
}

func (s *sDictData) Create(ctx context.Context, in *admin.DictDataCreateReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.DictData.Ctx(ctx).Data(in).OmitEmpty().Insert()
		liberr.IsNil(ctx, err, "添加字典数据失败")
	})
	return
}

func (s *sDictData) Update(ctx context.Context, in *admin.DictDataUpdateReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.DictData.Ctx(ctx).WherePri(in.DictCode).Data(in).OmitEmpty().Update()
		liberr.IsNil(ctx, err, "更新字典数据失败")
	})
	return
}

func (s *sDictData) Delete(ctx context.Context, in *admin.DictDataDeleteReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.DictData.Ctx(ctx).WhereIn(dao.DictData.Columns().DictCode, in.Ids).Delete()
		liberr.IsNil(ctx, err, "删除字典数据失败")
	})
	return
}

func (s *sDictData) Show(ctx context.Context, in *admin.DictDataShowReq) (res *admin.DictDataShowRes, err error) {
	err = dao.DictData.Ctx(ctx).WherePri(in.Id).Scan(&res)
	return res, err
}

func (s *sDictData) List(ctx context.Context, in *admin.DictDataListReq) (res *admin.DictDataListRes, err error) {
	// 初始化条件与页码信息
	in.ConditionOrPaginate()
	m := dao.DictData.Ctx(ctx).Where(dao.DictData.Columns().DictType, in.DictType)
	if in.Status > 0 {
		m = m.Where(dao.DictData.Columns().Status, in.Status)
	}
	if in.DictLabel != "" {
		m = m.Where(dao.DictData.Columns().DictLabel, in.DictLabel)
	}

	res.Total, err = m.Count()
	liberr.IsNil(ctx, err, "获取字典数据列表失败")

	res.CurrentPage = in.PageNum
	err = m.Page(in.PageNum, in.PageSize).Order(in.OrderBy).Scan(&res.List)
	liberr.IsNil(ctx, err, "获取字典类型列表失败")
	return
}
