package helper

import (
	"excel_table/internal/domain"

	"github.com/xuri/excelize/v2"
)

func AllBorders() []excelize.Border {
	return []excelize.Border{
		{Type: "left", Style: 1, Color: domain.ColorAcaBorder},
		{Type: "top", Style: 1, Color: domain.ColorAcaBorder},
		{Type: "right", Style: 1, Color: domain.ColorAcaBorder},
		{Type: "bottom", Style: 1, Color: domain.ColorAcaBorder},
	}
}

func BordersX() []excelize.Border {
	return []excelize.Border{
		{Type: "top", Style: 1, Color: domain.ColorAcaBorder},
		{Type: "bottom", Style: 1, Color: domain.ColorAcaBorder},
	}
}

func BordersBottom() []excelize.Border {
	return []excelize.Border{
		{Type: "bottom", Style: 1, Color: domain.ColorAcaBorder},
	}
}
