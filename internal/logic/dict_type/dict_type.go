package dict_type

import (
	"api/api/v1/admin"
	"api/internal/dao"
	"api/internal/library/liberr"
	"api/internal/service"
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/frame/g"
)

type sDictType struct {
}

func init() {
	service.RegisterDictType(New())
}

func New() *sDictType {
	return &sDictType{}
}

func (s *sDictType) Create(ctx context.Context, in *admin.DictTypeCreateReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		// 添加类型
		_, err = dao.DictType.Ctx(ctx).Data(in).OmitEmpty().Insert()
		liberr.IsNil(ctx, err, "添加字典类型失败")
	})
	return err
}

func (s *sDictType) Update(ctx context.Context, in *admin.DictTypeUpdateReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		// 更新类型
		_, err = dao.DictType.Ctx(ctx).WherePri(in.Id).Data(in).OmitEmpty().Update()
		liberr.IsNil(ctx, err, "更新字典类型失败")
	})
	return err
}

func (s *sDictType) Delete(ctx context.Context, in *admin.DictTypeDeleteReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		// 查询当前 ids 对应的类型值
		result, _ := dao.DictType.Ctx(ctx).WhereIn(dao.DictType.Columns().DictId, in.Ids).Fields("dict_type").All()
		// 删除类型
		_, err = dao.DictType.Ctx(ctx).WhereIn(dao.DictType.Columns().DictId, in.Ids).Delete()
		// 删除类型匹配的数据
		_, err = dao.DictData.Ctx(ctx).WhereIn(dao.DictData.Columns().DictType, result.Array("dict_type")).Delete()
	})
	return
}

func (s *sDictType) List(ctx context.Context, in *admin.DictTypeListReq) (res *admin.DictTypeListRes, err error) {
	res = new(admin.DictTypeListRes)
	err = g.Try(ctx, func(ctx context.Context) {
		// 初始化条件与页码信息
		in.ConditionOrPaginate()
		m := dao.DictType.Ctx(ctx)
		if in.DictName != "" {
			m = m.WhereLike(dao.DictType.Columns().DictName, in.DictName)
		}
		if in.Status != "" {
			m = m.Where(dao.DictType.Columns().Status, gconv.Uint(in.Status))
		}
		if in.DictType != "" {
			m = m.Where(dao.DictType.Columns().DictType, in.DictType)
		}
		res.Total, err = m.Count()
		liberr.IsNil(ctx, err, "获取字典类型列表失败")
		res.CurrentPage = in.PageNum
		err = m.Page(in.PageNum, in.PageSize).Order(in.OrderBy).Scan(&res.List)
		liberr.IsNil(ctx, err, "获取字典类型列表失败")
	})
	return
}
