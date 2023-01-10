package service

import (
	"api/internal/library/liberr"
	"api/internal/model"
	srv "api/internal/service"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

type gft struct {
	options *model.TokenOptions
	gT      srv.IGfToken
	lock    *sync.Mutex
}

var gftService = &gft{
	options: nil,
	gT:      nil,
	lock:    &sync.Mutex{},
}

func GfToken() srv.IGfToken {
	//if gftService.gT == nil {
	gftService.lock.Lock()
	defer gftService.lock.Unlock()
	//if gftService.gT == nil {
	ctx := gctx.New()
	err := g.Cfg("token").MustGet(ctx, "gfToken").Struct(&gftService.options)
	liberr.IsNil(ctx, err)
	gftService.gT = srv.GfToken(gftService.options)
	//}
	//}
	return gftService.gT
}
