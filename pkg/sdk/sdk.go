package sdk

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"

	"github.com/jackspirou/syscerts"
	"golang.org/x/oauth2"
)

const (
	pathAuthLogin      = "%s/api/auth/login"
	pathProfile        = "%s/api/profile/self"
	pathProfileToken   = "%s/api/profile/token"
	pathRegistries     = "%s/api/registries"
	pathRegistry       = "%s/api/registries/%v"
	pathRegistryMember = "%s/api/registries/%v/%v"
	pathTags           = "%s/api/tags"
	pathTag            = "%s/api/tags/%v"
	pathRepos          = "%s/api/repos"
	pathRepo           = "%s/api/repos/%v"
	pathOrgs           = "%s/api/orgs"
	pathOrg            = "%s/api/orgs/%v"
	pathOrgUser        = "%s/api/orgs/%v/users"
	pathOrgTeam        = "%s/api/orgs/%v/teams"
	pathUsers          = "%s/api/users"
	pathUser           = "%s/api/users/%v"
	pathUserTeam       = "%s/api/users/%v/teams"
	pathUserOrg        = "%s/api/users/%v/orgs"
	pathTeams          = "%s/api/teams"
	pathTeam           = "%s/api/teams/%v"
	pathTeamUser       = "%s/api/teams/%v/users"
	pathTeamOrg        = "%s/api/teams/%v/orgs"
)

// ClientAPI describes a client API.
type ClientAPI interface {
	// SetClient sets the default http client. This should
	// be used in conjunction with golang.org/x/oauth2 to
	// authenticate requests to the API.
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

	// RegistrySync synchronizes a registry.
	RegistrySync(string) error

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
	OrgUserList(OrgUserParams) ([]*UserOrg, error)

	// OrgUserAppend appends a user to a org.
	OrgUserAppend(OrgUserParams) error

	// OrgUserPerm updates perms for org user.
	OrgUserPerm(OrgUserParams) error

	// OrgUserDelete remove a user from a org.
	OrgUserDelete(OrgUserParams) error

	// OrgTeamList returns a list of related teams for a org.
	OrgTeamList(OrgTeamParams) ([]*TeamOrg, error)

	// OrgTeamAppend appends a team to a org.
	OrgTeamAppend(OrgTeamParams) error

	// OrgTeamPerm updates perms for org team.
	OrgTeamPerm(OrgTeamParams) error

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
	UserTeamList(UserTeamParams) ([]*TeamUser, error)

	// UserTeamAppend appends a team to a user.
	UserTeamAppend(UserTeamParams) error

	// UserTeamPerm updates perms for user team.
	UserTeamPerm(UserTeamParams) error

	// UserTeamDelete remove a team from a user.
	UserTeamDelete(UserTeamParams) error

	// UserOrgList returns a list of related orgs for a user.
	UserOrgList(UserOrgParams) ([]*UserOrg, error)

	// UserOrgAppend appends a org to a user.
	UserOrgAppend(UserOrgParams) error

	// UserOrgPerm updates perms for user org.
	UserOrgPerm(UserOrgParams) error

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
	TeamUserList(TeamUserParams) ([]*TeamUser, error)

	// TeamUserAppend appends a user to a team.
	TeamUserAppend(TeamUserParams) error

	// TeamUserPerm updates perms for team user.
	TeamUserPerm(TeamUserParams) error

	// TeamUserDelete remove a user from a team.
	TeamUserDelete(TeamUserParams) error

	// TeamOrgList returns a list of related orgs for a team.
	TeamOrgList(TeamOrgParams) ([]*TeamOrg, error)

	// TeamOrgAppend appends a org to a team.
	TeamOrgAppend(TeamOrgParams) error

	// TeamOrgPerm updates perms for team org.
	TeamOrgPerm(TeamOrgParams) error

	// TeamOrgDelete remove a org from a team.
	TeamOrgDelete(TeamOrgParams) error
}

