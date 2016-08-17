package umschlag

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"

	"github.com/jackspirou/syscerts"
	"golang.org/x/oauth2"
)

const (
	pathAuthLogin     = "%s/api/auth/login"
	pathProfile       = "%s/api/profile/self"
	pathProfileToken  = "%s/api/profile/token"
	pathUsers         = "%s/api/users"
	pathUser          = "%s/api/users/%v"
	pathUserTeam      = "%s/api/users/%v/teams"
	pathUserNamespace = "%s/api/users/%v/namespaces"
	pathTeams         = "%s/api/teams"
	pathTeam          = "%s/api/teams/%v"
	pathTeamUser      = "%s/api/teams/%v/users"
	pathTeamNamespace = "%s/api/teams/%v/namespaces"
	pathNamespaces    = "%s/api/namespaces"
	pathNamespace     = "%s/api/namespaces/%v"
	pathNamespaceUser = "%s/api/namespaces/%v/users"
	pathNamespaceTeam = "%s/api/namespaces/%v/teams"
)

// DefaultClient implements the client interface.
type DefaultClient struct {
	client *http.Client
	base   string
	token  string
}

// NewClient returns a client for the specified URL.
func NewClient(uri string) ClientAPI {
	return &DefaultClient{
		client: http.DefaultClient,
		base:   uri,
	}
}

// NewClientToken returns a client that authenticates
// all outbound requests with the given token.
func NewClientToken(uri, token string) ClientAPI {
	config := oauth2.Config{}

	auther := config.Client(
		oauth2.NoContext,
		&oauth2.Token{
			AccessToken: token,
		},
	)

	if trans, ok := auther.Transport.(*oauth2.Transport); ok {
		trans.Base = &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			TLSClientConfig: &tls.Config{
				RootCAs: syscerts.SystemRootsPool(),
			},
		}
	}

	return &DefaultClient{
		client: auther,
		base:   uri,
		token:  token,
	}
}

// IsAuthenticated checks if we already provided an authentication
// token for our client requests. If it returns false you can update
// the client after fetching a valid token.
func (c *DefaultClient) IsAuthenticated() bool {
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
		"Umschlag CLI",
	)

	resp, err := c.client.Do(req)

	if err != nil {
		return false
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return false
	}

	return true
}

// SetClient sets the default http client. This should
// be used in conjunction with golang.org/x/oauth2 to
// authenticate requests to the Umschlag API.
func (c *DefaultClient) SetClient(client *http.Client) {
	c.client = client
}

