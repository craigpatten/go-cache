package lru

import . "go-cache"

type cacheEntry struct {
	object CacheObject
}

func newCacheEntry() *cacheEntry {
	return &cacheEntry{}
}