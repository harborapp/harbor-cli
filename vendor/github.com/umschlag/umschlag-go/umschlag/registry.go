package umschlag

import (
	"time"
)

// Registry represents a registry API response.
type Registry struct {
	ID        int64     `json:"id"`
	Slug      string    `json:"slug"`
	Name      string    `json:"name"`
	Host      string    `json:"host"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Orgs      []*Org    `json:"orgs,omitempty"`
}

func (s *Registry) String() string {
	return s.Name
}
