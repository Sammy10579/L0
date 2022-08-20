package cache

import "sync"

type Cache struct {
	Cache sync.Map
}
