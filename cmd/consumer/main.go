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

	if _, err = st.Subscribe("orders", func(m *stan.Msg) {
		m.Ack()
		fmt.Printf("Received a message: %s\n", string(m.Data))
	}); err != nil {
		return
	}

	fmt.Println("Server is listening...")
	if err := http.ListenAndServe((":8080"), nil); err != nil {
		log.Fatal(err)
	}

}
