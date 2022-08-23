 # Create table in postgress db.
```
CREATE TABLE orders
(
id BIGSERIAL PRIMARY KEY,
orderuuid VARCHAR(255) NOT NULL UNIQUE,
data JSON NOT NULL
)
```
# Run nats-streaming + nuts ui + pgsql:
```
docker-compose up -d
```
# Run producer:
```
go run cmd/producer/main.go
```
# Run consumer:
```
go run cmd/consumer/main.go
```

