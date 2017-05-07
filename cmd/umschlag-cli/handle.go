package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/umschlag/umschlag-go/umschlag"
	"github.com/urfave/cli"
)

// HandleFunc is the real handle implementation.
type HandleFunc func(c *cli.Context, client umschlag.ClientAPI) error

// Handle wraps the command function handler.
func Handle(c *cli.Context, fn HandleFunc) error {
	var (
		server = c.GlobalString("server")
		token  = c.GlobalString("token")

		client umschlag.ClientAPI
	)

	if server == "" {
		fmt.Fprintf(os.Stderr, "Error: You must provide the server address.\n")
		os.Exit(1)
	}

	if _, err := url.Parse(server); err != nil {
		fmt.Fprintf(os.Stderr, "Error: Invalid server address, bad format?.\n")
		os.Exit(1)
	}

	if token == "" {
		client = umschlag.NewClient(
			server,
		)
	} else {
		client = umschlag.NewClientToken(
			server,
			token,
		)
	}

	if err := fn(c, client); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		os.Exit(2)
	}

	return nil
}
