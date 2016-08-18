package umschlag

import (
	"time"
)

// Message represents a standard response.
type Message struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Token represents a session token.
type Token struct {
	Token  string `json:"token"`
	Expire string `json:"expire,omitempty"`
}

// Profile represents a profile API response.
type Profile struct {
	ID        int64     `json:"id"`
	Slug      string    `json:"slug"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *Profile) String() string {
	return s.Username
}

// Registry represents a registry API response.
type Registry struct {
	ID        int64     `json:"id"`
	Slug      string    `json:"slug"`
	Name      string    `json:"name"`
	Host      string    `json:"host"`
	UseSSL    bool      `json:"use_ssl"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Orgs      []*Org    `json:"orgs,omitempty"`
}

func (s *Registry) String() string {
	return s.Name
}

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

// User represents a user API response.
type User struct {
	ID        int64     `json:"id"`
	Slug      string    `json:"slug"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Teams     []*Team   `json:"teams,omitempty"`
	Orgs      []*Org    `json:"orgs,omitempty"`
}

func (s *User) String() string {
	return s.Username
}

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
