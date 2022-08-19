Create table in postgressdb or import.

CREATE TABLE orders
(
id BIGSERIAL PRIMARY KEY,
num VARCHAR(255) NOT NULL UNIQUE,
payload JSON NOT NULL
)


1)Run nats-streaming + nuts ui + pgsql:

docker-compose up -d

2)Run producer:

go run cmd/producer/main.go

3)Run consumer:

go run cmd/consumer/main.go
