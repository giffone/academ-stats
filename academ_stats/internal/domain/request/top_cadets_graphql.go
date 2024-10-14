// structs use for parsing requests to other services
package request

import (
	"time"
)

type TopCadets struct {
	Cadets []Cadet `json:"cadets"`
}

type Cadet struct {
	UserID int    `json:"userId"`
	Login  string `json:"login"`
	Level  int    `json:"level"`
	User   struct {
		Attrs struct {
			Gender      string    `json:"gender,omitempty"`
			LastName    string    `json:"lastName,omitempty"`
			FirstName   string    `json:"firstName,omitempty"`
			DateOfBirth time.Time `json:"dateOfBirth,omitempty"`
		} `json:"attrs"`
		Registrations []struct {
			Admission int `json:"admission,omitempty"`
		} `json:"registrations,omitempty"`
	} `json:"user"`
	Journey map[string]Object `json:"journey"` // key - {module, checkpoint, piscine}
}

type Object struct {
	XPTotal struct {
		Sum struct {
			Total float32 `json:"total"`
		} `json:"sum"`
	} `json:"xpTotal"`
	Paths []Path `json:"paths"`
}

type Path struct {
	XP   float32 `json:"xp"`
	Path struct {
		Attrs struct {
			ID       int    `json:"id"`
			PathName string `json:"pathName"`
			Language string `json:"language"`
		} `json:"attrs"`
	} `json:"path"`
	Event struct {
		RegistrationID int       `json:"registrationId"`
		CreatedAt      time.Time `json:"createdAt"`
		Object         struct {
			EventName string `json:"eventName"`
		} `json:"object"`
	} `json:"event"`
}
