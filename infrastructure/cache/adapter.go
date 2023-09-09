package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var caching = cache.New(24*time.Hour, 24*time.Hour)

func GetCache(cacheKey string) interface{} {
	if val, ok := caching.Get(cacheKey); ok {
		return val
	}

	return nil
}

func SetCache(cacheKey string, value interface{}, expTime time.Duration) {
	if expTime == 0 {
		expTime = cache.DefaultExpiration
	}

	caching.Set(cacheKey, value, expTime)
}

func RemoveFromCache(cacheKey string) {
	caching.Delete(cacheKey)
}
