package parser

import (
	"academ_stats/internal/domain"
	"academ_stats/internal/repository/pb/excel_table"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// TopCadetsGrpc reads internal struct TopCadets and copy data to grpc
func TopCadetsGrpc(tc *domain.TopCadets) *excel_table.TopCadetsRequest {
	if tc == nil {
		return nil
	}
	req := excel_table.TopCadetsRequest{
		Period: tc.Period,
		Cadets: make([]*excel_table.CadetData, len(tc.Cadets)),
	}

	for i, v := range tc.Cadets {
		req.Cadets[i] = &excel_table.CadetData{
			Personal: parsePersonalGrpc(&v.Personal),
			Hours:    parseHoursGrpc(&v.Hours),
			Journey:  parseJourneyGrpc(v.Journey),
		}
	}

	return &req
}

func parsePersonalGrpc(p *domain.Personal) *excel_table.Personal {
	if p == nil {
		return nil
	}

	return &excel_table.Personal{
		Id:       int32(p.ID),
		Login:    p.Login,
		FullName: p.FullName,
		Age:      int32(p.Age),
		Gender:   p.Gender,
		Admission: &excel_table.Admission{
			Name:  p.Admission.Name,
			Color: p.Admission.Color,
		},
		Level:   int32(p.Level),
		XpTotal: int32(p.XPTotal),
	}
}

func parseHoursGrpc(h *domain.HoursDTO) *excel_table.Hours {
	if h == nil {
		return nil
	}

	m := make([]*excel_table.MonthStr, len(h.Month))

	for i, v := range h.Month {
		m[i] = &excel_table.MonthStr{
			Year:  v.Year,
			Month: v.Month,
			Hours: int32(v.Hours),
		}
	}

	return &excel_table.Hours{
		Total: int32(h.Total),
		Month: m,
	}
}

func parseJourneyGrpc(journeys map[string][]domain.EventDTO) map[string]*excel_table.EventList {
	eList := make(map[string]*excel_table.EventList, len(journeys))

	for k, j := range journeys {
		eList[k] = &excel_table.EventList{
			Events: make([]*excel_table.Event, len(j)),
		}

		for i, e := range j {
			eList[k].Events[i] = &excel_table.Event{
				Id:        int32(e.ID),
				Name:      e.Name,
				XpTotal:   int32(e.XPTotal),
				CreatedAt: timestamppb.New(e.CreatedAt),
				Languages: parseLanguageGrpc(e.Languages),
			}
		}
	}

	return eList
}

func parseLanguageGrpc(languages []domain.LanguageDTO) []*excel_table.Language {
	lang := make([]*excel_table.Language, len(languages))

	for i, l := range languages {
		lang[i] = &excel_table.Language{
			Name:  l.Name,
			Tasks: parseTaskGrpc(l.Tasks),
		}
	}

	return lang
}

func parseTaskGrpc(tasks []domain.TaskDTO) []*excel_table.Task {
	ts := make([]*excel_table.Task, len(tasks))

	for i, t := range tasks {
		ts[i] = &excel_table.Task{
			Id:   int32(t.ID),
			Name: t.Name,
			Xp:   int32(t.XP),
		}
	}

	return ts
}
