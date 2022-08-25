package app

import (
	"context"
	"fmt"
	"log"

	"L0/pkg/order"
	"github.com/jackc/pgx/v4"
	"github.com/nats-io/stan.go"
)

type Config struct {
	ListenedAddr string
	DBConnection string
	Nats         struct {
		URL       string
		ClusterID string
		ClientID  string
	}
}

type Application struct {
	config Config
}

func NewApplication(config Config) *Application {
	return &Application{config: config}
}

func (a *Application) Run() {
	ctx := context.Background()

	errCantInitialize := func(e error) {
		log.Fatal(fmt.Errorf("cant initialize application: %w", e))
	}

	stanConnection, err := stan.Connect(
		a.config.Nats.ClusterID,
		a.config.Nats.ClientID,
		stan.NatsURL(a.config.Nats.URL),
	)
	if err != nil {
		errCantInitialize(err)
	}
	defer stanConnection.Close()

	dbConn, err := pgx.Connect(ctx, a.config.DBConnection)
	if err != nil {
		errCantInitialize(err)
	}
	defer dbConn.Close(ctx)

	orderService := order.NewService(dbConn)
	runNatsListener(stanConnection, orderService)
	runHttpServer(a.config.ListenedAddr, orderService)
}
