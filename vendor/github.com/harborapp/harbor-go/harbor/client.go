package harbor

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/jackspirou/syscerts"
	"golang.org/x/oauth2"
)

const (
	pathProfile      = "%s/api/profile/self"
	pathProfileToken = "%s/api/profile/token"
	pathUsers        = "%s/api/users"
	pathUser         = "%s/api/users/%v"
	pathUserTeam     = "%s/api/users/%v/teams"
	pathTeams        = "%s/api/teams"
	pathTeam         = "%s/api/teams/%v"
	pathTeamUser     = "%s/api/teams/%v/users"
)

// DefaultClient implements the client interface.
type DefaultClient struct {
	client *http.Client
	base   string
}

// NewClient returns a client for the specified URL.
func NewClient(uri string) ClientAPI {
	return &DefaultClient{
		http.DefaultClient,
		uri,
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

	auther.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs: syscerts.SystemRootsPool(),
		},
	}

	return &DefaultClient{
		auther,
		uri,
	}
}

// SetClient sets the default http client. This should
// be used in conjunction with golang.org/x/oauth2 to
// authenticate requests to the Harbor API.
func (c *DefaultClient) SetClient(client *http.Client) {
	c.client = client
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

// Helper function for making an GET request.
func (c *DefaultClient) get(rawurl string, out interface{}) error {
	return c.do(rawurl, "GET", nil, out)
}

// Helper function for making an POST request.
func (c *DefaultClient) post(rawurl string, in, out interface{}) error {
	return c.do(rawurl, "POST", in, out)
}

// Helper function for making an PUT request.
func (c *DefaultClient) put(rawurl string, in, out interface{}) error {
	return c.do(rawurl, "PUT", in, out)
}

// Helper function for making an PATCH request.
func (c *DefaultClient) patch(rawurl string, in, out interface{}) error {
	return c.do(rawurl, "PATCH", in, out)
}

// Helper function for making an DELETE request.
func (c *DefaultClient) delete(rawurl string, in interface{}) error {
	return c.do(rawurl, "DELETE", in, nil)
}

// Helper function to make an HTTP request
func (c *DefaultClient) do(rawurl, method string, in, out interface{}) error {
	body, err := c.stream(
		rawurl,
		method,
		in,
		out,
	)

	if err != nil {
		return err
	}

	defer body.Close()

	if out != nil {
		return json.NewDecoder(body).Decode(out)
	}

	return nil
}

// Helper function to stream an HTTP request
func (c *DefaultClient) stream(rawurl, method string, in, out interface{}) (io.ReadCloser, error) {
	uri, err := url.Parse(rawurl)

	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter

	if in != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(in)

		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, uri.String(), buf)

	if err != nil {
		return nil, err
	}

	req.Header.Set(
		"User-Agent",
		"Harbor CLI",
	)

	if in != nil {
		req.Header.Set(
			"Content-Type",
			"application/json",
		)
	}

	resp, err := c.client.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode > http.StatusPartialContent {
		defer resp.Body.Close()
		out, _ := ioutil.ReadAll(resp.Body)

		msg := &Message{}
		parse := json.Unmarshal(out, msg)

		if parse != nil {
			return nil, fmt.Errorf(string(out))
		}

		return nil, fmt.Errorf(msg.Message)
	}

	return resp.Body, nil
}