// AuthLogin signs in based on credentials and returns a token.
func (c *DefaultClient) AuthLogin(username, password string) (*Token, error) {
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
func (c *DefaultClient) ProfileToken() (*Token, error) {
	out := &Token{}

	uri := fmt.Sprintf(pathProfileToken, c.base)
	err := c.get(uri, out)

	return out, err
}

// ProfileGet returns a profile.
func (c *DefaultClient) ProfileGet() (*Profile, error) {
	out := &Profile{}

	uri := fmt.Sprintf(pathProfile, c.base)
	err := c.get(uri, out)

	return out, err
}

// ProfilePatch updates a profile.
func (c *DefaultClient) ProfilePatch(in *Profile) (*Profile, error) {
	out := &Profile{}

	uri := fmt.Sprintf(pathProfile, c.base)
	err := c.patch(uri, in, out)

	return out, err
}

// UserList returns a list of all users.
func (c *DefaultClient) UserList() ([]*User, error) {
	var out []*User

	uri := fmt.Sprintf(pathUsers, c.base)
	err := c.get(uri, &out)

	return out, err
}

// UserGet returns a user.
func (c *DefaultClient) UserGet(id string) (*User, error) {
	out := &User{}

	uri := fmt.Sprintf(pathUser, c.base, id)
	err := c.get(uri, out)

	return out, err
}

// UserPost creates a user.
func (c *DefaultClient) UserPost(in *User) (*User, error) {
	out := &User{}

	uri := fmt.Sprintf(pathUsers, c.base)
	err := c.post(uri, in, out)

	return out, err
}

// UserPatch updates a user.
func (c *DefaultClient) UserPatch(in *User) (*User, error) {
	out := &User{}

	uri := fmt.Sprintf(pathUser, c.base, in.ID)
	err := c.patch(uri, in, out)

	return out, err
}

// UserDelete deletes a user.
func (c *DefaultClient) UserDelete(id string) error {
	uri := fmt.Sprintf(pathUser, c.base, id)
	err := c.delete(uri, nil)

	return err
}

// UserTeamList returns a list of related teams for a user.
func (c *DefaultClient) UserTeamList(opts UserTeamParams) ([]*Team, error) {
	var out []*Team

	uri := fmt.Sprintf(pathUserTeam, c.base, opts.User)
	err := c.get(uri, &out)

	return out, err
}

// UserTeamAppend appends a team to a user.
func (c *DefaultClient) UserTeamAppend(opts UserTeamParams) error {
	uri := fmt.Sprintf(pathUserTeam, c.base, opts.User)
	err := c.patch(uri, opts, nil)

	return err
}

// UserTeamDelete remove a team from a user.
func (c *DefaultClient) UserTeamDelete(opts UserTeamParams) error {
	uri := fmt.Sprintf(pathUserTeam, c.base, opts.User)
	err := c.delete(uri, opts)

	return err
}

// UserNamespaceList returns a list of related namespaces for a user.
func (c *DefaultClient) UserNamespaceList(opts UserNamespaceParams) ([]*Namespace, error) {
	var out []*Namespace

	uri := fmt.Sprintf(pathUserNamespace, c.base, opts.User)
	err := c.get(uri, &out)

	return out, err
}

// UserNamespaceAppend appends a namespace to a user.
func (c *DefaultClient) UserNamespaceAppend(opts UserNamespaceParams) error {
	uri := fmt.Sprintf(pathUserNamespace, c.base, opts.User)
	err := c.patch(uri, opts, nil)

	return err
}

// UserNamespaceDelete remove a namespace from a user.
func (c *DefaultClient) UserNamespaceDelete(opts UserNamespaceParams) error {
	uri := fmt.Sprintf(pathUserNamespace, c.base, opts.User)
	err := c.delete(uri, opts)

	return err
}

// TeamList returns a list of all teams.
func (c *DefaultClient) TeamList() ([]*Team, error) {
	var out []*Team

	uri := fmt.Sprintf(pathTeams, c.base)
	err := c.get(uri, &out)

	return out, err
}

// TeamGet returns a team.
func (c *DefaultClient) TeamGet(id string) (*Team, error) {
	out := &Team{}

	uri := fmt.Sprintf(pathTeam, c.base, id)
	err := c.get(uri, out)

	return out, err
}

// TeamPost creates a team.
func (c *DefaultClient) TeamPost(in *Team) (*Team, error) {
	out := &Team{}

	uri := fmt.Sprintf(pathTeams, c.base)
	err := c.post(uri, in, out)

	return out, err
}

// TeamPatch updates a team.
func (c *DefaultClient) TeamPatch(in *Team) (*Team, error) {
	out := &Team{}

	uri := fmt.Sprintf(pathTeam, c.base, in.ID)
	err := c.patch(uri, in, out)

	return out, err
}

// TeamDelete deletes a team.
func (c *DefaultClient) TeamDelete(id string) error {
	uri := fmt.Sprintf(pathTeam, c.base, id)
	err := c.delete(uri, nil)

	return err
}

// TeamUserList returns a list of related users for a team.
func (c *DefaultClient) TeamUserList(opts TeamUserParams) ([]*User, error) {
	var out []*User

	uri := fmt.Sprintf(pathTeamUser, c.base, opts.Team)
	err := c.get(uri, &out)

	return out, err
}

// TeamUserAppend appends a user to a team.
func (c *DefaultClient) TeamUserAppend(opts TeamUserParams) error {
	uri := fmt.Sprintf(pathTeamUser, c.base, opts.Team)
	err := c.patch(uri, opts, nil)

	return err
}

// TeamUserDelete remove a user from a team.
func (c *DefaultClient) TeamUserDelete(opts TeamUserParams) error {
	uri := fmt.Sprintf(pathTeamUser, c.base, opts.Team)
	err := c.delete(uri, opts)

	return err
}

// TeamNamespaceList returns a list of related namespaces for a team.
func (c *DefaultClient) TeamNamespaceList(opts TeamNamespaceParams) ([]*Namespace, error) {
	var out []*Namespace

	uri := fmt.Sprintf(pathTeamNamespace, c.base, opts.Team)
	err := c.get(uri, &out)

	return out, err
}

// TeamNamespaceAppend appends a namespace to a team.
func (c *DefaultClient) TeamNamespaceAppend(opts TeamNamespaceParams) error {
	uri := fmt.Sprintf(pathTeamNamespace, c.base, opts.Team)
	err := c.patch(uri, opts, nil)

	return err
}

// TeamNamespaceDelete remove a namespace from a team.
func (c *DefaultClient) TeamNamespaceDelete(opts TeamNamespaceParams) error {
	uri := fmt.Sprintf(pathTeamNamespace, c.base, opts.Team)
	err := c.delete(uri, opts)

	return err
}

// NamespaceList returns a list of all namespaces.
func (c *DefaultClient) NamespaceList() ([]*Namespace, error) {
	var out []*Namespace

	uri := fmt.Sprintf(pathNamespaces, c.base)
	err := c.get(uri, &out)

	return out, err
}

// NamespaceGet returns a namespace.
func (c *DefaultClient) NamespaceGet(id string) (*Namespace, error) {
	out := &Namespace{}

	uri := fmt.Sprintf(pathNamespace, c.base, id)
	err := c.get(uri, out)

	return out, err
}

// NamespacePost creates a namespace.
func (c *DefaultClient) NamespacePost(in *Namespace) (*Namespace, error) {
	out := &Namespace{}

	uri := fmt.Sprintf(pathNamespaces, c.base)
	err := c.post(uri, in, out)

	return out, err
}

// NamespacePatch updates a namespace.
func (c *DefaultClient) NamespacePatch(in *Namespace) (*Namespace, error) {
	out := &Namespace{}

	uri := fmt.Sprintf(pathNamespace, c.base, in.ID)
	err := c.patch(uri, in, out)

	return out, err
}

// NamespaceDelete deletes a namespace.
func (c *DefaultClient) NamespaceDelete(id string) error {
	uri := fmt.Sprintf(pathNamespace, c.base, id)
	err := c.delete(uri, nil)

	return err
}

// NamespaceUserList returns a list of related users for a namespace.
func (c *DefaultClient) NamespaceUserList(opts NamespaceUserParams) ([]*User, error) {
	var out []*User

	uri := fmt.Sprintf(pathNamespaceUser, c.base, opts.Namespace)
	err := c.get(uri, &out)

	return out, err
}

// NamespaceUserAppend appends a user to a namespace.
func (c *DefaultClient) NamespaceUserAppend(opts NamespaceUserParams) error {
	uri := fmt.Sprintf(pathNamespaceUser, c.base, opts.Namespace)
	err := c.patch(uri, opts, nil)

	return err
}

// NamespaceUserDelete remove a user from a namespace.
func (c *DefaultClient) NamespaceUserDelete(opts NamespaceUserParams) error {
	uri := fmt.Sprintf(pathNamespaceUser, c.base, opts.Namespace)
	err := c.delete(uri, opts)

	return err
}

// NamespaceTeamList returns a list of related teams for a namespace.
func (c *DefaultClient) NamespaceTeamList(opts NamespaceTeamParams) ([]*Team, error) {
	var out []*Team

	uri := fmt.Sprintf(pathNamespaceTeam, c.base, opts.Namespace)
	err := c.get(uri, &out)

	return out, err
}

// NamespaceTeamAppend appends a team to a namespace.
func (c *DefaultClient) NamespaceTeamAppend(opts NamespaceTeamParams) error {
	uri := fmt.Sprintf(pathNamespaceTeam, c.base, opts.Namespace)
	err := c.patch(uri, opts, nil)

	return err
}

// NamespaceTeamDelete remove a team from a namespace.
func (c *DefaultClient) NamespaceTeamDelete(opts NamespaceTeamParams) error {
	uri := fmt.Sprintf(pathNamespaceTeam, c.base, opts.Namespace)
	err := c.delete(uri, opts)

	return err
}
