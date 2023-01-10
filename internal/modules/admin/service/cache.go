package service

import (
	"api/internal/library/libcache"
	"api/internal/modules/admin/consts"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

type ICache interface {
	libcache.ICache
}

type sCache struct {
	*libcache.GfCache
	prefix string
}

var (
	c              = sCache{}
	cacheContainer *libcache.GfCache
	lock           = &sync.Mutex{}
)

// Cache 缓存初始化
func Cache() ICache {
	var (
		ch  = c
		ctx = gctx.New()
	)
	// 缓存前缀
	prefix := g.Cfg().MustGet(ctx, "admin.cache.prefix").String()
	// 缓存类型
	model := g.Cfg().MustGet(ctx, "admin.cache.model").String()
	if cacheContainer == nil {
		if cacheContainer == nil {
			lock.Lock()
			if model == consts.CacheModelRedis {
				cacheContainer = libcache.NewRedis(prefix)
			} else {
				cacheContainer = libcache.New(prefix)
			}
		}
		lock.Unlock()
	}
	ch.GfCache = cacheContainer
	return &ch
}
