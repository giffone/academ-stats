package test

import (
	"excel_table/internal/domain"
	"excel_table/internal/helper/parser"
	"excel_table/internal/repository/pb/excel_table"
	"excel_table/internal/service/top_cadets"
	"fmt"
)

var Period string

func Build(dirTc, dirQue string) (*domain.Sheet, error) {
	srcTc, err := ReadFile[excel_table.TopCadetsRequest](dirTc)
	if err != nil {
		return nil, fmt.Errorf("can not read src top cadets file: %s", err)
	}
	Period = srcTc.GetPeriod()

	// read table queue
	srcQ, err := ReadFile[map[string]domain.TableQueue](dirQue)
	if err != nil {
		return nil, fmt.Errorf("can not read src table queue file: %s", err)
	}

	storageTopCadets := top_cadets.TopCadetsStorage()

	src, err := parser.TopCadets(srcTc)
	if err != nil {
		return nil, fmt.Errorf("parser request: %s", err)
	}

	if err = storageTopCadets.Parse(src, *srcQ); err != nil {
		return nil, fmt.Errorf("parser storage: %s", err)
	}

	res, err := storageTopCadets.Build()
	if err != nil {
		return nil, fmt.Errorf("build: %s", err)
	}

	return res, nil
}
