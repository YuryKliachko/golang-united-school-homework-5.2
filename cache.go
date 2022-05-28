package cache

import (
	"time"
)

type CacheItem struct {
	value   string
	expires bool
	timeout time.Time
}

func (ci CacheItem) isExpired() bool {
	if !ci.expires {
		return false
	} else {
		return ci.timeout.UnixMilli() < time.Now().UnixMilli()
	}
}

type Cache struct {
	cache map[string]CacheItem
}

func NewCache() Cache {
	return Cache{cache: map[string]CacheItem{}}
}

func (c Cache) Get(key string) (string, bool) {
	item, exists := c.cache[key]
	if !exists || item.isExpired() {
		return "", false
	}
	return item.value, true
}

func (c Cache) Put(key, value string) {
	cacheItem := CacheItem{value: value, expires: false}
	c.cache[key] = cacheItem
}

func (c Cache) Keys() []string {
	keys := make([]string, 0, len(c.cache))
	for k := range c.cache {
		item := c.cache[k]
		if item.isExpired() {
			continue
		}
		keys = append(keys, k)
	}
	return keys
}

func (c Cache) PutTill(key, value string, deadline time.Time) {
	c.cache[key] = CacheItem{value: value, expires: true, timeout: deadline}
}
