package parser

import (
	"academ_stats/internal/domain"
	"academ_stats/internal/domain/request"
	"fmt"
	"time"
)

// PersonalGraphql reads data from graphql request and copy data to internal struct Personal
func PersonalGraphql(c *request.Cadet, adm map[int]domain.Piscine) domain.Personal {
	cd := domain.Personal{
		PersonalDTO: domain.PersonalDTO{
			ID:       c.UserID,
			Login:    c.Login,
			FullName: fmt.Sprintf("%s %s", c.User.Attrs.FirstName, c.User.Attrs.LastName),
			Age:      int(time.Since(c.User.Attrs.DateOfBirth).Hours() / 24 / 365),
			Gender:   c.User.Attrs.Gender,
			Level:    c.Level,
		},
	}

	for _, v := range c.Journey {
		cd.XPTotal += int(v.XPTotal.Sum.Total)
	}

	if adm == nil {
		return cd
	}

	// cadet admission [when he entered basic training]
	for _, v := range c.User.Registrations {
		if a, ok := adm[v.Admission]; ok {
			if cd.Admission.Name != "" {
				cd.Admission.Name = fmt.Sprintf("%s / %s", cd.Admission.Name, a.Title)
			} else {
				cd.Admission.Name = a.Title
			}
			cd.Admission.Color = a.Color
		}
	}

	return cd
}
