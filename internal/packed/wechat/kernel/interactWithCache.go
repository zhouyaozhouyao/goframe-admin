package kernel

import (
	"api/internal/packed/wechat/cache"
)

type InteractWithCache struct {
	Cache  CacheInterface
	prefix string
}

type CacheInterface cache.ICache

func NewInteractsWithCache(client CacheInterface, prefix string) *InteractWithCache {
	interact := &InteractWithCache{
		Cache:  client,
		prefix: prefix,
	}
	if client == nil {
		interact.Cache = cache.New(prefix)
	}

	return interact
}

func NewRedisClient() CacheInterface {
	return cache.NewRedis("powerwechat.access_token.")
}

func (interactCache *InteractWithCache) GetCache() CacheInterface {
	if interactCache.Cache != nil {
		return interactCache.Cache
	}

	// create default cache
	interactCache.Cache = cache.New(interactCache.prefix)

	return interactCache.Cache
}

func (interactCache *InteractWithCache) SetCache(cache CacheInterface) *InteractWithCache {

	interactCache.Cache = cache

	return interactCache
}
