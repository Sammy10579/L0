package main

import (
	"fmt"
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
	st, err := stan.Connect(
		stanClusterID,
		clientID,
		stan.NatsURL(os.Getenv("NATS_STREAMING_URL")),
	)

	if err != nil {
		log.Fatal(err)
	}
	defer st.Close()

	/*	if _, err = st.Subscribe("orders", func(msg *stan.Msg) {
			order := &storage.Order{}
			massage := json.Unmarshal(order.Data, &order)
		}); err != nil {
			return
		}
		fmt.Printf("Received a message: %s\n", massage)*/

	fmt.Println("Server is listening...")
	if err := http.ListenAndServe((":8080"), nil); err != nil {
		log.Fatal(err)
	}

}
