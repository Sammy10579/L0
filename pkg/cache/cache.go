package cache

import "sync"

func NewCache(orders sync.Map) Cache {
	return Cache{
		Cache: orders,
	}
}
