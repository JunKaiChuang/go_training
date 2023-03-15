// Store the data in the local memory
package localcache

import "time"

const (
	// TTL is the default time to live(seconds) for a cache entry
	TTL = 30
)

var (
	// function variable of time.Now
	timeNow = time.Now
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
		cacheMap: make(map[string]cacheData),
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

func expireAt() time.Time {
	return timeNow().Add(TTL * time.Second)
}
