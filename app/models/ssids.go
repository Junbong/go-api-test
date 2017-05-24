package models

import (
	"time"
	"fmt"
)

type Ssids struct {
	Id          int
	Name        string
	CountryCode string
	SsidType    string
	Regex       bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (s *Ssids) String() string {
	return fmt.Sprintf(
		"ssids{id=%d, name=%s, country_code=%s, ssid_tye=%s, regex=%t, created_at=%s, updated_at=%s}",
		s.Id, s.Name, s.CountryCode, s.SsidType, s.Regex, s.CreatedAt, s.UpdatedAt)
}
