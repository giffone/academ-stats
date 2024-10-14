package postgres

import (
	"academ_stats/internal/domain"
	"context"
	"time"
)

type Piscine interface {
	PiscineList(ctx context.Context) (map[int]domain.Piscine, error)
}

func (s *storage) PiscineList(ctx context.Context) (map[int]domain.Piscine, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	rows, err := s.pool.Query(ctx, "select id, title, color from env_tracker.piscines")
	if err != nil {
		return nil, customErr("query", err)
	}
	defer rows.Close()

	var (
		id    int
		title string
		color string
	)

	piscines := make(map[int]domain.Piscine, 50)

	for rows.Next() {
		if err := rows.Scan(
			&id,
			&title,
			&color,
		); err != nil {
			return nil, customErr("in iterate row", err)
		}
		piscines[id] = domain.Piscine{
			ID:    id,
			Title: title,
			Color: color,
		}
	}

	if err = rows.Err(); err != nil {
		return nil, customErr("rows at all", err)
	}

	return piscines, nil
}
