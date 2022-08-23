package order

import (
	"L0/pkg/storage"
	"context"
)

type Service struct {
	st *storage.Storage
}

func NewService(st *storage.Storage) *Service {
	return &Service{st: st}
}

func (s *Service) Save(ctx context.Context, order *storage.Order) error {
	return s.st.Create(ctx, order)
}

func (s *Service) ByUUID(ctx context.Context, orderuuid string) (*storage.Order, error) {
	return s.st.ByUUID(ctx, orderuuid)
}
