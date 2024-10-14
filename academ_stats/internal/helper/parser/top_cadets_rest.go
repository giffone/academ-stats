package parser

import (
	"academ_stats/internal/domain"
	"academ_stats/internal/domain/response"
)

// TopCadetsRest reads internal struct TopCadets and copy data to response struct
func TopCadetsRest(tc *domain.TopCadets) *response.TopCadets {
	if tc == nil {
		return nil
	}
	res := response.TopCadets{
		Cadets: make([]response.CadetData, len(tc.Cadets)),
	}

	for i, v := range tc.Cadets {
		res.Cadets[i] = response.CadetData{
			Personal: response.Personal{
				PersonalDTO: v.Personal.PersonalDTO,
				Admission:   v.Personal.Admission.Name,
				HoursTotal:  v.Hours.Total,
				Color:       v.Personal.Admission.Color,
			},
			Journey: v.Journey,
		}
	}

	return &res
}
