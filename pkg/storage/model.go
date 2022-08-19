package storage

type Order struct {
	ID          int    `db:"id"`
	OrderUuid   string `db:"order_uid"`
	TrackNumber string `db:"track_number"`
	Data        []byte `db:"order_data"`
}
