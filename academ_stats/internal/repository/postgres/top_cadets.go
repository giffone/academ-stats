package postgres

import (
	"academ_stats/internal/domain/response"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

const topCadets = "top_cadets"

type Cadets interface {
	TopCadets(ctx context.Context, body []byte) (*response.TopCadetsResponse, error)
}

func (s *storage) TopCadets(ctx context.Context, body []byte) (*response.TopCadetsResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	// check last data
	var (
		bodyLast  []byte
		createdAt time.Time
	)

	query := fmt.Sprintf(`SELECT body, createdAt
	FROM academ_stats.snapshot
	WHERE stat_name = '%s'
	ORDER BY createdAt DESC
	LIMIT 1;`, topCadets)

	if err := s.pool.QueryRow(ctx, query).Scan(&bodyLast, &createdAt); err != nil && err != pgx.ErrNoRows {
		return nil, customErr("query row", err)
	}

	oneWeekAgo := time.Now().AddDate(0, 0, -7)

	if createdAt.After(oneWeekAgo) {
		return &response.TopCadetsResponse{
			Current:  body,
			Last:     bodyLast,
			LastDate: createdAt,
		}, nil
	}

	// add new data
	n, err := s.pool.Exec(ctx,
		`INSERT INTO academ_stats.snapshot (stat_name, body) 
		VALUES ($1, $2);`,
		topCadets,
		body,
	)

	if err != nil {
		return nil, customErr("exec", err)
	}

	if n.RowsAffected() == 0 {
		return nil, response.ErrNotCreated
	}

	return &response.TopCadetsResponse{
		Current:  body,
		Last:     bodyLast,
		LastDate: createdAt,
	}, nil
}
