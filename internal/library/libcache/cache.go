package libcache

import (
	"context"
	"reflect"
	"sync"
	"time"

	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gogf/gf/v2/encoding/gjson"

	"github.com/gogf/gf/v2/container/gvar"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/os/gcache"
)

type ICache interface {
	Get(ctx context.Context, key string) *gvar.Var
	Set(ctx context.Context, key string, value interface{}, duration time.Duration, tag ...string)
	Remove(ctx context.Context, key string) interface{}
	Removes(ctx context.Context, keys []string)
	RemoveByTag(ctx context.Context, tag string)
	RemoveByTags(ctx context.Context, tag []string)
	SetIfNotExist(ctx context.Context, key string, value interface{}, duration time.Duration, tag string) bool
	GetOrSet(ctx context.Context, key string, value interface{}, duration time.Duration, tag string) interface{}
	GetOrSetFunc(ctx context.Context, key string, f gcache.Func, duration time.Duration, tag string) interface{}
	GetOrSetFuncLock(ctx context.Context, key string, f gcache.Func, duration time.Duration, tag string) interface{}
	Contains(ctx context.Context, key string) bool
	Data(ctx context.Context) map[interface{}]interface{}
	Keys(ctx context.Context) []interface{}
	KeyStrings(ctx context.Context) []string
	Values(ctx context.Context) []interface{}
	Size(ctx context.Context) int
}

type GfCache struct {
	CachePrefix string        //缓存前缀
	cache       *gcache.Cache // 缓存
	tagSetMux   sync.Mutex    // 锁
}

func New(prefix string) *GfCache {
	return &GfCache{
		CachePrefix: prefix,
		cache:       gcache.New(),
	}
}

// NewRedis 创建redis缓存
func NewRedis(prefix string) *GfCache {
	cache := New(prefix)
	cache.cache.SetAdapter(gcache.NewAdapterRedis(g.Redis()))
	return cache
}

// Get 获取缓存
func (c *GfCache) Get(ctx context.Context, key string) *gvar.Var {
	v, err := c.cache.Get(ctx, c.CachePrefix+key)
	if err != nil {
		g.Log("exception").Error(ctx, err)
	}
	return v
}

// Set 设置缓存
func (c *GfCache) Set(ctx context.Context, key string, value interface{}, duration time.Duration, tag ...string) {
	// 开启锁
	c.tagSetMux.Lock()
	if len(tag) > 0 {
		c.cacheTagKey(ctx, key, tag[0])
	}

	if err := c.cache.Set(ctx, c.CachePrefix+key, value, duration); err != nil {
		g.Log("exception").Error(ctx, err)
	}
	// 释放锁
	c.tagSetMux.Unlock()
}

// Remove 删除缓存
func (c *GfCache) Remove(ctx context.Context, key string) interface{} {
	v, _ := c.cache.Remove(ctx, c.CachePrefix+key)
	return v
}

// Removes 批量删除缓存
func (c *GfCache) Removes(ctx context.Context, keys []string) {
	keyWithPrefix := make([]interface{}, len(keys))
	for k, v := range keys {
		keyWithPrefix[k] = c.CachePrefix + v
	}
	_, _ = c.cache.Remove(ctx, keyWithPrefix...)
}

// RemoveByTag 根据tag删除缓存
func (c *GfCache) RemoveByTag(ctx context.Context, tag string) {
	c.tagSetMux.Lock()
	tagKey := c.setTagKey(tag)
	//删除tagKey 对应的 key和值
	keys := c.Get(ctx, tagKey)
	if !keys.IsNil() {
		// 如果是字符串
		if kStr, ok := keys.Val().(string); ok {
			js, err := gjson.DecodeToJson(kStr)
			if err != nil {
				g.Log("exception").Error(ctx, err)
				return
			}
			ks := gconv.SliceStr(js.Interface())
			c.Removes(ctx, ks)
		} else {
			// 如果是数组
			ks := gconv.SliceStr(keys.Val())
			c.Removes(ctx, ks)
		}
	}
	// 删除tagKey
	c.Remove(ctx, tagKey)
	c.tagSetMux.Unlock()
}

