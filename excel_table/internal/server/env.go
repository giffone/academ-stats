package server

import (
	"context"
	"excel_table/config"
	"excel_table/internal/repository/postgres"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Env struct {
	pool *pgxpool.Pool
}

func NewEnv(ctx context.Context, cfg *config.Config) *Env {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	return &Env{
		pool: postgres.NewPostgres(ctx, cfg),
	}
}

func (e *Env) Stop(ctx context.Context) {
	// ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	// defer cancel()

	e.pool.Close()
	log.Println("envorinments stopped")
}
