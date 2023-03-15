package localcache

type Cache interface {
	// Get the data from the cache with k(key)
	Get(k string) any
	// Set the v(data) to the cache with k(key)
	Set(k string, v any)
}
