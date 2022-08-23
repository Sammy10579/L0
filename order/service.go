package order

import (
	"L0/pkg/storage"
	"context"
	"encoding/json"
	"github.com/nats-io/stan.go"
	"log"
)

type Service struct {
	st *storage.Storage
}

func NewService(st *storage.Storage) *Service {
	return &Service{st: st}
}

type natsMessage struct {
	OrderUuid string `json:"order_uid"`
}

func (s *Service) Load(m *stan.Msg) {
	var msg natsMessage

	if err := json.Unmarshal(m.Data, &msg); err != nil {
		log.Println(err)
		return
	}

	order := &storage.Order{}
	order.Data = m.Data
	order.OrderUuid = msg.OrderUuid

	if err := s.st.Safe(context.Background(), order); err != nil {
		log.Println(err)
	}
}
