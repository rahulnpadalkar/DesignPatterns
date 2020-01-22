package cache

type CacheReplacemnetAlgo interface {
	Replace(c *Cache)
}