package audit_process

import (
	"api/api/v1/admin"
	"api/internal/dao"
	"api/internal/library/liberr"
	"api/internal/model"
	"api/internal/model/do"
	"api/internal/model/entity"
	"api/internal/service"
	"context"

	"github.com/gogf/gf/v2/database/gdb"

	"github.com/gogf/gf/v2/util/gutil"

	"github.com/gogf/gf/v2/container/gmap"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/frame/g"
)

type sAuditProcess struct {
}

func init() {
	service.RegisterAuditProcess(New())
}

func New() *sAuditProcess {
	return &sAuditProcess{}
}

// Create 创建审核类型
func (s *sAuditProcess) Create(ctx context.Context, req *admin.AuditProcessCreateReq) (err error) {
	// 新增审核类型
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.ConfigAuditProcess.Ctx(ctx).Data(req).OmitEmpty().Insert()
		liberr.IsNil(ctx, err, "添加审核类型失败")
	})
	return err
}

// Update 更新审核类型
func (s *sAuditProcess) Update(ctx context.Context, req *admin.AuditProcessUpdateReq) (err error) {
	// 开启事务
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			// 删除现有流程数据
			if _, err = dao.ConfigAuditProcess.Ctx(ctx).TX(tx).Where(dao.ConfigAuditProcess.Columns().ServiceType, req.ServiceType).Delete(); err != nil {
				liberr.IsNil(ctx, err, "删除审核流程失败")
			}
			insertData := make([]do.ConfigAuditProcess, 0)
			for k, v := range req.AuditProcessArr {
				if len(v.UserId) > 0 {
					for _, u := range v.UserId {
						insertData = append(insertData, do.ConfigAuditProcess{
							ServiceType:       req.ServiceType,
							AuditDepartmentId: v.DepartmentId,
							AuditUserId:       u,
							Procedure:         k + 1,
							ServiceName:       req.ServiceName,
						})
					}
				} else {
					insertData = append(insertData, do.ConfigAuditProcess{
						ServiceType:       req.ServiceType,
						AuditDepartmentId: v.DepartmentId,
						AuditUserId:       0,
						Procedure:         k + 1,
						ServiceName:       req.ServiceName,
					})
				}
			}
			// 添加审核流程
			if _, err = dao.ConfigAuditProcess.Ctx(ctx).TX(tx).Data(insertData).Insert(); err != nil {
				liberr.IsNil(ctx, err, "添加审核流程失败")
			}
		})
		return err
	})
	return err
}

// Delete 删除审核类型
func (s *sAuditProcess) Delete(ctx context.Context, req *admin.AuditProcessDeleteReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.ConfigAuditProcess.Ctx(ctx).Where(do.ConfigAuditProcess{
			ServiceType: req.ServiceType,
			Procedure:   req.Process,
		}).Delete()
		liberr.IsNil(ctx, err, "删除审核类型失败")
	})
	return err
}

// Detail 获取审核类型详情
func (s *sAuditProcess) Detail(ctx context.Context, req *admin.AuditProcessDetailReq) (res *admin.AuditProcessDetailRes, err error) {
	res = new(admin.AuditProcessDetailRes)
	var list []*entity.ConfigAuditProcess
	treeMap := gmap.NewTreeMap(gutil.ComparatorInt, true)
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.ConfigAuditProcess.Ctx(ctx).Where(dao.ConfigAuditProcess.Columns().ServiceType, req.ServiceType).Order("procedure asc").Scan(&list)
		liberr.IsNil(ctx, err, "获取审核类型详情失败")
		// 声明一个 TreeMap 类型
		var item = make(map[int]*model.AuditProcessDetailItem, 0)
		for _, v := range list {
			if item[v.Procedure] == nil {
				if gconv.Int(v.AuditUserId) > 0 {
					item[v.Procedure] = &model.AuditProcessDetailItem{
						Procedure:    v.Procedure,
						DepartmentId: v.AuditDepartmentId,
						UserId:       []int{gconv.Int(v.AuditUserId)},
					}
				} else {
					item[v.Procedure] = &model.AuditProcessDetailItem{
						Procedure:    v.Procedure,
						DepartmentId: v.AuditDepartmentId,
						UserId:       []int{},
					}
				}
			} else {
				item[v.Procedure].UserId = append(item[v.Procedure].UserId, gconv.Int(v.AuditUserId))
			}
		}
		// 进行有序遍历
		for k, v := range item {
			treeMap.Set(k, v)
		}
		// 转换
		lists := treeMap.Values()
		_ = gconv.Struct(lists, &res.Lists)
	})
	return res, err
}

// List 获取审核类型列表
func (s *sAuditProcess) List(ctx context.Context, req *admin.AuditProcessListReq) (res *admin.AuditProcessListRes, err error) {
	res = new(admin.AuditProcessListRes)

	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.ConfigAuditProcess.Ctx(ctx).Group(dao.ConfigAuditProcess.Columns().ServiceType).Scan(&res.List)
		liberr.IsNil(ctx, err, "获取审核类型列表失败")
	})
	return res, err
}
