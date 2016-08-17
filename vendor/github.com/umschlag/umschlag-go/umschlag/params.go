package umschlag

// TeamUserParams is used to assign users to a team.
type TeamUserParams struct {
	Team string `json:"team"`
	User string `json:"user"`
}

// UserTeamParams is used to assign teams to a user.
type UserTeamParams struct {
	User string `json:"user"`
	Team string `json:"team"`
}

// NamespaceUserParams is used to assign users to a namespace.
type NamespaceUserParams struct {
	Namespace string `json:"namespace"`
	User      string `json:"user"`
}

// UserNamespaceParams is used to assign namespaces to a user.
type UserNamespaceParams struct {
	User      string `json:"user"`
	Namespace string `json:"namespace"`
}

// NamespaceTeamParams is used to assign teams to a namespace.
type NamespaceTeamParams struct {
	Namespace string `json:"namespace"`
	Team      string `json:"team"`
}

// TeamNamespaceParams is used to assign namespaces to a team.
type TeamNamespaceParams struct {
	Team      string `json:"team"`
	Namespace string `json:"namespace"`
}
