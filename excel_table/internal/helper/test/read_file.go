package test

import (
	"encoding/json"
	"excel_table/internal/domain"
	"excel_table/internal/domain/request"
	"excel_table/internal/repository/pb/excel_table"
	"fmt"
	"os"
)

type TestingStruct interface {
	request.TopCadets | excel_table.TopCadetsRequest | domain.TableQueue | map[string]domain.TableQueue
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