// Default implements the client interface.
type Default struct {
	client *http.Client
	base   string
	token  string
}

// NewClient returns a client for the specified URL.
func NewClient(uri string) ClientAPI {
	return &Default{
		client: http.DefaultClient,
		base:   uri,
	}
}

// NewClientToken returns a client that authenticates
// all outbound requests with the given token.
func NewClientToken(uri, token string) ClientAPI {
	config := oauth2.Config{}

	client := config.Client(
		context.Background(),
		&oauth2.Token{
			AccessToken: token,
		},
	)

	if trans, ok := client.Transport.(*oauth2.Transport); ok {
		trans.Base = &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			TLSClientConfig: &tls.Config{
				RootCAs: syscerts.SystemRootsPool(),
			},
		}
	}

	return &Default{
		client: client,
		base:   uri,
		token:  token,
	}
}

// IsAuthenticated checks if we already provided an authentication
// token for our client requests. If it returns false you can update
// the client after fetching a valid token.
func (c *Default) IsAuthenticated() bool {
	if c.token == "" {
		return false
	}

	uri, err := url.Parse(fmt.Sprintf(pathProfileToken, c.base))

	if err != nil {
		return false
	}

	req, err := http.NewRequest("GET", uri.String(), nil)

	if err != nil {
		return false
	}

	req.Header.Set(
		"User-Agent",
		UserAgent,
	)

	resp, err := c.client.Do(req)

	if err != nil {
		return false
	}

	defer resp.Body.Close()

	return resp.StatusCode == http.StatusUnauthorized
}

// SetClient sets the default http client. This should
// be used in conjunction with golang.org/x/oauth2 to
// authenticate requests to the API.
func (c *Default) SetClient(client *http.Client) {
	c.client = client
}

// AuthLogin signs in based on credentials and returns a token.
func (c *Default) AuthLogin(username, password string) (*Token, error) {
	out := &Token{}

	in := struct {
		Username string
		Password string
	}{
		username,
		password,
	}

	uri := fmt.Sprintf(pathAuthLogin, c.base)
	err := c.post(uri, in, out)

	return out, err
}

// ProfileToken returns a profile.
func (c *Default) ProfileToken() (*Token, error) {
	out := &Token{}

	uri := fmt.Sprintf(pathProfileToken, c.base)
	err := c.get(uri, out)

	return out, err
}

// ProfileGet returns a profile.
func (c *Default) ProfileGet() (*Profile, error) {
	out := &Profile{}

	uri := fmt.Sprintf(pathProfile, c.base)
	err := c.get(uri, out)

	return out, err
}

// ProfilePatch updates a profile.
func (c *Default) ProfilePatch(in *Profile) (*Profile, error) {
	out := &Profile{}

	uri := fmt.Sprintf(pathProfile, c.base)
	err := c.put(uri, in, out)

	return out, err
}

// RegistryList returns a list of all registries.
func (c *Default) RegistryList() ([]*Registry, error) {
	var out []*Registry

	uri := fmt.Sprintf(pathRegistries, c.base)
	err := c.get(uri, &out)

	return out, err
}

// RegistryGet returns a registry.
func (c *Default) RegistryGet(id string) (*Registry, error) {
	out := &Registry{}

	uri := fmt.Sprintf(pathRegistry, c.base, id)
	err := c.get(uri, out)

	return out, err
}

// RegistryPost creates a registry.
func (c *Default) RegistryPost(in *Registry) (*Registry, error) {
	out := &Registry{}

	uri := fmt.Sprintf(pathRegistries, c.base)
	err := c.post(uri, in, out)

	return out, err
}

// RegistryPatch updates a registry.
func (c *Default) RegistryPatch(in *Registry) (*Registry, error) {
	out := &Registry{}

	uri := fmt.Sprintf(pathRegistry, c.base, in.ID)
	err := c.put(uri, in, out)

	return out, err
}

