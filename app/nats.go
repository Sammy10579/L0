package app

import (
	"context"
	"encoding/json"
	"log"

	"L0/pkg/order"
	"github.com/nats-io/stan.go"
)

type natsMessage struct {
	OrderUuid string `json:"order_uid"`
}

func runNatsListener(conn stan.Conn, service OrderService) {
	if _, err := conn.Subscribe("orders", func(m *stan.Msg) {
		ctx := context.Background()
		var msg natsMessage
		log.Printf("Received a message: %s\n\n", string(m.Data))
		if err := json.Unmarshal(m.Data, &msg); err != nil {
			log.Println(err)
			return
		}

		if err := service.Save(ctx, order.Order{
			Uid:     msg.OrderUuid,
			Payload: m.Data,
		}); err != nil {
			log.Println("message could not be saved: " + err.Error())
			return
		}
	}, stan.StartWithLastReceived()); err != nil {
		log.Fatalf("cant subscribe to topic %s", "orders")
	}
}
