package service

import (
	"bytes"
	"context"
	"excel_table/internal/domain"
	"excel_table/internal/domain/response"
	"excel_table/internal/helper/parser"
	"excel_table/internal/repository/excel"
	"excel_table/internal/repository/pb/excel_table"
	"excel_table/internal/repository/postgres"
	"excel_table/internal/service/top_cadets"
	"fmt"
	"log"
)

// TOP CADETS

//          ----------------------------------------------------------------------------------------------------------------------------
//  head 1  | Personal                   | Hours         | Module                                      | Checkpoint           | Piscine |  1 title
//          |                            ------------------------------------------------------------------------------------------------
//  head 2  |                            |               | Go                             | other      | Go    | Go     | ... |   ...   |  2 language
//          |                            ------------------------------------------------------------------------------------------------
//  head 3  |                            |               | 5000xp      | 10000xp   | ...  | 5000xp     | #74   | #75    |     |   ...   |  3 xp / id
//          -----------------------------------------------------------------------------------------------------------------------------
//  head 4  | UserID | Login | Age | ... | Jan24 - May24 | go-reloaded | ascii-art | a... | math-skils | 8 feb | 14 feb |     |         |  4 user fields / task name / date
//          -----------------------------------------------------------------------------------------------------------------------------
//  data    | user   | login | 21  | ... |      345      | done        | done      |      | done       | 3     | 7      |     |         |
//          -----------------------------------------------------------------------------------------------------------------------------

func NewServiceGrpc(storage postgres.Storage) *ServiceGrpc {
	return &ServiceGrpc{storage: storage}
}

type ServiceGrpc struct {
	storage postgres.Storage
	excel_table.UnimplementedExcelTableServer
}

func (s *ServiceGrpc) GetTopCadets(ctx context.Context, in *excel_table.TopCadetsRequest) (*excel_table.TopCadetsResponse, error) {
	// check request
	if in == nil || len(in.Cadets) == 0 {
		return nil, response.ErrReqEmpty
	}

	// parse data
	src, err := parser.TopCadets(in)
	if err != nil {
		return nil, fmt.Errorf("get top cadets: %v", err)
	}

	// get indexes for excel table to sort data
	queue, err := s.storage.GetQueue(ctx)
	if err != nil {
		return nil, fmt.Errorf("get top cadets: get queue: %v", err)
	}

	// create storage for cadets data [like linked list]
	// this is needed to save the data in the hierarchy chain and then sort it later
	storage := top_cadets.TopCadetsStorage()
	// read src, queue and save into storage
	if err = storage.Parse(src, queue); err != nil {
		return nil, fmt.Errorf("get top cadets: parse: %v", err)
	}

	// top cadets sheet
	sheet, err := storage.Build()
	if err != nil {
		return nil, fmt.Errorf("get top cadets: build: %v", err)
	}

	// create file
	file, err := excel.ExcelFile([]*domain.Sheet{sheet})
	if err != nil {
		return nil, fmt.Errorf("get top cadets: excel file: %v", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Println(err)
		}
	}()

	// write file to buffer for making []byte
	var buffer bytes.Buffer
	if err := file.Write(&buffer); err != nil {
		return nil, fmt.Errorf("get top cadets: excel file: write to buffer: %v", err)
	}

	return &excel_table.TopCadetsResponse{
		Message: "Created",
		File:    buffer.Bytes(),
	}, nil
}