// RegistryDelete deletes a registry.
func (c *Default) RegistryDelete(id string) error {
	uri := fmt.Sprintf(pathRegistry, c.base, id)
	err := c.delete(uri, nil)

	return err
}

// RegistrySync synchronizes a registry.
func (c *Default) RegistrySync(id string) error {
	uri := fmt.Sprintf(pathRegistryMember, c.base, id, "sync")
	err := c.post(uri, nil, nil)

	return err
}

// TagList returns a list of all tags.
func (c *Default) TagList() ([]*Tag, error) {
	var out []*Tag

	uri := fmt.Sprintf(pathTags, c.base)
	err := c.get(uri, &out)

	return out, err
}

// TagGet returns a tag.
func (c *Default) TagGet(id string) (*Tag, error) {
	out := &Tag{}

	uri := fmt.Sprintf(pathTag, c.base, id)
	err := c.get(uri, out)

	return out, err
}

// TagDelete deletes a tag.
func (c *Default) TagDelete(id string) error {
	uri := fmt.Sprintf(pathTag, c.base, id)
	err := c.delete(uri, nil)

	return err
}

// RepoList returns a list of all repos.
func (c *Default) RepoList() ([]*Repo, error) {
	var out []*Repo

	uri := fmt.Sprintf(pathRepos, c.base)
	err := c.get(uri, &out)

	return out, err
}

// RepoGet returns a repo.
func (c *Default) RepoGet(id string) (*Repo, error) {
	out := &Repo{}

	uri := fmt.Sprintf(pathRepo, c.base, id)
	err := c.get(uri, out)

	return out, err
}

// RepoDelete deletes a repo.
func (c *Default) RepoDelete(id string) error {
	uri := fmt.Sprintf(pathRepo, c.base, id)
	err := c.delete(uri, nil)

	return err
}

// OrgList returns a list of all orgs.
func (c *Default) OrgList() ([]*Org, error) {
	var out []*Org

	uri := fmt.Sprintf(pathOrgs, c.base)
	err := c.get(uri, &out)

	return out, err
}

// OrgGet returns a org.
func (c *Default) OrgGet(id string) (*Org, error) {
	out := &Org{}

	uri := fmt.Sprintf(pathOrg, c.base, id)
	err := c.get(uri, out)

	return out, err
}

// OrgPost creates a org.
func (c *Default) OrgPost(in *Org) (*Org, error) {
	out := &Org{}

	uri := fmt.Sprintf(pathOrgs, c.base)
	err := c.post(uri, in, out)

	return out, err
}

// OrgPatch updates a org.
func (c *Default) OrgPatch(in *Org) (*Org, error) {
	out := &Org{}

	uri := fmt.Sprintf(pathOrg, c.base, in.ID)
	err := c.put(uri, in, out)

	return out, err
}

// OrgDelete deletes a org.
func (c *Default) OrgDelete(id string) error {
	uri := fmt.Sprintf(pathOrg, c.base, id)
	err := c.delete(uri, nil)

	return err
}

// OrgUserList returns a list of related users for a org.
func (c *Default) OrgUserList(opts OrgUserParams) ([]*UserOrg, error) {
	var out []*UserOrg

	uri := fmt.Sprintf(pathOrgUser, c.base, opts.Org)
	err := c.get(uri, &out)

	return out, err
}

// OrgUserAppend appends a user to a org.
func (c *Default) OrgUserAppend(opts OrgUserParams) error {
	uri := fmt.Sprintf(pathOrgUser, c.base, opts.Org)
	err := c.post(uri, opts, nil)

	return err
}

// OrgUserPerm updates perms for org user.
func (c *Default) OrgUserPerm(opts OrgUserParams) error {
	uri := fmt.Sprintf(pathOrgUser, c.base, opts.Org)
	err := c.put(uri, opts, nil)

	return err
}

// OrgUserDelete remove a user from a org.
func (c *Default) OrgUserDelete(opts OrgUserParams) error {
	uri := fmt.Sprintf(pathOrgUser, c.base, opts.Org)
	err := c.delete(uri, opts)

	return err
}

