package main

import (
	"L0/order"
	"L0/pkg/storage"
	"context"
	"encoding/json"
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

type natsMessage struct {
	OrderUuid string `json:"order_uid"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	if err = conn.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	defer conn.Close(ctx)

	store := storage.NewStorage(conn)
	if err := store.Load(context.Background()); err != nil {
		log.Println("error load cache from db", err)
	}
	service := order.NewService(store)

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
		var msg natsMessage
		fmt.Printf("Received a message: %s\n", string(m.Data))
		if err := json.Unmarshal(m.Data, &msg); err != nil {
			log.Println(err)
			return
		}

		if err = service.Create(ctx, &storage.Order{
			OrderUuid: msg.OrderUuid,
			Data:      m.Data,
		}); err != nil {
			log.Println("massage could not save: " + err.Error())
			return
		}

		service.Save(&storage.Order{
			OrderUuid: msg.OrderUuid,
			Data:      m.Data,
		})

		m.Ack()
	}); err != nil {
		return
	}

	http.HandleFunc("/orders", func(writer http.ResponseWriter, request *http.Request) {
		uids := request.URL.Query()["uid"]
		if len(uids) != 1 || uids[0] == "" {
			writer.Write([]byte("invalid GET argument uid"))
			return
		}

		o, err := service.ByUUID(ctx, uids[0])
		if err != nil {
			log.Println(err)
			writer.Write([]byte("can't get order from storage"))
			return
		}
		writer.Write(o.Data)
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

//
