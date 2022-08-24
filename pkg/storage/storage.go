package storage

import (
	"context"
	"fmt"
	"log"
	"sync"
)

type Storage struct {
	db Queries
	m  sync.Map
}

var ord Order

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

func (s *Storage) ByUUID(ctx context.Context, orderuuid string) (*Order, error) {
	q := `SELECT id, data FROM orders WHERE orderuuid = $1`
	var ord Order
	if err := s.db.QueryRow(ctx, q, orderuuid).Scan(&ord.ID, &ord.Data); err != nil {
		return nil, fmt.Errorf("error get uuid: %w", err)
	}
	return &ord, nil
}

func (s *Storage) Load(ctx context.Context) error {
	rows, err := s.db.Query(ctx, `SELECT * FROM orders ORDER BY id`)
	if err != nil {
		return fmt.Errorf("error load orders from db: %w", err)
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var order Order
		err := rows.Scan(&order.ID, &order.OrderUuid, &order.Data)
		if err != nil {
			log.Fatal(err)
		}
		orders = append(orders, order)
		s.m.Store(order.OrderUuid, order.Data)
	}

	return nil
}
