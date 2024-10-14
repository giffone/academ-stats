package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Storage interface {
	TableQueue
}

// Pg need to use pgx or pgxmock
type Pg interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Close()
}

func NewStorage(pool Pg) Storage {
	return &storage{pool: pool}
}

type storage struct {
	pool Pg
}