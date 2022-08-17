package main

import (
	"github.com/google/uuid"
	"github.com/nats-io/stan.go"
	"log"
	"os"
	"time"
)

func main() {
	sc, err := stan.Connect(
		"test-cluster",
		"order_producer",
		stan.NatsURL(os.Getenv("NATS_STREAMING_URL")),
	)
	if err != nil {
		log.Println(err)

		return
	}
	defer sc.Close()

	t := time.NewTicker(time.Second)
	for range t.C {
		log.Println("send order to orders")

		if err = sc.Publish("orders", genOrder()); err != nil {
			log.Fatal(err)
		}
	}

}
func genOrder() []byte {
	orderUid := uuid.New()
	trackNumber := uuid.New()
	return []byte("{" +
		" \"order_uid\": \"" + orderUid.String() + "\"," +
		"  \"track_number\": \"" + trackNumber.String() + "\"," +
		"  \"entry\": \"WBIL\"," +
		"  \"delivery\": {\n    \"name\": \"Test Testov\"," +
		"    \"phone\": \"+9720000000\"," +
		"    \"zip\": \"2639809\"," +
		"    \"city\": \"Kiryat Mozkin\"," +
		"    \"address\": \"Ploshad Mira 15\"," +
		"    \"region\": \"Kraiot\"," +
		"    \"email\": \"test@gmail.com\"\n  }," +
		"  \"payment\": {\n    \"transaction\": \"b563feb7b2b84b6test\"," +
		"    \"request_id\": \"\"," +
		"    \"currency\": \"USD\"," +
		"    \"provider\": \"wbpay\"," +
		"    \"amount\": 1817," +
		"    \"payment_dt\": 1637907727," +
		"    \"bank\": \"alpha\"," +
		"    \"delivery_cost\": 1500," +
		"    \"goods_total\": 317," +
		"    \"custom_fee\": 0\n  }," +
		"  \"items\": [\n    {\n      \"chrt_id\": 9934930," +
		"      \"track_number\": \"WBILMTESTTRACK\"," +
		"      \"price\": 453," +
		"      \"rid\": \"ab4219087a764ae0btest\"," +
		"      \"name\": \"Mascaras\"," +
		"      \"sale\": 30,\n      \"size\": \"0\"," +
		"      \"total_price\": 317," +
		"      \"nm_id\": 2389212," +
		"      \"brand\": \"Vivienne Sabo\"," +
		"     \"status\": 202\n    }\n  ]," +
		"  \"locale\": \"en\"," +
		" \"internal_signature\": \"\"," +
		"  \"customer_id\": \"test\"," +
		"  \"delivery_service\": \"meest\"," +
		" \"shardkey\": \"9\"," +
		"  \"sm_id\": 99," +
		"  \"date_created\": \"2021-11-26T06:22:19Z\"," +
		"  \"oof_shard\": \"1\"}")
}
