package umschlag

// UserOrg represents a user org API response.
type UserOrg struct {
	User *User  `json:"user,omitempty"`
	Org  *Org   `json:"org,omitempty"`
	Perm string `json:"perm,omitempty"`
}