// OrgTeamList returns a list of related teams for a org.
func (c *Default) OrgTeamList(opts OrgTeamParams) ([]*TeamOrg, error) {
	var out []*TeamOrg

	uri := fmt.Sprintf(pathOrgTeam, c.base, opts.Org)
	err := c.get(uri, &out)

	return out, err
}

// OrgTeamAppend appends a team to a org.
func (c *Default) OrgTeamAppend(opts OrgTeamParams) error {
	uri := fmt.Sprintf(pathOrgTeam, c.base, opts.Org)
	err := c.post(uri, opts, nil)

	return err
}

// OrgTeamPerm updates perms for org team.
func (c *Default) OrgTeamPerm(opts OrgTeamParams) error {
	uri := fmt.Sprintf(pathOrgTeam, c.base, opts.Org)
	err := c.put(uri, opts, nil)

	return err
}

// OrgTeamDelete remove a team from a org.
func (c *Default) OrgTeamDelete(opts OrgTeamParams) error {
	uri := fmt.Sprintf(pathOrgTeam, c.base, opts.Org)
	err := c.delete(uri, opts)

	return err
}

// UserList returns a list of all users.
func (c *Default) UserList() ([]*User, error) {
	var out []*User

	uri := fmt.Sprintf(pathUsers, c.base)
	err := c.get(uri, &out)

	return out, err
}

// UserGet returns a user.
func (c *Default) UserGet(id string) (*User, error) {
	out := &User{}

	uri := fmt.Sprintf(pathUser, c.base, id)
	err := c.get(uri, out)

	return out, err
}

// UserPost creates a user.
func (c *Default) UserPost(in *User) (*User, error) {
	out := &User{}

	uri := fmt.Sprintf(pathUsers, c.base)
	err := c.post(uri, in, out)

	return out, err
}

// UserPatch updates a user.
func (c *Default) UserPatch(in *User) (*User, error) {
	out := &User{}

	uri := fmt.Sprintf(pathUser, c.base, in.ID)
	err := c.put(uri, in, out)

	return out, err
}

// UserDelete deletes a user.
func (c *Default) UserDelete(id string) error {
	uri := fmt.Sprintf(pathUser, c.base, id)
	err := c.delete(uri, nil)

	return err
}

// UserTeamList returns a list of related teams for a user.
func (c *Default) UserTeamList(opts UserTeamParams) ([]*TeamUser, error) {
	var out []*TeamUser

	uri := fmt.Sprintf(pathUserTeam, c.base, opts.User)
	err := c.get(uri, &out)

	return out, err
}

// UserTeamAppend appends a team to a user.
func (c *Default) UserTeamAppend(opts UserTeamParams) error {
	uri := fmt.Sprintf(pathUserTeam, c.base, opts.User)
	err := c.post(uri, opts, nil)

	return err
}

// UserTeamPerm updates perms for user team.
func (c *Default) UserTeamPerm(opts UserTeamParams) error {
	uri := fmt.Sprintf(pathUserTeam, c.base, opts.User)
	err := c.put(uri, opts, nil)

	return err
}

// UserTeamDelete remove a team from a user.
func (c *Default) UserTeamDelete(opts UserTeamParams) error {
	uri := fmt.Sprintf(pathUserTeam, c.base, opts.User)
	err := c.delete(uri, opts)

	return err
}

// UserOrgList returns a list of related orgs for a user.
func (c *Default) UserOrgList(opts UserOrgParams) ([]*UserOrg, error) {
	var out []*UserOrg

	uri := fmt.Sprintf(pathUserOrg, c.base, opts.User)
	err := c.get(uri, &out)

	return out, err
}

// UserOrgAppend appends a org to a user.
func (c *Default) UserOrgAppend(opts UserOrgParams) error {
	uri := fmt.Sprintf(pathUserOrg, c.base, opts.User)
	err := c.post(uri, opts, nil)

	return err
}

