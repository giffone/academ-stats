package server

import (
	"academ_stats/internal/config"
	"academ_stats/internal/repository/postgres"
	"academ_stats/internal/repository/zero_one_api"
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Env struct {
	pool  *pgxpool.Pool
	zoApi zero_one_api.ZeroOneApi
}

func NewEnv(ctx context.Context, cfg *config.Config) *Env {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	zoApi, err := zero_one_api.NewZeroOneApi(cfg)
	if err != nil {
		log.Fatalf("environment: %s", err)
	}

	return &Env{
		pool:  postgres.NewPostgres(ctx, cfg),
		zoApi: zoApi,
	}
}

func (e *Env) Stop(ctx context.Context) {
	// ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	// defer cancel()

	e.pool.Close()
	log.Println("envorinments stopped")
}
