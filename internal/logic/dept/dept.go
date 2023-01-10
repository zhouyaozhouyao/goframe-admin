package dept

import (
	"api/internal/library/liberr"
	"api/internal/model"
	"api/internal/model/do"
	"api/internal/model/entity"
	"api/internal/modules/admin/consts"
	adminService "api/internal/modules/admin/service"
	"context"

	"github.com/gogf/gf/v2/text/gstr"

	"github.com/gogf/gf/v2/util/gconv"

	"api/api/v1/admin"
	"api/internal/dao"
	"api/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

type sDept struct {
}

func init() {
	service.RegisterDept(New())
}

func New() *sDept {
	return &sDept{}
}

// Create 添加组织架构
func (s *sDept) Create(ctx context.Context, in *admin.DeptCreateReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.Dept.Ctx(ctx).Data(do.Dept{
			ParentId:  in.ParentId,
			DeptName:  in.DeptName,
			OrderNum:  in.OrderNum,
			Leader:    in.Leader,
			Phone:     in.Phone,
			Email:     in.Email,
			Status:    in.Status,
			CreatedBy: adminService.Context().GetUserId(ctx),
			UpdatedBy: adminService.Context().GetUserId(ctx),
		}).OmitEmpty().Insert()
		// 删除缓存
		liberr.IsNil(ctx, err, "添加组织架构失败")
		adminService.Cache().Remove(ctx, consts.CacheSysDept)
	})
	return
}

func (s *sDept) Update(ctx context.Context, in *admin.DeptUpdateReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.Dept.Ctx(ctx).WherePri(in.DeptId).Data(do.Dept{
			ParentId:  in.ParentId,
			DeptName:  in.DeptName,
			OrderNum:  in.OrderNum,
			Leader:    in.Leader,
			Phone:     in.Phone,
			Email:     in.Email,
			Status:    in.Status,
			UpdatedBy: adminService.Context().GetUserId(ctx),
		}).OmitEmpty().Update()
		liberr.IsNil(ctx, err, "更新组织架构失败")
		adminService.Cache().Remove(ctx, consts.CacheSysDept)
	})
	return
}

func (s *sDept) Delete(ctx context.Context, in *admin.DeptDeleteReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.Dept.Ctx(ctx).WherePri(in.DeptId).Delete()
		liberr.IsNil(ctx, err, "删除组织架构失败")
		adminService.Cache().Remove(ctx, consts.CacheSysDept)
	})
	return
}

func (s *sDept) Info(ctx context.Context, in *admin.DeptInfoReq) (out *entity.Dept, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.Dept.Ctx(ctx).WherePri(in.DeptId).Scan(&out)
	})
	return
}

// GetList 获取组织架构列表
func (s *sDept) GetList(ctx context.Context, req *admin.DeptListReq) (list []*entity.Dept, err error) {
	// 查询缓存没有则进行设置缓存
	list, err = s.GetListCache(ctx)
	if err != nil {
		return
	}
	rList := make([]*entity.Dept, 0, len(list))
	if req.DeptName != "" {
		for _, v := range list {
			if v.DeptName != "" && !gstr.Contains(v.DeptName, req.DeptName) {
				continue
			}
			rList = append(rList, v)
		}
		list = rList
	}
	return
}

// GetListCache 设置与获取缓存
func (s *sDept) GetListCache(ctx context.Context) (list []*entity.Dept, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		cache := adminService.Cache()
		// 从缓存中获取
		iList := cache.GetOrSetFuncLock(ctx, consts.CacheSysDept, func(ct context.Context) (value interface{}, err error) {
			err = dao.Dept.Ctx(ctx).Scan(&list)
			liberr.IsNil(ctx, err, "查询组织架构失败")
			value = list
			return
		}, 0, consts.CacheSysAuthTag)
		if iList != nil {
			err = gconv.Struct(iList, &list)
			liberr.IsNil(ctx, err)
		}
	})
	return
}

func (s *sDept) Tree(pid int64, list []*entity.Dept) (deptTree []*model.DeptTreeOutput) {
	deptTree = make([]*model.DeptTreeOutput, 0, len(list))
	for _, v := range list {
		if v.ParentId == pid {
			t := &model.DeptTreeOutput{
				Dept: v,
			}
			child := s.Tree(v.DeptId, list)
			if len(child) > 0 {
				t.Children = child
			}
			deptTree = append(deptTree, t)
		}
	}
	return
}
