// structs for internal use
package top_cadets

import (
	"errors"
	"excel_table/internal/domain"
	"excel_table/internal/domain/request"
	"excel_table/internal/helper"
	"excel_table/pkg/sorted_map"

	"github.com/xuri/excelize/v2"
)

func TopCadetsStorage() *Maps {
	m := Maps{
		// storage for head will save by "head_1" -> "head_2" -> "head_3" -> "head_4"
		// head_1: Module, Checkpoints, Piscine
		// head_2: Languages
		// head_3: xp or event_id
		// head_4: event_name or date
		Head: sorted_map.NewStorage[domain.HeadIndex, domain.Head](100),
		// Data will save by keys - "level" -> "xp_total" -> "user_id" -> map[head_1]map[head_4]user_data
		// need sort: level - desc, xp_total - desc
		Data: sorted_map.NewStorage[int, map[domain.HeadIndex]map[domain.HeadIndex]domain.Cell](100),
	}

	// set first key of head - {Personal, Hours, Module, Checkpoints}
	m.Head.Storage().GetOrSet(domain.IndexPersonal, domain.Head{
		ID:     domain.IndexPersonal,
		Title:  domain.NamePersonal,
		RaiseX: -1, // minus current position
	})

	m.Head.Storage().GetOrSet(domain.IndexHours, domain.Head{
		ID:     domain.IndexHours,
		Title:  domain.NameHours,
		RaiseX: -1, // minus current position
	})

	m.Head.Storage().GetOrSet(domain.IndexModule, domain.Head{
		ID:     domain.IndexModule,
		Title:  domain.NameModule,
		RaiseX: -1, // minus current position
	})

	m.Head.Storage().GetOrSet(domain.IndexCheckpoint, domain.Head{
		ID:     domain.IndexCheckpoint,
		Title:  domain.NameCheckpoint,
		RaiseX: -1, // minus current position
	})

	return &m
}

type Maps struct {
	Head sorted_map.Storage[domain.HeadIndex, domain.Head]                              //
	Data sorted_map.Storage[int, map[domain.HeadIndex]map[domain.HeadIndex]domain.Cell] //
}

func (m *Maps) Parse(req *request.TopCadets, tableQueue map[string]domain.TableQueue) error {
	if len(req.Cadets) == 0 {
		return errors.New("no data for parsing")
	}
	if tableQueue == nil {
		tableQueue = map[string]domain.TableQueue{
			"": {
				Title: domain.NameLangOther,
				Queue: domain.DefaultQueue,
			},
		}
	}
	var fields []fields

	for _, value := range req.Cadets {
		buf := m.Data.Storage().GetOrSet(value.Personal.Level, nil) // need sort
		buf = buf.Next.GetOrSet(value.Personal.XPTotal, nil)        // need sort

		c := currCadet{
			head: m.Head.Storage(),
			data: make(map[domain.HeadIndex]map[domain.HeadIndex]domain.Cell),
			cellStyle: excelize.Style{
				Border: helper.BordersBottom(),
				// Fill: excelize.Fill{
				// 	Type:    "pattern",
				// 	Pattern: 1, // solid
				// 	Color:   []string{value.Personal.Admission.Color},
				// },
				Font: &excelize.Font{
					Size:  domain.DefaultFontSize,
					Color: value.Personal.Admission.Color,
				},
				Alignment: &excelize.Alignment{
					Horizontal: domain.AlignHorizontalLeft,
					Vertical:   domain.AlignVerticalCenter,
				},
				NumFmt: 3, // #,##0
			},
			tq: tableQueue,
		}

		// parse user data, journey
		// make head for journey or update
		if value.Journey != nil {
			// personal
			c.data[domain.IndexPersonal] = make(map[domain.HeadIndex]domain.Cell)
			fields = c.personalParse(value.Personal)

			// module
			if j, ok := value.Journey[domain.NameModule]; ok {
				c.data[domain.IndexModule] = make(map[domain.HeadIndex]domain.Cell)
				c.moduleParse(j)
			}

			// checkpoint
			if j, ok := value.Journey[domain.NameCheckpoint]; ok {
				c.data[domain.IndexCheckpoint] = make(map[domain.HeadIndex]domain.Cell)
				c.checkpointParse(j)
			}

			// hours
			if value.Hours.Total > 0 {
				c.data[domain.IndexHours] = make(map[domain.HeadIndex]domain.Cell)
				c.hoursParse(value.Hours.Total)
			}
		}
		buf.Next.GetOrSet(value.Personal.ID, c.data)
	}

	// make head for personal, hours
	personalHead(fields, m.Head.Storage(), m.Head.Depth())
	hoursHead(req.Period, m.Head.Storage(), m.Head.Depth())

	return nil
}

type currCadet struct {
	head      sorted_map.Maps[domain.HeadIndex, domain.Head]
	data      map[domain.HeadIndex]map[domain.HeadIndex]domain.Cell
	cellStyle excelize.Style
	tq        map[string]domain.TableQueue
}

func (m *Maps) Build() (*domain.Sheet, error) {
	if m.Head.Depth() == 0 {
		return nil, errors.New("storage depth is zero")
	}

	res := result{
		head:    make([][]domain.Head, 0, m.Head.Depth()),
		data:    make([][]domain.Cell, 0, len(m.Data.ReadByDepth(2))), // 2 - index of depth storage, where store all users
		headIDs: make(map[domain.HeadIndex]map[domain.HeadIndex]int, m.Head.Depth()),
	}

	// make head
	res.makeHead(m.Head.Storage(), 0)
	// save ids of head to quick find id
	res.readIDs()
	// make user data
	res.makeData(m.Data.Storage())

	l := 0
	if len(res.data) > 0 {
		l = len(res.data[0])
	} else {
		return nil, errors.New("useful cadet data length is zero")
	}

	return &domain.Sheet{
		Name:   "Top cadets",
		Active: true,
		Head:   res.head,
		Data:   res.data,
		HeadStyle: &excelize.Style{
			Fill: excelize.Fill{
				Type:    "pattern",
				Pattern: 1, // solid
				Color:   []string{domain.ColorAcaHeadFill},
			},
			Border: helper.AllBorders(),
			Font: &excelize.Font{
				Bold:  true,
				Size:  domain.DefaultFontSize,
				Color: domain.ColorAcaHeadFont,
			},
			Alignment: &excelize.Alignment{
				Horizontal: domain.AlignHorizontalCenter,
				Vertical:   domain.AlignVerticalCenter,
			},
		},

		Length: l, // length of useful data
	}, nil
}

func (m *Maps) GetHead() sorted_map.Storage[domain.HeadIndex, domain.Head] {
	return m.Head
}
