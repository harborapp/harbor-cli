package umschlag

import (
	"time"
)

// Team represents a team API response.
type Team struct {
	ID        int64     `json:"id"`
	Slug      string    `json:"slug"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Users     []*User   `json:"users,omitempty"`
	Orgs      []*Org    `json:"orgs,omitempty"`
}

func (s *Team) String() string {
	return s.Name
}
