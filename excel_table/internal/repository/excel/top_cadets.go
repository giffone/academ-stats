package excel

import (
	"excel_table/internal/domain"
	"excel_table/internal/helper"
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

func ExcelFile(data []*domain.Sheet) (*excelize.File, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("no data to create file")
	}
	// create new excel file
	ef := eFile{
		file: excelize.NewFile(),
	}

	for _, v := range data {
		index, err := ef.file.NewSheet(v.Name)
		if err != nil {
			return nil, err
		}
		if v.Active {
			// set main list
			ef.file.SetActiveSheet(index)
		}
		if err = ef.excelFile(v); err != nil {
			return nil, err
		}
	}

	return ef.file, nil
}

type eFile struct {
	file *excelize.File
}

func (e *eFile) excelFile(data *domain.Sheet) error {
	// set header style
	hStyle, err := e.file.NewStyle(data.HeadStyle)
	if err != nil {
		log.Printf("excelize: set style: header: %s\n", err)
		hStyle = -1
	}

	// generate excel alphabet by length data, example [A, ..., ZZ]
	ad := helper.Addr{
		Alphabet: helper.GenExcelColumnAlphabet(data.Length),
	}

	excelY := 1 // excel row starts from 1, not from 0
	excelX := 0 // this column will also contain the cell raiseX

	// add header
	for y := 0; y < len(data.Head); y++ {
		for x := 0; x < len(data.Head[y]); x++ {
			head := data.Head[y][x]
			// make excel column
			topLeft, bottomRight, merge := ad.Addr(excelY, head.RaiseY, excelX, head.RaiseX)
			excelX += head.RaiseX

			// set cell data
			if err = e.file.SetCellValue(data.Name, topLeft, head.Title); err != nil {
				return fmt.Errorf("excelize: set cell value: sheet [%s]: head [%d][%d]: with data \"%s\": %v", data.Name, y, x, head.Title, err)
			}
			// need customize
			if hStyle >= 0 {
				err = e.file.SetCellStyle(data.Name, topLeft, bottomRight, hStyle)
				if err != nil {
					log.Printf("excelize: set cell style: sheet [%s]: head [%d][%d]: with data \"%s\": %v", data.Name, y, x, head.Title, err)
				}
			}
			// need merge
			if merge {
				err = e.file.MergeCell(data.Name, topLeft, bottomRight)
				if err != nil {
					log.Printf("excelize: merge cell: sheet [%s]: head [%d][%d]: with data \"%s\": %v", data.Name, y, x, head.Title, err)
				}
			}
			excelX++
		}
		excelY++
		excelX = 0
	}

	// clear
	excelX = 0

	// add user data
	for y := 0; y < len(data.Data); y++ {
		for x := 0; x < len(data.Data[y]); x++ {
			cadet := data.Data[y][x]

			// make excel column
			topLeft, bottomRight, merge := ad.Addr(excelY, cadet.RaiseY, excelX, cadet.RaiseX)
			excelX += cadet.RaiseX

			// set cell data
			if err = e.file.SetCellValue(data.Name, topLeft, cadet.GetData()); err != nil {
				return fmt.Errorf("excelize: set cell value: sheet [%s]: cadet [%d][%d]: with data \"%s\": %v", data.Name, y, x, cadet.GetData(), err)
			}

			// set cell style if nil
			if cadet.Style == nil {
				cadet.Style = &excelize.Style{
					Border: helper.BordersBottom(),
				}
			}

			// get style
			style, err := e.file.NewStyle(cadet.Style)
			if err != nil {
				log.Printf("excelize: set style: cadet: %s\n", err)
				style = -1
			}
			err = e.file.SetCellStyle(data.Name, topLeft, bottomRight, style)
			if err != nil {
				log.Printf("excelize: set cell style: sheet [%s]: cadet [%d][%d]: with data \"%s\": %v", data.Name, y, x, cadet.GetData(), err)
			}

			// need merge
			if merge {
				err = e.file.MergeCell(data.Name, topLeft, bottomRight)
				if err != nil {
					log.Printf("excelize: merge cell: sheet [%s]: cadet [%d][%d]: with data \"%s\": %v", data.Name, y, x, cadet.GetData(), err)
				}
			}
			excelX++
		}
		excelY++
		excelX = 0
	}

	return nil
}
