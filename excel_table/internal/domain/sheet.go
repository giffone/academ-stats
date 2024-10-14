package domain

import (
	"time"

	"github.com/xuri/excelize/v2"
)

type Sheet struct {
	Name      string          // name of sheet
	Active    bool            // is active sheet in excel file
	Head      [][]Head        //
	Data      [][]Cell        // useful data
	HeadStyle *excelize.Style //
	Length    int             // length of []Cell
}

type Head struct {
	ID     HeadIndex // identify data
	Title  string    // useful data
	RaiseX int       // add cells by horisontal
	RaiseY int       // add cells by vertical
	Main   *Head     // head_1 (need to check Main ID)
}

type Cell struct {
	ID     HeadIndex       // identify data
	data   any             // useful data
	RaiseX int             // add cells by horisontal
	RaiseY int             // add cells by vertical
	Style  *excelize.Style //
}

// it need for use only allowed excel types
func (c *Cell) SetData(v any) {
	if v == nil {
		c.data = CellValueNil
		return
	}

	if checkExcelType(v) {
		c.data = v
	} else {
		c.data = CellValueTypeNotAvailable
	}
}
func (c *Cell) GetData() interface{} {
	return c.data
}

type Style struct {
	FontColor       string  //
	FontSize        float64 //
	Boards          bool    // all boards
	Bold            bool    //
	FillColor       string  //
	AlignmentCenter bool    //
}

func checkExcelType(val any) bool {
	switch val.(type) {
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64,
		string,
		[]byte,
		time.Duration,
		time.Time,
		bool:
		return true
	default:
		return false
	}
}
