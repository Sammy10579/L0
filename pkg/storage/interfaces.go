package storage

import (
	"context"
	"github.com/jackc/pgconn"
)

type Queries interface {
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
}