// RemoveByTags 根据tags批量删除缓存
func (c *GfCache) RemoveByTags(ctx context.Context, tag []string) {
	for _, v := range tag {
		c.RemoveByTag(ctx, v)
	}
}

// SetIfNotExist 设置缓存，如果不存在则设置缓存 duration 为0时永久缓存
func (c *GfCache) SetIfNotExist(ctx context.Context, key string, value interface{}, duration time.Duration, tag string) bool {
	// 开启锁
	c.tagSetMux.Lock()
	// 释放锁
	defer c.tagSetMux.Unlock()
	c.cacheTagKey(ctx, key, tag)
	v, _ := c.cache.SetIfNotExist(ctx, c.CachePrefix+key, value, duration)
	return v
}

// GetOrSet 获取缓存，缓存不存在则返回value并设置缓存
func (c *GfCache) GetOrSet(ctx context.Context, key string, value interface{}, duration time.Duration, tag string) interface{} {
	c.tagSetMux.Lock()
	defer c.tagSetMux.Unlock()
	c.cacheTagKey(ctx, key, tag)
	v, _ := c.cache.GetOrSet(ctx, c.CachePrefix+key, value, duration)
	return v
}

// GetOrSetFunc 获取缓存，缓存不存在则使用函数进行设置
func (c *GfCache) GetOrSetFunc(ctx context.Context, key string, f gcache.Func, duration time.Duration, tag string) interface{} {
	c.tagSetMux.Lock()
	defer c.tagSetMux.Unlock()
	c.cacheTagKey(ctx, key, tag)
	v, _ := c.cache.GetOrSetFunc(ctx, c.CachePrefix+key, f, duration)
	return v
}

func (c *GfCache) GetOrSetFuncLock(ctx context.Context, key string, f gcache.Func, duration time.Duration, tag string) interface{} {
	c.tagSetMux.Lock()
	defer c.tagSetMux.Unlock()
	c.cacheTagKey(ctx, key, tag)
	v, _ := c.cache.GetOrSetFuncLock(ctx, c.CachePrefix+key, f, duration)
	return v
}

func (c *GfCache) Contains(ctx context.Context, key string) bool {
	v, _ := c.cache.Contains(ctx, c.CachePrefix+key)
	return v
}

func (c *GfCache) Data(ctx context.Context) map[interface{}]interface{} {
	v, _ := c.cache.Data(ctx)
	return v
}

// Keys returns all keys in the cache as slice.
func (c *GfCache) Keys(ctx context.Context) []interface{} {
	v, _ := c.cache.Keys(ctx)
	return v
}

// KeyStrings returns all keys in the cache as string slice.
func (c *GfCache) KeyStrings(ctx context.Context) []string {
	v, _ := c.cache.KeyStrings(ctx)
	return v
}

// Values returns all values in the cache as slice.
func (c *GfCache) Values(ctx context.Context) []interface{} {
	v, _ := c.cache.Values(ctx)
	return v
}

// Size returns the size of the cache.
func (c *GfCache) Size(ctx context.Context) int {
	v, _ := c.cache.Size(ctx)
	return v
}

// setTagKey 设置tagKey名称
func (c *GfCache) setTagKey(tag string) string {
	if tag != "" {
		tag = "tag_" + tag
	}
	return tag
}

// cacheTagKey 缓存tagKey
func (c *GfCache) cacheTagKey(ctx context.Context, key interface{}, tag string) {
	tagKey := c.CachePrefix + c.setTagKey(tag)
	if tagKey != "" {
		tagValue := []interface{}{key}
		value, _ := c.cache.Get(ctx, tagKey)
		if value != nil {
			var keyValue []interface{}
			// 若是字符串
			if kStr, ok := value.Val().(string); ok {
				js, err := gjson.DecodeToJson(kStr)
				if err != nil {
					g.Log("exception").Error(ctx, err)
					return
				}
				keyValue = gconv.SliceAny(js.Interface())
			} else {
				keyValue = gconv.SliceAny(value)
			}
			for _, v := range keyValue {
				// 比较两个值类型与值是否相等
				if !reflect.DeepEqual(key, v) {
					tagValue = append(tagValue, v)
				}
			}
		}
		_ = c.cache.Set(ctx, tagKey, tagValue, 0)
	}
}
