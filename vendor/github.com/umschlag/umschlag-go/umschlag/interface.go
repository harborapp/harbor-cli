package umschlag

import (
	"net/http"
)

//go:generate mockery -all -case=underscore

// ClientAPI describes a client API.
type ClientAPI interface {
	// SetClient sets the default http client. This should
	// be used in conjunction with golang.org/x/oauth2 to
	// authenticate requests to the Umschlag API.
	SetClient(client *http.Client)

	// IsAuthenticated checks if we already provided an authentication
	// token for our client requests. If it returns false you can update
	// the client after fetching a valid token.
	IsAuthenticated() bool

	// AuthLogin signs in based on credentials and returns a token.
	AuthLogin(string, string) (*Token, error)

	// ProfileToken returns a token.
	ProfileToken() (*Token, error)

	// ProfileGet returns a profile.
	ProfileGet() (*Profile, error)

	// ProfilePatch updates a profile.
	ProfilePatch(*Profile) (*Profile, error)

	// RegistryList returns a list of all registries.
	RegistryList() ([]*Registry, error)

	// RegistryGet returns a registry.
	RegistryGet(string) (*Registry, error)

	// RegistryPost creates a registry.
	RegistryPost(*Registry) (*Registry, error)

	// RegistryPatch updates a registry.
	RegistryPatch(*Registry) (*Registry, error)

	// RegistryDelete deletes a registry.
	RegistryDelete(string) error

	// TagList returns a list of all tags.
	TagList() ([]*Tag, error)

	// TagGet returns a tag.
	TagGet(string) (*Tag, error)

	// TagDelete deletes a tag.
	TagDelete(string) error

	// RepoList returns a list of all repos.
	RepoList() ([]*Repo, error)

	// RepoGet returns a repo.
	RepoGet(string) (*Repo, error)

	// RepoDelete deletes a repo.
	RepoDelete(string) error

	// OrgList returns a list of all orgs.
	OrgList() ([]*Org, error)

	// OrgGet returns a org.
	OrgGet(string) (*Org, error)

	// OrgPost creates a org.
	OrgPost(*Org) (*Org, error)

	// OrgPatch updates a org.
	OrgPatch(*Org) (*Org, error)

	// OrgDelete deletes a org.
	OrgDelete(string) error

	// OrgUserList returns a list of related users for a org.
	OrgUserList(OrgUserParams) ([]*User, error)

	// OrgUserAppend appends a user to a org.
	OrgUserAppend(OrgUserParams) error

	// OrgUserDelete remove a user from a org.
	OrgUserDelete(OrgUserParams) error

	// OrgTeamList returns a list of related teams for a org.
	OrgTeamList(OrgTeamParams) ([]*Team, error)

	// OrgTeamAppend appends a team to a org.
	OrgTeamAppend(OrgTeamParams) error

	// OrgTeamDelete remove a team from a org.
	OrgTeamDelete(OrgTeamParams) error

	// UserList returns a list of all users.
	UserList() ([]*User, error)

	// UserGet returns a user.
	UserGet(string) (*User, error)

	// UserPost creates a user.
	UserPost(*User) (*User, error)

	// UserPatch updates a user.
	UserPatch(*User) (*User, error)

	// UserDelete deletes a user.
	UserDelete(string) error

	// UserTeamList returns a list of related teams for a user.
	UserTeamList(UserTeamParams) ([]*Team, error)

	// UserTeamAppend appends a team to a user.
	UserTeamAppend(UserTeamParams) error

	// UserTeamDelete remove a team from a user.
	UserTeamDelete(UserTeamParams) error

	// UserOrgList returns a list of related orgs for a user.
	UserOrgList(UserOrgParams) ([]*Org, error)

	// UserOrgAppend appends a org to a user.
	UserOrgAppend(UserOrgParams) error

	// UserOrgDelete remove a org from a user.
	UserOrgDelete(UserOrgParams) error

	// TeamList returns a list of all teams.
	TeamList() ([]*Team, error)

	// TeamGet returns a team.
	TeamGet(string) (*Team, error)

	// TeamPost creates a team.
	TeamPost(*Team) (*Team, error)

	// TeamPatch updates a team.
	TeamPatch(*Team) (*Team, error)

	// TeamDelete deletes a team.
	TeamDelete(string) error

	// TeamUserList returns a list of related users for a team.
	TeamUserList(TeamUserParams) ([]*User, error)

	// TeamUserAppend appends a user to a team.
	TeamUserAppend(TeamUserParams) error

	// TeamUserDelete remove a user from a team.
	TeamUserDelete(TeamUserParams) error

	// TeamOrgList returns a list of related orgs for a team.
	TeamOrgList(TeamOrgParams) ([]*Org, error)

	// TeamOrgAppend appends a org to a team.
	TeamOrgAppend(TeamOrgParams) error

	// TeamOrgDelete remove a org from a team.
	TeamOrgDelete(TeamOrgParams) error
}
