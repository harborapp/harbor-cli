package umschlag

import (
	"time"
)

// Tag represents a tag API response.
type Tag struct {
	ID        int64     `json:"id"`
	Repo      *Repo     `json:"repo,omitempty"`
	RepoID    int64     `json:"repo_id"`
	Slug      string    `json:"slug"`
	Name      string    `json:"name"`
	FullName  string    `json:"full_name"`
	Public    bool      `json:"public"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *Tag) String() string {
	return s.Name
}
