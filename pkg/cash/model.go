package cache

import (
	"sync"
	"time"
)

type Order struct {
	ID          int64  `db:"id"`
	OrderUuid   string `db:"order_uid"`
	TrackNumber string `db:"track_number"`
}

type Cache struct {
	sync.RWMutex
	defaultExpiration time.Duration
	cleanupInterval   time.Duration
	Order             map[string]int
}
