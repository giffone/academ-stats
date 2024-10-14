package parser

import (
	"academ_stats/internal/domain"
	"academ_stats/internal/domain/request"

	"sync"
)

// TopCadetsGraphql reads data from graphql request and copy data to internal struct TopCadets
func TopCadetsGraphql(data *request.TopCadets, adm map[int]domain.Piscine, t map[int]domain.HoursDTO, pr *domain.Period) *domain.TopCadets {
	wg := sync.WaitGroup{}

	period := ""

	if pr != nil {
		period = pr.String()
	}

	tc := domain.TopCadets{
		Period: period,
		Cadets: make([]domain.CadetData, len(data.Cadets)),
	}

	for i, cadet := range data.Cadets {
		cd := domain.CadetData{
			Personal: PersonalGraphql(&cadet, adm),
			Journey:  make(map[string][]domain.EventDTO),
		}

		chData := make(chan map[string][]domain.EventDTO, len(cadet.Journey))

		for _, v := range cadet.Journey {
			if len(v.Paths) == 0 {
				continue
			}
			wg.Add(1)
			go parse(&wg, v, chData)
		}

		wg.Wait()
		// need close channel for use range
		close(chData)

		// save events {module, checkpoint, piscine}
		// range channel to catch data
		for data := range chData {
			for k, v := range data {
				cd.Journey[k] = v
			}
		}

		if h, ok := t[cadet.UserID]; ok {
			cd.Hours = h
		}

		tc.Cadets[i] = cd
	}

	return &tc
}

// buffer struct for saving the order of "registration_id" index and included "language"
// like map[registration_id]regSlice
type regBuf struct {
	regID map[int]*regSlice
}

type regSlice struct {
	i     int            // index in slice
	iLang map[string]int // index in slice for languages -> map[language]index
}

type HeaderList struct {
	Name  string
	IDGen int
	ID    int
}

func parse(wg *sync.WaitGroup, obj request.Object, ch chan<- map[string][]domain.EventDTO) {
	defer wg.Done()

	if len(obj.Paths) == 0 {
		return
	}

	data := []domain.EventDTO{}
	buf := regBuf{regID: make(map[int]*regSlice)}

	eventName := ""

	for i, path := range obj.Paths {
		if eventName == "" {
			eventName = path.Event.Object.EventName
		}

		// find or create new event
		newObj, indexO := buf.findEvent(&path)
		if newObj != nil {
			if eventName != domain.NameCheckpoint {
				newObj.XPTotal = int(obj.XPTotal.Sum.Total)
			}
			data = append(data, *newObj)
		}

		// add language
		newLang, indexL := buf.regID[path.Event.RegistrationID].findLanguage(&obj, i)
		if newLang != nil {
			data[indexO].Languages = append(data[indexO].Languages, *newLang)
		}

		// add XP
		data[indexO].Languages[indexL].Tasks = append(data[indexO].Languages[indexL].Tasks,
			domain.TaskDTO{
				ID:   path.Path.Attrs.ID,
				XP:   int(path.XP),
				Name: path.Path.Attrs.PathName,
			},
		)

		if eventName == domain.NameCheckpoint {
			// calc manually total xp
			data[indexO].XPTotal += int(path.XP)
		}
	}

	ch <- map[string][]domain.EventDTO{
		eventName: data,
	}
}

func (b *regBuf) findEvent(path *request.Path) (*domain.EventDTO, int) {
	// check if a "registration_id" exists in the list
	s, ok := b.regID[path.Event.RegistrationID]
	// create object
	if !ok {
		o := domain.EventDTO{
			Name:      path.Event.Object.EventName,
			ID:        path.Event.RegistrationID,
			CreatedAt: path.Event.CreatedAt,
			Languages: make([]domain.LanguageDTO, 0, 10), // length ~ how many languages in platform
		}

		newIndex := len(b.regID)

		// save index of registration_id
		b.regID[path.Event.RegistrationID] = &regSlice{
			i:     newIndex,
			iLang: make(map[string]int, 10),
		}

		return &o, newIndex
	}

	return nil, s.i
}

// for all
func (b *regSlice) findLanguage(rObj *request.Object, pathID int) (*domain.LanguageDTO, int) {
	lang := rObj.Paths[pathID].Path.Attrs.Language
	if lang == "" {
		lang = domain.NameLangOther
	}
	// check if a language exists in the list
	iLang, ok := b.iLang[lang]
	// create language
	if !ok {
		l := domain.LanguageDTO{
			Name:  lang,
			Tasks: make([]domain.TaskDTO, 0, len(rObj.Paths)), // length ~ how many records in request
		}

		newIndex := len(b.iLang)

		// save index of language
		b.iLang[lang] = newIndex
		// update changes
		iLang = b.iLang[lang]

		return &l, iLang
	}

	return nil, iLang
}
