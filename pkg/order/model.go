package order

type Order struct {
	ID      int    `db:"id"`
	Uid     string `db:"uid"`
	Payload []byte `db:"payload"`
}
