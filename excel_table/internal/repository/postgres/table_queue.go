package postgres

import (
	"context"
	"excel_table/internal/domain"
	"fmt"
	"time"
)

const (
	QueryGetQueue = "SELECT value, value_title, queue FROM excel_table.table_queue WHERE value_type = '%s';"
)

type TableQueue interface {
	GetQueue(ctx context.Context) (map[string]domain.TableQueue, error)
	CreateQueue(ctx context.Context, value string, dto domain.TableQueue) error
	CreateQueueDefault(ctx context.Context, value string) error
}

func (s *storage) GetQueue(ctx context.Context) (map[string]domain.TableQueue, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	q := fmt.Sprintf(QueryGetQueue, domain.SqlFieldLanguage)

	rows, err := s.pool.Query(ctx, q)
	if err != nil {
		return nil, customErr("query", err)
	}
	defer rows.Close()

	queue := make(map[string]domain.TableQueue, 20)
	var value string

	for rows.Next() {
		q := domain.TableQueue{}
		if err := rows.Scan(
			&value,
			&q.Title,
			&q.Queue,
		); err != nil {
			return nil, customErr("in iterate row", err)
		}
		queue[value] = q
	}

	if err = rows.Err(); err != nil {
		return nil, customErr("rows at all", err)
	}

	return queue, nil
}

func (s *storage) CreateQueue(ctx context.Context, value string, dto domain.TableQueue) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// insert
	if _, err := s.pool.Exec(ctx,
		`INSERT INTO 
		excel_table.table_queue (value, value_title, queue, value_type) 
		VALUES ($1, $2, $3, $4);`,
		value,
		dto.Title,
		dto.Queue,
		domain.SqlFieldLanguage,
	); err != nil {
		return customErr("exec", err)
	}

	return nil
}

func (s *storage) CreateQueueDefault(ctx context.Context, value string) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// insert only value, other - defaul values [title_value = "other", queue = 1000]
	if _, err := s.pool.Exec(ctx,
		`INSERT INTO 
		excel_table.table_queue (value, value_type, queue) 
		VALUES ($1, $2, $3);`,
		value,
		domain.SqlFieldLanguage,
		domain.DefaultQueue,
	); err != nil {
		return customErr("exec", err)
	}

	return nil
}
