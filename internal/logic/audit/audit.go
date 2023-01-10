package audit

import (
	"api/api/v1/admin"
	"api/internal/dao"
	"api/internal/library/liberr"
	"api/internal/model"
	"api/internal/model/do"
	"api/internal/model/entity"
	"api/internal/modules/admin/consts"
	adminService "api/internal/modules/admin/service"
	"api/internal/service"
	"context"

	"github.com/gogf/gf/v2/database/gdb"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/frame/g"
)

type sAudit struct {
}

func init() {
	service.RegisterAudit(New())
}

func New() *sAudit {
	return &sAudit{}
}

func (s *sAudit) Create(ctx context.Context, in *model.AuditCreateInput) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		exists := s.auditExists(ctx, in.ServiceId)
		if exists {
			liberr.IsNil(ctx, err, "当前存在待审核数据，请先审核")
			return
		}

		// 创建审核
		_, err = dao.Audit.Ctx(ctx).Data(do.Audit{
			PlatformId:  in.PlatformId,
			ServiceName: in.ServiceName,
			ServiceType: in.ServiceType,
			ServiceId:   in.ServiceId,
			FinalStep:   in.FinalStep,
			CurrentStep: in.CurrentStep,
			IndexColumn: gconv.String(in.ServiceType) + "-" + gconv.String(in.CurrentStep),
			ApplyId:     in.ApplyId,
			AuditTime:   gtime.Now(),
			Remark:      in.Remark,
			Status:      1,
		}).OmitEmpty().Insert()
		liberr.IsNil(ctx, err, "创建审核失败")
	})
	return
}

// Update 审核状态改变
func (s *sAudit) Update(ctx context.Context, in *admin.AuditUpdateReq) (err error) {
	var audit = new(entity.Audit)
	var date = gtime.New()
	// 查询当前审核记录
	err = dao.Audit.Ctx(ctx).Where(do.Audit{Id: in.Id}).Scan(&audit)
	if err != nil {
		return err
	}

	// 查询下一步配置
	// 开启事务
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			// 添加审核记录
			if _, err := dao.AuditRecord.Ctx(ctx).TX(tx).Data(do.AuditRecord{
				AuditId:     in.Id,
				ServiceType: audit.ServiceType,
				ServiceId:   audit.ServiceId,
				AuditorId:   in.AuditorId,
				AuditRemark: in.AuditRemark,
				AuditStep:   audit.CurrentStep,
			}).OmitEmpty().Insert(); err != nil {
				liberr.IsNil(ctx, err, "审核记录添加失败")
				return
			}

			// 审核拒绝
			if in.Action == consts.AuditStatusRefuse {
				// 更新当前审核状态
				if _, err = dao.Audit.Ctx(ctx).TX(tx).Where(do.Audit{Id: in.Id}).Data(do.Audit{
					Status:      consts.AuditStatusRefuse,
					AuditRemark: in.AuditRemark,
					AuditTime:   date,
					AuditorId:   in.AuditorId,
				}).OmitEmpty().Update(); err != nil {
					return
				}
			} else {
				// 审核通过
				// 获取需要谁审核
				procedure := s.byPassAudit(ctx, audit, audit.CurrentStep+1)

				updateData := do.Audit{
					CurrentStep: procedure,
					IndexColumn: gconv.String(audit.ServiceType) + "-" + gconv.String(procedure),
					AuditorId:   adminService.Context().GetUserId(ctx),
					AuditRemark: in.AuditRemark,
					AuditTime:   date,
				}
				// 判断是否为最后一次审核
				if audit.FinalStep == audit.CurrentStep {
					updateData.CurrentStep = audit.CurrentStep
					updateData.IndexColumn = gconv.String(audit.ServiceType) + "-" + gconv.String(audit.CurrentStep)
					updateData.Status = consts.AuditStatusPass
				}

				// 开始更新
				if _, err := dao.Audit.Ctx(ctx).TX(tx).Where(do.Audit{Id: in.Id}).Data(updateData).OmitEmpty().Update(); err != nil {
					return
				}

				// 如果是审核完成需要发送推送
				if updateData.Status == consts.AuditStatusPass {
					// 发送推送
					adminService.Queue().Push(model.PushInput{
						ServiceId:   audit.ServiceId,
						ServiceType: audit.ServiceType,
						ApplyRemark: audit.Remark,
						AuditRemark: audit.AuditRemark,
						Result:      in.Action,
					})
				}
			}
		})
		return err
	})

	return err
}

