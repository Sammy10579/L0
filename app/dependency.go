package app

import (
	"context"

	"L0/pkg/order"
)

type OrderService interface {
	PayloadByUid(ctx context.Context, uid string) ([]byte, error)
	Save(ctx context.Context, order order.Order) error
}
