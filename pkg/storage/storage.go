package storage

import (
	"context"
	"fmt"
	"sync"
)

type Storage struct {
	db    Queries
	cache sync.Map
}

func NewStorage(conn Queries) *Storage {
	return &Storage{db: conn}
}

func (s *Storage) Create(ctx context.Context, order *Order) error {
	q := "INSERT INTO orders (orderuuid, data) VALUES ($1, $2) RETURNING id"
	if err := s.db.QueryRow(ctx, q, order.OrderUuid, order.Data).Scan(&order.ID); err != nil {
		return fmt.Errorf("error adding order: %w", err)
	}
	return nil
}