func (s *sAudit) Delete(ctx context.Context, in []uint) (err error) {
	return
}

// List 我的审核列表
func (s *sAudit) List(ctx context.Context, req *admin.AuditListReq) (res *admin.AuditListRes, err error) {
	res = new(admin.AuditListRes)
	// 获取用户的审核配置列表
	auditConfig, _ := dao.ConfigAuditProcess.Ctx(ctx).Where(dao.ConfigAuditProcess.Columns().AuditUserId, req.AuditUserId).Fields("procedure, service_type").All()
	if auditConfig.IsEmpty() {
		return &admin.AuditListRes{}, nil
	}

	var indexColumn []string
	for _, v := range auditConfig {
		indexColumn = append(indexColumn, gconv.String(v.Map()["service_type"])+"-"+gconv.String(v.Map()["procedure"]))
	}

	total, _ := dao.Audit.Ctx(ctx).WhereIn(dao.Audit.Columns().IndexColumn, indexColumn).Count()
	if total < 1 {
		return &admin.AuditListRes{}, nil
	}

	err = dao.Audit.Ctx(ctx).WhereIn(dao.Audit.Columns().IndexColumn, indexColumn).Page(req.PageNum, req.PageSize).Scan(&res.List)
	return
}

func (s *sAudit) Detail(ctx context.Context, req *admin.AuditDetailReq) (res *admin.AuditDetailRes, err error) {
	res = new(admin.AuditDetailRes)
	err = dao.Audit.Ctx(ctx).Where(dao.Audit.Columns().Id, req.Id).Scan(&res)
	// 获取审核日志
	if res.CurrentStep >= 1 {
		auditRecord, _ := dao.AuditRecord.Ctx(ctx).Where(dao.AuditRecord.Columns().AuditId, req.Id).All()
		res.AuditRecord = auditRecord.List()
	}
	return res, err
}

func (s *sAudit) auditExists(ctx context.Context, serviceId string, id ...int64) (exists bool) {
	m := dao.Audit.Ctx(ctx)
	if len(id) > 0 {
		m = m.WhereNot(dao.Audit.Columns().Id, id[0])
	}
	count, err := m.Where(dao.Audit.Columns().ServiceId, serviceId).Where(dao.Audit.Columns().Status, consts.AuditStatusPending).Count()
	if err != nil {
		g.Log().Error(ctx, "查询记录是否存在", g.Map{"content": err})
		return true
	}

	return count > 0
}

// byPassAudit 检测审核配置
func (s *sAudit) byPassAudit(ctx context.Context, audit *entity.Audit, procedure int) int {
	auditConfExit, _ := dao.ConfigAuditProcess.Ctx(ctx).Where(do.ConfigAuditProcess{
		ServiceType: audit.ServiceType,
		Procedure:   procedure,
		Type:        1,
	}).One()

	if auditConfExit.IsEmpty() {
		return procedure
	}

	// 检测该审核是否离职
	userStatus, _ := dao.User.Ctx(ctx).Where(dao.User.Columns().Id, auditConfExit.Map()["audit_user_id"]).Value(dao.User.Columns().UserStatus)
	if userStatus.Int() == consts.UserStatusDisable && userStatus != nil {
		return s.byPassAudit(ctx, audit, gconv.Int(auditConfExit.Map()["procedure"])+1)
	}
	// 存在流程
	if !auditConfExit.IsEmpty() {
		return procedure
	}

	return procedure + 1
}
