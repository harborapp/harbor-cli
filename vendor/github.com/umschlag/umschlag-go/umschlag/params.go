package umschlag

// TeamUserParams is used to assign users to a team.
type TeamUserParams struct {
	Team string `json:"team"`
	User string `json:"user"`
	Perm string `json:"perm"`
}

// UserTeamParams is used to assign teams to a user.
type UserTeamParams struct {
	User string `json:"user"`
	Team string `json:"team"`
	Perm string `json:"perm"`
}

// OrgUserParams is used to assign users to a org.
type OrgUserParams struct {
	Org  string `json:"org"`
	User string `json:"user"`
	Perm string `json:"perm"`
}

// UserOrgParams is used to assign orgs to a user.
type UserOrgParams struct {
	User string `json:"user"`
	Org  string `json:"org"`
	Perm string `json:"perm"`
}

// OrgTeamParams is used to assign teams to a org.
type OrgTeamParams struct {
	Org  string `json:"org"`
	Team string `json:"team"`
	Perm string `json:"perm"`
}

// TeamOrgParams is used to assign orgs to a team.
type TeamOrgParams struct {
	Team string `json:"team"`
	Org  string `json:"org"`
	Perm string `json:"perm"`
}
