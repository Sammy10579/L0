package main

import (
	"L0/pkg/storage"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"os"
)

const (
	stanClusterID = "test-cluster"
	clientID      = "order-consumer"
)

type natsMessage struct {
	TrackNumber string `json:"track_number"`
}

func main() {
	st, err := stan.Connect(
		stanClusterID,
		clientID,
		stan.NatsURL(os.Getenv("NATS_STREAMING_URL")),
	)

	if err != nil {
		log.Fatal(err)
	}
	defer st.Close()

	sub, err := st.Subscribe("orders", func(m *stan.Msg) {
		order := &storage.Order{}
		err := json.Unmarshal(m.Data, &order)
		if err != nil {
			log.Print(err)
			return
		}

		fmt.Printf("Received a message: %s\n", string(m.Data))

	})

	if err != nil {
		fmt.Println("Subscribe is not connected")
	}

	sub.Unsubscribe()
	st.Close()
}
