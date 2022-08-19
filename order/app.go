package order

import (
	"L0/pkg/storage"
	"context"
	"encoding/json"
	"github.com/nats-io/stan.go"
	"log"
)

type App struct {
	st *storage.Storage
}

func NewApp(st *storage.Storage) *App {
	return &App{st: st}
}

func (a *App) Consumer(s *stan.Msg) {
	var msg natsMessage

	order := &storage.Order{}
	order.Data = s.Data

	if err := json.Unmarshal(s.Data, &msg); err != nil {
		log.Println(err)
		return
	}

	order := &storage.Order{}
	if err := a.st.AddOrder(context.Background(), order); err != nil {
		log.Println(err)
	}
}
