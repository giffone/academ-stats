package domain

import (
	"fmt"
	"time"
)

type Period struct {
	FromDate time.Time
	ToDate   time.Time
}

func(p Period) String() string {
	return fmt.Sprintf("%s - %s",p.FromDate.Format("Jan06"), p.ToDate.Format("Jan06"))
}