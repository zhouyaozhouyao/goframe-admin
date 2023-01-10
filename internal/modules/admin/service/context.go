package service

import (
	"api/internal/model"
	"api/internal/modules/admin/consts"
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

type IContext interface {
	Init(r *ghttp.Request, customCtx *model.Context)
	Get(ctx context.Context) *model.Context
	SetUser(ctx context.Context, ctxUser *model.ContextUser)
	GetLoginUser(ctx context.Context) *model.ContextUser
	GetUserId(ctx context.Context) uint64
	GetCode(ctx context.Context) string
}

type sContext struct {
}

// Context 上下文管理服务
var contextService = sContext{}

func Context() IContext {
	return &contextService
}

// Init 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func (s *sContext) Init(r *ghttp.Request, customCtx *model.Context) {

	r.SetCtxVar(consts.CtxKey, customCtx)
}

// Get 获得上下文变量，如果没有设置，那么返回nil
func (s *sContext) Get(ctx context.Context) *model.Context {
	value := ctx.Value(consts.CtxKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

// SetUser 将上下文信息设置到上下文请求中，注意是完整覆盖
func (s *sContext) SetUser(ctx context.Context, ctxUser *model.ContextUser) {
	s.Get(ctx).User = ctxUser
}

// GetLoginUser 获取当前登陆用户信息
func (s *sContext) GetLoginUser(ctx context.Context) *model.ContextUser {
	context := s.Get(ctx)
	if context == nil {
		return nil
	}
	return context.User
}

func (s *sContext) GetUserId(ctx context.Context) uint64 {
	context := s.Get(ctx)
	if context == nil {
		return 0
	}
	return context.User.Id
}

// GetCode 获取平台Code
func (s *sContext) GetCode(ctx context.Context) string {
	context := s.Get(ctx)
	if context == nil {
		return ""
	}
	return context.Code
}
