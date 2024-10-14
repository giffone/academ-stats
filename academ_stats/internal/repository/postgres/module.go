package postgres

import (
	"academ_stats/internal/domain"
	"context"
	"time"
)

type Module interface {
	ModuleList(ctx context.Context) ([]domain.Module, error)
}

func (s *storage) ModuleList(ctx context.Context) ([]domain.Module, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	rows, err := s.pool.Query(ctx, `SELECT id, title
	FROM env_tracker.divs
	WHERE id >= 66`) // first created module is 66
	if err != nil {
		return nil, customErr("query", err)
	}

	defer rows.Close()

	var (
		id    int
		title string
	)

	modules := make([]domain.Module, 0, 5)

	for rows.Next() {
		if err := rows.Scan(
			&id,
			&title,
		); err != nil {
			return nil, customErr("in iterate row", err)
		}
		modules = append(modules, domain.Module{
			ID:    id,
			Title: title,
		})
	}

	if err = rows.Err(); err != nil {
		return nil, customErr("rows at all", err)
	}

	return modules, nil
}
