package log

import (
	"api/internal/dao"
	"api/internal/library/libUtils"
	"api/internal/model"
	"api/internal/model/do"
	"api/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/os/grpool"

	"github.com/mssola/user_agent"
)

type sLoginLog struct {
	Pool *grpool.Pool
}

var (
	loginLogService = sLoginLog{
		Pool: grpool.New(100),
	}
)

func init() {
	service.RegisterLoginLog(New())
}

func New() *sLoginLog {
	return &loginLogService
}

// Invoke 写入登录日志
func (s *sLoginLog) Invoke(ctx context.Context, data *model.LoginLogInput) {
	_ = s.Pool.Add(ctx, func(ctx context.Context) {
		ua := user_agent.New(data.UserAgent)
		browser, _ := ua.Browser()
		logData := &do.LoginLog{
			LoginName:     data.Username,
			Ipaddr:        data.Ip,
			LoginLocation: libUtils.GetCityByIp(data.Ip),
			Browser:       browser,
			Os:            ua.OS(),
			Status:        data.Status,
			Msg:           data.Msg,
			LoginTime:     gtime.Now(),
			Module:        data.Module,
		}
		_, err := dao.LoginLog.Ctx(ctx).Insert(logData)
		if err != nil {
			g.Log().Error(ctx, err)
		}
	})
}
