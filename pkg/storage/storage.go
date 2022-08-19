package storage

type Storage struct {
	db    Queries
	cache map[int]string
}

func NewStorage(db Queries) *Storage {
	return &Storage{db: db}
}
