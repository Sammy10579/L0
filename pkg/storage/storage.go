package storage

type Storage struct {
	db Queries
}

func NewStorage(conn Queries) *Storage {
	return &Storage{db: conn}
}
