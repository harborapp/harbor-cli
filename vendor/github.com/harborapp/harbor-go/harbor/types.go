package harbor

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
	Expite string `json:"expire,omitempty"`
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
}

func (s *User) String() string {
	return s.Username
}

// Team represents a user API response.
type Team struct {
	ID        int64     `json:"id"`
	Slug      string    `json:"slug"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Users     []*User   `json:"users,omitempty"`
}

func (s *Team) String() string {
	return s.Name
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
