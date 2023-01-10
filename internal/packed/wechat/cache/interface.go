package cache

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/container/gvar"
)

type ICache interface {
	// Get 读取缓存信息
	Get(ctx context.Context, key string) (*gvar.Var, error)
	// Set 设置缓存信息
	Set(ctx context.Context, key string, value interface{}, duration time.Duration) error
	// Remove 删除缓存信息
	Remove(ctx context.Context, key string) error
	// Has 判断缓存是否存在
	Has(ctx context.Context, key string) (bool, error)
}
