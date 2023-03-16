// Store the data in the local memory
package localcache

import "time"

const (
	// cacheTTL is the default time to live(seconds) for a cache entry
	cacheTTL = 30
)

var (
	// function variable of time.Now
	timeNow = time.Now
	// function variable of make cache map
	initCacheMap = func() map[string]cacheData {
		return make(map[string]cacheData)
	}
)

type cacheData struct {
	Data     any
	ExpireAt time.Time
}

type cache struct {
	cacheMap map[string]cacheData
}

// Returns a new Cache instance
func New() Cache {
	return &cache{
		cacheMap: initCacheMap(),
	}
}

// Get the data from the cache with k(key)
func (c cache) Get(k string) any {
	now := timeNow()
	val, ok := c.cacheMap[k]

	if !ok {
		return nil
	}

	if now.After(val.ExpireAt) {
		delete(c.cacheMap, k)
		return nil
	}

	return val.Data
}

// Set the v(data) to the cache with k(key)
func (c cache) Set(k string, v any) {
	c.cacheMap[k] = cacheData{
		Data:     v,
		ExpireAt: expireAt(),
	}
}

// Get the cache map
func (c cache) CacheMap() map[string]cacheData {
	return c.cacheMap
}

func expireAt() time.Time {
	return timeNow().Add(cacheTTL * time.Second)
}
