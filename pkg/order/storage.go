package order

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

type storage struct {
	db *pgx.Conn
}

func (s *storage) create(ctx context.Context, o Order) error {
	q := `INSERT INTO orders (uid, payload) VALUES ($1, $2)`
	if _, err := s.db.Exec(ctx, q, o.Uid, o.Payload); err != nil {
		return fmt.Errorf("error adding order: %w", err)
	}
	return nil
}

func (s *storage) all(ctx context.Context) ([]Order, error) {
	rows, err := s.db.Query(ctx, `SELECT * FROM orders`)
	if err != nil {
		return nil, fmt.Errorf("error load orders from db: %w", err)
	}
	defer rows.Close()

	var returnedOrders []Order
	for rows.Next() {
		var o Order
		err := rows.Scan(&o.ID, &o.Uid, &o.Payload)
		if err != nil {
			return nil, err
		}
		returnedOrders = append(returnedOrders, o)
	}

	return returnedOrders, nil
}
