package umschlag

// TeamOrg represents a team org API response.
type TeamOrg struct {
	Team *Team  `json:"team,omitempty"`
	Org  *Org   `json:"org,omitempty"`
	Perm string `json:"perm,omitempty"`
}
