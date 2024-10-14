package test

import (
	"academ_stats/internal/domain"
	"academ_stats/internal/domain/request"
	"academ_stats/internal/domain/response"
	"academ_stats/internal/repository/pb/excel_table"
	"academ_stats/internal/repository/pb/session_manager"

	"encoding/json"
	"fmt"
	"os"
)

type TestingStruct interface {
	request.TopCadets | domain.TopCadets | response.TopCadets | domain.HoursDTO |
		session_manager.CadetsTimeResponse | excel_table.TopCadetsRequest |
		map[int]domain.HoursDTO
}

func ReadFile[T TestingStruct](file_dir string) (*T, error) {
	var tc T

	file, err := os.ReadFile(file_dir)
	if err != nil {
		return nil, fmt.Errorf("open file: %s", err)
	}

	err = json.Unmarshal(file, &tc)
	if err != nil {
		return nil, fmt.Errorf("unmarshal: %s", err)
	}

	return &tc, nil
}
