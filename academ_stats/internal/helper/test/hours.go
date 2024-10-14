package test

import (
	"academ_stats/internal/domain"
	"academ_stats/internal/repository/pb/session_manager"
	"time"
)

const (
	UserAabdikha = 870
	UserAolzhass = 634
)

var Period = domain.Period{
	FromDate: time.Date(2024, time.January, 1, 0, 0, 0, 0, time.Local),
	ToDate:   time.Date(2024, time.June, 0, 0, 0, 0, 0, time.Local),
}

var ReqGrpcHours = session_manager.CadetsTimeResponse{
	Cadets: []*session_manager.Cadet{
		{
			Id:    UserAabdikha,
			Total: 100,
			Month: []*session_manager.MonthNum{
				{
					Year:  "2024",
					Month: session_manager.MonthNumber_january,
					Hours: 50,
				},
				{
					Year:  "2024",
					Month: session_manager.MonthNumber_february,
					Hours: 50,
				},
			},
		},
		{
			Id:    UserAolzhass,
			Total: 200,
			Month: []*session_manager.MonthNum{
				{
					Year:  "2024",
					Month: session_manager.MonthNumber_january,
					Hours: 50,
				},
				{
					Year:  "2024",
					Month: session_manager.MonthNumber_february,
					Hours: 50,
				},
				{
					Year:  "2024",
					Month: session_manager.MonthNumber_march,
					Hours: 50,
				},
				{
					Year:  "2024",
					Month: session_manager.MonthNumber_may,
					Hours: 50,
				},
			},
		},
	},
}
