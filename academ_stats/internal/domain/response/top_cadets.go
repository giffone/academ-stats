// structs to use when sending responses to requests from other services
package response

import (
	"academ_stats/internal/domain"
	"encoding/json"
	"time"
)

type (
	TopCadetsResponse struct {
		Current  json.RawMessage `json:"current,omitempty"`
		Last     json.RawMessage `json:"last,omitempty"`
		LastDate time.Time       `json:"last_date,omitempty"`
	}
)

type (
	TopCadets struct {
		Cadets []CadetData `json:"cadets"` // sorted by "Level" desc, "Login" asc
	}

	CadetData struct {
		Personal Personal                     `json:"personal"`
		Journey  map[string][]domain.EventDTO `json:"journey"` // key - {Module, Checkpoint, Piscine}
	}

	Personal struct {
		domain.PersonalDTO
		Admission  string `json:"admission"`   //
		HoursTotal int    `json:"hours_total"` //
		Color      string `json:"color"`       //
	}
)
