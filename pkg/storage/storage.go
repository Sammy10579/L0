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
	q := `INSERT INTO orders (orderuuid, data) VALUES ($1, $2) RETURNING id`
	if err := s.db.QueryRow(ctx, q, order.OrderUuid, order.Data).Scan(&order.ID); err != nil {
		return fmt.Errorf("error adding order: %w", err)
	}
	return nil
}

func (s *Storage) ByUUID(ctx context.Context, orderuuid string) (Order, error) {
	q := `SELECT id, data FROM orders WHERE orderuuid = $1`
	var ord Order
	if err := s.db.QueryRow(ctx, q, orderuuid).Scan(&ord.ID, ord.Data); err != nil {
		return Order{}, fmt.Errorf("error getting uuid: %w", err)
	}
	return ord, nil
}

func (s *Storage) LoadCache(ctx context.Context, order *Order) error {
	q := `SELECT * FROM orders ORDER BY id`
	var orders []Order
	if err := s.db.QueryRow(ctx, q, order.OrderUuid, order.Data).Scan(&order.ID, &order.OrderUuid, &order.Data); err != nil {
		var order Order
		orders = append(orders, order)
		s.cache.Store(order.OrderUuid, order.Data)
	}
	return nil
}
