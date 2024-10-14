// structs use for parsing requests from other services
package request

import (
	"excel_table/internal/domain"
	"time"
)

// copy of protobuf
type (
	TopCadets struct {
		Period string
		Cadets []CadetData
	}

	CadetData struct {
		Personal Personal
		Hours    Hours
		Journey  map[string][]Event
	}

	Personal struct {
		ID        int
		Login     string
		FullName  string
		Age       int
		Gender    string
		Admission Admission
		Level     int
		XPTotal   int
	}

	Admission struct {
		Name  string
		Color string
	}

	Hours struct {
		Total int
		Month []Month
	}

	Month struct {
		Year  string
		Month string
		Hours int
	}

	Event struct {
		ID        domain.HeadIndex
		Name      string
		XPTotal   int
		CreatedAt time.Time
		Languages []Language
	}

	Language struct {
		Name  string
		Tasks []Task
	}

	Task struct {
		ID   domain.HeadIndex
		Name string
		XP   int
	}
)
