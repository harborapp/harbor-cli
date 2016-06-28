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
	pathProfile = "%s/api/profile"
)

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

// SetClient sets the default http client. This should
// be used in conjunction with golang.org/x/oauth2 to
// authenticate requests to the Solder API.
func (c *DefaultClient) SetClient(client *http.Client) {
	c.client = client
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
func (c *DefaultClient) delete(rawurl string) error {
	return c.do(rawurl, "DELETE", nil, nil)
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
