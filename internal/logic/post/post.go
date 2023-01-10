package post

import (
	"api/api/v1/admin"
	"api/internal/dao"
	"api/internal/library/liberr"
	"api/internal/model/entity"
	"api/internal/service"
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/frame/g"
)

type sPost struct {
}

func init() {
	service.RegisterPost(New())
}

func New() *sPost {
	return &sPost{}
}

func (s *sPost) List(ctx context.Context, in *admin.PostListReq) (res *admin.PostListRes, err error) {
	res = new(admin.PostListRes)
	err = g.Try(ctx, func(ctx context.Context) {
		// 初始化条件与页码信息
		in.ConditionOrPaginate()
		m := dao.Post.Ctx(ctx)
		if in.PostName != "" {
			m = m.WhereLike(dao.Post.Columns().PostName, in.PostName)
		}
		if in.Status != "" {
			m = m.Where(dao.Post.Columns().Status, gconv.Uint(in.Status))
		}
		res.Total, err = m.Count()
		liberr.IsNil(ctx, err, "获取岗位数据失败")
		res.CurrentPage = in.PageNum
		err = m.Page(in.PageNum, in.PageSize).Order(in.OrderBy).Scan(&res.List)
		liberr.IsNil(ctx, err, "获取岗位数据失败")
	})
	return
}

func (s *sPost) Create(ctx context.Context, in *admin.PostCreateReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.Post.Ctx(ctx).Data(in).OmitEmpty().Save()
		liberr.IsNil(ctx, err, "添加岗位失败")
	})
	return
}

func (s *sPost) Update(ctx context.Context, in *admin.PostUpdateReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.Post.Ctx(ctx).Data(in).OmitEmpty().Where(dao.Post.Columns().PostId, in.PostId).Update()
		liberr.IsNil(ctx, err, "更新岗位失败")
	})
	return
}

func (s *sPost) Delete(ctx context.Context, in *admin.PostDeleteReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.Post.Ctx(ctx).WhereIn(dao.Post.Columns().PostId, in.Ids).Delete()
		liberr.IsNil(ctx, err, "删除岗位失败")
	})
	return
}

// GetUsedPost 获取正常的岗位
func (s *sPost) GetUsedPost(ctx context.Context) (list []*entity.Post, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.Post.Ctx(ctx).Where(dao.Post.Columns().Status, 1).Order(dao.Post.Columns().PostSort+" asc", dao.Post.Columns().PostId+" id").Scan(&list)
	})
	return
}
