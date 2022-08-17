package main

import (
	"github.com/nats-io/stan.go"
	"log"
	"os"
)

func main() {
	st, err := stan.Connect(
		"test-cluster",
		"order-consumer",
		stan.NatsURL(os.Getenv("NATS_STREAMING_URL")),
	)

	if err != nil {
		log.Fatal(err)
	}
	defer st.Close()
}
