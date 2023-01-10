package cache

import (
	"context"
	"sync"
	"time"

	"github.com/gogf/gf/v2/container/gvar"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/os/gcache"
)

type GfCache struct {
	Prefix    string        // 前缀
	cache     *gcache.Cache // 缓存实例
	tagSetMux sync.Mutex    // 锁
}

var (
	cache *gcache.Cache
	once  sync.Once
)

// New 默认缓存
func New(prefix string) *GfCache {
	once.Do(func() {
		cache = gcache.New()
	})
	return &GfCache{
		Prefix: prefix,
		cache:  cache,
	}
}

// NewRedis 改为redis缓存
func NewRedis(prefix string) *GfCache {
	cache := New(prefix)
	cache.cache.SetAdapter(gcache.NewAdapterRedis(g.Redis()))
	return cache
}

// 不同的缓存，按规范实现即可

// Get 获取缓存
func (c *GfCache) Get(ctx context.Context, key string) (*gvar.Var, error) {
	return c.cache.Get(ctx, c.Prefix+key)
}

// Set 设置缓存
func (c *GfCache) Set(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	// 开启锁
	c.tagSetMux.Lock()
	err := c.cache.Set(ctx, c.Prefix+key, value, duration)
	// 释放锁
	c.tagSetMux.Unlock()
	return err
}

// Remove 删除缓存
func (c *GfCache) Remove(ctx context.Context, key string) error {
	_, err := c.cache.Remove(ctx, c.Prefix+key)
	return err
}

// Has 缓存是否存在
func (c *GfCache) Has(ctx context.Context, key string) (bool, error) {
	return c.cache.Contains(ctx, c.Prefix+key)
}
