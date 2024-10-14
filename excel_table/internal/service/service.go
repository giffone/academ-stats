package service

import (
	"context"
	"excel_table/internal/domain"
	"excel_table/internal/repository/postgres"
	"fmt"
)

type Service interface {
	GetQueue(ctx context.Context) (map[string]domain.TableQueue, error)
}

func New(storage postgres.Storage) Service {
	return &service{
		storage: storage,
	}
}

type service struct {
	storage postgres.Storage
}

func (s *service) GetQueue(ctx context.Context) (map[string]domain.TableQueue, error) {
	q, err := s.storage.GetQueue(ctx)
	if err != nil {
		return nil, fmt.Errorf("get queue %s", err)
	}

	return q, nil
}
