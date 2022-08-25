package order

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
)

type Service struct {
	storage *storage
	cache   *inMemoryCache
}

func NewService(conn *pgx.Conn) *Service {
	cache := &inMemoryCache{}
	storage := &storage{db: conn}

	orders, err := storage.all(context.Background())
	if err != nil {
		log.Fatal(fmt.Errorf("cannot warm up cache: %w", err))
	}
	for i := range orders {
		cache.set(orders[i].Uid, orders[i].Payload)
	}

	return &Service{
		storage: storage,
		cache:   cache,
	}
}

func (s *Service) Save(ctx context.Context, order Order) error {
	if err := s.storage.create(ctx, order); err != nil {
		return err
	}
	s.cache.set(order.Uid, order.Payload)

	return nil
}

func (s *Service) PayloadByUid(_ context.Context, uid string) ([]byte, error) {
	// we are get payload from cache, so we don't use context
	return s.cache.get(uid)
}
