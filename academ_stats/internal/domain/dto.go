// structs for internal use
package domain

import (
	"time"
)

// for general use
type (
	PersonalDTO struct {
		ID       int    `json:"id" db:"id"`               //
		Login    string `json:"login" db:"login"`         //
		FullName string `json:"full_name" db:"full_name"` //
		Age      int    `json:"age" db:"age"`             //
		Gender   string `json:"gender" db:"gender"`       //
		Level    int    `json:"level" db:"level"`         //
		XPTotal  int    `json:"xp_total" db:"xp_total"`   // calc xp by all events
	}

	AdmissionDTO struct {
		Name  string `json:"name" db:"name"`   // example "piscine23sep"
		Color string `json:"color" db:"color"` // example #000000
	}

	HoursDTO struct {
		Total int        `json:"total"` //
		Month []MonthDTO `json:"month"` //
	}

	MonthDTO struct {
		Year  string `json:"year"`  //
		Month string `json:"month"` //
		Hours int    `json:"hours"` //
	}

	EventDTO struct {
		ID        int           `json:"id"`         //
		Name      string        `json:"name"`       //
		XPTotal   int           `json:"xp_total"`   // calc xp by all tasks in event
		CreatedAt time.Time     `json:"created_at"` //
		Languages []LanguageDTO `json:"languages"`  // will sort by language
	}

	LanguageDTO struct {
		Name  string    `json:"name"`  //
		Tasks []TaskDTO `json:"tasks"` // slice sorted by "XP" asc
	}

	TaskDTO struct {
		ID   int    `json:"id"`   //
		Name string `json:"name"` //
		XP   int    `json:"xp"`   // xp by each task
	}
)

// Journey = [Module, Checkpoint, Piscine][]Event

// if Module = usually 1 Event
// [
//     Event{ID: 66, Name: "Module", ...}
// ]

// if Checkpoint = n-Events
// [
//     Event{ID: 76, Name: "Checkpoint", ...},
//     Event{ID: 77, Name: "Checkpoint", ...}, ...
// ]

// Event = Module (n-lang)
// ID: 66
// Name: "Module"
// Languages: [
//     Language{
//         Name: "Go"
//         Tasks: [
//             Task{Name: go-reloaded, XP: 5000},
//             Task{Name: ascii-art, XP: 10000}, ...
//         ]
//     },
//     Language{
//         Name: "other"
//         Tasks: [
//             Task{Name: math-skills, XP: 6000}, ...
//         ]
//     },
// ]

// Event = Checkpoint (1 lang)
// ID: 76
// Name: "Checkpoint"
// Languages:[
//     Language{
//         Name: "go"
//         Tasks: [
//             Task{Name: doop, XP: 300},
//             Task{Name: atoiBase, XP: 400}, ...
//         ]
//     }
// ]
