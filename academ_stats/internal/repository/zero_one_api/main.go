package zero_one_api

import (
	"academ_stats/internal/config"
	"fmt"
	"log"
	"time"

	zero "github.com/giffone/zerone_api/v2"
)

const RefreshTokenPreNotify = 100 * time.Hour

type ZeroOneApi interface {
	TopCadets
	Token() (int64, string, error)
}

func NewZeroOneApi(cfg *config.Config) (ZeroOneApi, error) {
	log.Println("[zero-one-api] init...")
	log.Println("[zero-one-api] conn...")

	cli, err := zero.CreateClient(cfg.ZeroAddr, cfg.ZeroToken, RefreshTokenPreNotify, cfg.Debug)
	if err != nil {
		return nil, fmt.Errorf("zero one: client create: %w", err)
	}

	log.Println("[zero-one-api] init done...")
	return &zeroOneApi{
		cli: cli,
	}, nil
}

type zeroOneApi struct {
	cli zero.Client
}

func (z *zeroOneApi) Token() (int64, string, error) {
	return z.cli.TokenBase()
}
