package cache

import (
	"time"
)

type CacheItem struct {
	value   string
	timeout time.Time
}

type Cache struct {
	cache map[string]CacheItem
}

func NewCache() Cache {
	return Cache{cache: map[string]CacheItem{}}
}

func (c Cache) Get(key string) (string, bool) {
	item, exists := c.cache[key]
	if item.timeout.UnixMilli() > time.Now().UnixMilli() {
		return item.value, false
	}
	return item.value, exists
}

func (c Cache) Put(key, value string) {
	cacheItem := CacheItem{value: value}
	c.cache[key] = cacheItem
}

func (c Cache) Keys() []string {
	keys := make([]string, 0, len(c.cache))
	for k := range c.cache {
		keys = append(keys, k)
	}
	return keys
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	c.cache[key] = CacheItem{value: value, timeout: deadline}
}
