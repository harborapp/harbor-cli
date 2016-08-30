package umschlag

import (
	"time"
)

// Repo represents a repo API response.
type Repo struct {
	ID        int64     `json:"id"`
	Org       *Org      `json:"org,omitempty"`
	OrgID     int64     `json:"org_id"`
	Slug      string    `json:"slug"`
	Name      string    `json:"name"`
	FullName  string    `json:"full_name"`
	Public    bool      `json:"public"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Tags      []*Tag    `json:"tags,omitempty"`
}

func (s *Repo) String() string {
	return s.Name
}
