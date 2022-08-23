package main

import (
	"L0/order"
	"L0/pkg/storage"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"github.com/nats-io/stan.go"
	"log"
	"net/http"
	"os"
)

const (
	stanClusterID = "test-cluster"
	clientID      = "order-consumer"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	if err = conn.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}

	defer conn.Close(context.Background())

	store := storage.NewStorage(conn)
	_ = order.NewService(store)

	st, err := stan.Connect(
		stanClusterID,
		clientID,
		stan.NatsURL(os.Getenv("NATS_STREAMING_URL")),
	)

	if err != nil {
		log.Fatal(err)
	}
	defer st.Close()

	if _, err = st.Subscribe("orders", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
		m.Ack()
	}); err != nil {
		return
	}

	fmt.Println("Server is listening...")
	if err := http.ListenAndServe((":8080"), nil); err != nil {
		log.Fatal(err)
	}

}
