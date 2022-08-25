package order

import (
	"fmt"
	"sync"
)

type inMemoryCache struct {
	m sync.Map
}

func (c *inMemoryCache) set(key string, value []byte) {
	c.m.Store(key, value)
}

func (c *inMemoryCache) get(key string) ([]byte, error) {
	val, ok := c.m.Load(key)
	if !ok {
		return nil, fmt.Errorf("key %s doesn't stored in cache", key)
	}
	return val.([]byte), nil
}
