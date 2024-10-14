package parser

import (
	"excel_table/internal/domain"
	"excel_table/internal/domain/request"
	"excel_table/internal/domain/response"
	"excel_table/internal/repository/pb/excel_table"
)

func TopCadets(req *excel_table.TopCadetsRequest) (*request.TopCadets, error) {
	if req == nil {
		return nil, response.ErrReqEmpty
	}

	data := req.GetCadets()
	if data == nil {
		return nil, response.ErrReqCadetDataEmpty
	}

	if len(req.Cadets) == 0 {
		return nil, response.ErrNoData
	}

	tc := request.TopCadets{
		Period: req.GetPeriod(),
		Cadets: make([]request.CadetData, len(data)),
	}

	for i, cadet := range data {
		tc.Cadets[i] = request.CadetData{
			Personal: personalParse(cadet),
			Hours:    hoursParse(cadet),
			Journey:  journeyParse(cadet),
		}
	}

	return &tc, nil
}

func personalParse(data *excel_table.CadetData) request.Personal {
	pr := data.GetPersonal()
	if pr == nil {
		return request.Personal{}
	}
	p := request.Personal{
		ID:       int(pr.GetId()),
		Login:    pr.GetLogin(),
		FullName: pr.GetFullName(),
		Age:      int(pr.GetAge()),
		Gender:   pr.GetGender(),
		Level:    int(pr.GetLevel()),
		XPTotal:  int(pr.GetXpTotal()),
	}

	if pr.GetAdmission() != nil {
		p.Admission = request.Admission{
			Name:  pr.Admission.GetName(),
			Color: pr.Admission.GetColor(),
		}
	}

	return p
}

func hoursParse(data *excel_table.CadetData) request.Hours {
	h := data.GetHours()
	if h == nil {
		return request.Hours{}
	}

	lh := len(h.GetMonth())
	if lh == 0 {
		return request.Hours{
			Total: int(h.GetTotal()),
		}
	}

	m := make([]request.Month, lh)

	for i := 0; i < lh; i++ {
		m[i] = request.Month{
			Year:  h.Month[i].GetYear(),
			Month: h.Month[i].GetMonth(),
			Hours: int(h.Month[i].GetHours()),
		}
	}

	return request.Hours{
		Total: int(h.GetTotal()),
		Month: m,
	}
}

func journeyParse(data *excel_table.CadetData) map[string][]request.Event {
	jr := data.GetJourney()
	if jr == nil {
		return make(map[string][]request.Event, 0)
	}

	m := make(map[string][]request.Event)

	for k, v := range jr {
		if v == nil || len(v.GetEvents()) == 0 {
			continue
		}
		m[k] = eventsParse(v.GetEvents())
	}

	return m
}

func eventsParse(data []*excel_table.Event) []request.Event {
	e := make([]request.Event, 0, len(data))

	for i := 0; i < len(data); i++ {
		ev := data[i]
		if ev == nil {
			continue
		}
		e = append(e, request.Event{
			ID:        domain.HeadIndex(ev.GetId()),
			Name:      ev.GetName(),
			XPTotal:   int(ev.GetXpTotal()),
			CreatedAt: ev.GetCreatedAt().AsTime(),
			Languages: languagesParse(ev.GetLanguages()),
		})
	}

	return e
}

func languagesParse(data []*excel_table.Language) []request.Language {
	if data == nil {
		return nil
	}

	l := make([]request.Language, 0, len(data))

	for i := 0; i < len(data); i++ {
		lang := data[i]
		if lang == nil {
			continue
		}
		l = append(l, request.Language{
			Name:  lang.GetName(),
			Tasks: tasksParse(lang.GetTasks()),
		})
	}

	return l
}

func tasksParse(data []*excel_table.Task) []request.Task {
	if data == nil {
		return nil
	}

	t := make([]request.Task, 0, len(data))

	for i := 0; i < len(data); i++ {
		task := data[i]
		if task == nil {
			continue
		}
		t = append(t, request.Task{
			ID:   domain.HeadIndex(task.GetId()),
			Name: task.GetName(),
			XP:   int(task.GetXp()),
		})
	}

	return t
}
