package storage

import "sync"

type Storage struct {
	db    Queries
	cache sync.Map
}

func NewStorage(db Queries) *Storage {
	return &Storage{db: db}
}
