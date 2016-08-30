package umschlag

import (
	"time"
)

// Org represents a org API response.
type Org struct {
	ID         int64     `json:"id"`
	Registry   *Registry `json:"registry,omitempty"`
	RegistryID int64     `json:"registry_id"`
	Slug       string    `json:"slug"`
	Name       string    `json:"name"`
	Public     bool      `json:"public"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Repos      []*Repo   `json:"repos,omitempty"`
	Users      []*User   `json:"users,omitempty"`
	Teams      []*Team   `json:"teams,omitempty"`
}

func (s *Org) String() string {
	return s.Name
}
