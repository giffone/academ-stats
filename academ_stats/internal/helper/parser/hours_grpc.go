package parser

import (
	"academ_stats/internal/domain"
	"academ_stats/internal/repository/pb/session_manager"
)

// HoursGrpc reads response from grpc request and copy data to internal struct Hours
func HoursGrpc(req *session_manager.CadetsTimeResponse) map[int]domain.HoursDTO {
	if req == nil {
		return nil
	}

	res := make(map[int]domain.HoursDTO)

	for _, c := range req.Cadets {
		id := int(c.GetId())

		if id == 0 || c.GetTotal() == 0 {
			continue
		}

		lm := len(c.GetMonth())
		ms := make([]domain.MonthDTO, lm)

		if lm > 0 {
			for i, m := range c.GetMonth() {
				ms[i] = domain.MonthDTO{
					Year:  m.GetYear(),
					Month: m.GetMonth().String(),
					Hours: int(m.GetHours()),
				}
			}
		}

		res[id] = domain.HoursDTO{
			Total: int(c.GetTotal()),
			Month: ms,
		}
	}

	return res
}
