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

	// UserNamespaceList returns a list of related namespaces for a user.
	UserNamespaceList(UserNamespaceParams) ([]*Namespace, error)

	// UserNamespaceAppend appends a namespace to a user.
	UserNamespaceAppend(UserNamespaceParams) error

	// UserNamespaceDelete remove a namespace from a user.
	UserNamespaceDelete(UserNamespaceParams) error

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

	// TeamNamespaceList returns a list of related namespaces for a team.
	TeamNamespaceList(TeamNamespaceParams) ([]*Namespace, error)

	// TeamNamespaceAppend appends a namespace to a team.
	TeamNamespaceAppend(TeamNamespaceParams) error

	// TeamNamespaceDelete remove a namespace from a team.
	TeamNamespaceDelete(TeamNamespaceParams) error

	// NamespaceList returns a list of all namespaces.
	NamespaceList() ([]*Namespace, error)

	// NamespaceGet returns a namespace.
	NamespaceGet(string) (*Namespace, error)

	// NamespacePost creates a namespace.
	NamespacePost(*Namespace) (*Namespace, error)

	// NamespacePatch updates a namespace.
	NamespacePatch(*Namespace) (*Namespace, error)

	// NamespaceDelete deletes a namespace.
	NamespaceDelete(string) error

	// NamespaceUserList returns a list of related users for a namespace.
	NamespaceUserList(NamespaceUserParams) ([]*User, error)

	// NamespaceUserAppend appends a user to a namespace.
	NamespaceUserAppend(NamespaceUserParams) error

	// NamespaceUserDelete remove a user from a namespace.
	NamespaceUserDelete(NamespaceUserParams) error

	// NamespaceTeamList returns a list of related teams for a namespace.
	NamespaceTeamList(NamespaceTeamParams) ([]*Team, error)

	// NamespaceTeamAppend appends a team to a namespace.
	NamespaceTeamAppend(NamespaceTeamParams) error

	// NamespaceTeamDelete remove a team from a namespace.
	NamespaceTeamDelete(NamespaceTeamParams) error
}