// UserOrgPerm updates perms for user org.
func (c *Default) UserOrgPerm(opts UserOrgParams) error {
	uri := fmt.Sprintf(pathUserOrg, c.base, opts.User)
	err := c.put(uri, opts, nil)

	return err
}

// UserOrgDelete remove a org from a user.
func (c *Default) UserOrgDelete(opts UserOrgParams) error {
	uri := fmt.Sprintf(pathUserOrg, c.base, opts.User)
	err := c.delete(uri, opts)

	return err
}

// TeamList returns a list of all teams.
func (c *Default) TeamList() ([]*Team, error) {
	var out []*Team

	uri := fmt.Sprintf(pathTeams, c.base)
	err := c.get(uri, &out)

	return out, err
}

// TeamGet returns a team.
func (c *Default) TeamGet(id string) (*Team, error) {
	out := &Team{}

	uri := fmt.Sprintf(pathTeam, c.base, id)
	err := c.get(uri, out)

	return out, err
}

// TeamPost creates a team.
func (c *Default) TeamPost(in *Team) (*Team, error) {
	out := &Team{}

	uri := fmt.Sprintf(pathTeams, c.base)
	err := c.post(uri, in, out)

	return out, err
}

// TeamPatch updates a team.
func (c *Default) TeamPatch(in *Team) (*Team, error) {
	out := &Team{}

	uri := fmt.Sprintf(pathTeam, c.base, in.ID)
	err := c.put(uri, in, out)

	return out, err
}

// TeamDelete deletes a team.
func (c *Default) TeamDelete(id string) error {
	uri := fmt.Sprintf(pathTeam, c.base, id)
	err := c.delete(uri, nil)

	return err
}

// TeamUserList returns a list of related users for a team.
func (c *Default) TeamUserList(opts TeamUserParams) ([]*TeamUser, error) {
	var out []*TeamUser

	uri := fmt.Sprintf(pathTeamUser, c.base, opts.Team)
	err := c.get(uri, &out)

	return out, err
}

// TeamUserAppend appends a user to a team.
func (c *Default) TeamUserAppend(opts TeamUserParams) error {
	uri := fmt.Sprintf(pathTeamUser, c.base, opts.Team)
	err := c.post(uri, opts, nil)

	return err
}

// TeamUserPerm updates perms for team user.
func (c *Default) TeamUserPerm(opts TeamUserParams) error {
	uri := fmt.Sprintf(pathTeamUser, c.base, opts.Team)
	err := c.put(uri, opts, nil)

	return err
}

// TeamUserDelete remove a user from a team.
func (c *Default) TeamUserDelete(opts TeamUserParams) error {
	uri := fmt.Sprintf(pathTeamUser, c.base, opts.Team)
	err := c.delete(uri, opts)

	return err
}

// TeamOrgList returns a list of related orgs for a team.
func (c *Default) TeamOrgList(opts TeamOrgParams) ([]*TeamOrg, error) {
	var out []*TeamOrg

	uri := fmt.Sprintf(pathTeamOrg, c.base, opts.Team)
	err := c.get(uri, &out)

	return out, err
}

// TeamOrgAppend appends a org to a team.
func (c *Default) TeamOrgAppend(opts TeamOrgParams) error {
	uri := fmt.Sprintf(pathTeamOrg, c.base, opts.Team)
	err := c.post(uri, opts, nil)

	return err
}

// TeamOrgPerm updates perms for team org.
func (c *Default) TeamOrgPerm(opts TeamOrgParams) error {
	uri := fmt.Sprintf(pathTeamOrg, c.base, opts.Team)
	err := c.put(uri, opts, nil)

	return err
}

// TeamOrgDelete remove a org from a team.
func (c *Default) TeamOrgDelete(opts TeamOrgParams) error {
	uri := fmt.Sprintf(pathTeamOrg, c.base, opts.Team)
	err := c.delete(uri, opts)

	return err
}
