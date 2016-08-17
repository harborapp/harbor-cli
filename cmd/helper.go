package cmd

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

// GetIdentifierParam checks and returns the record id/slug parameter.
func GetIdentifierParam(c *cli.Context) string {
	val := c.String("id")

	if val == "" {
		fmt.Println("Error: You must provide an ID or a slug.")
		os.Exit(1)
	}

	return val
}

// GetUserParam checks and returns the user id/slug parameter.
func GetUserParam(c *cli.Context) string {
	val := c.String("user")

	if val == "" {
		fmt.Println("Error: You must provide a user ID or slug.")
		os.Exit(1)
	}

	return val
}

// GetTeamParam checks and returns the team id/slug parameter.
func GetTeamParam(c *cli.Context) string {
	val := c.String("team")

	if val == "" {
		fmt.Println("Error: You must provide a team ID or slug.")
		os.Exit(1)
	}

	return val
}

// GetNamespaceParam checks and returns the namespace id/slug parameter.
func GetNamespaceParam(c *cli.Context) string {
	val := c.String("namespace")

	if val == "" {
		fmt.Println("Error: You must provide a namespace ID or slug.")
		os.Exit(1)
	}

	return val
}

// GetRegistryParam checks and returns the registry id/slug parameter.
func GetRegistryParam(c *cli.Context) string {
	val := c.String("registry")

	if val == "" {
		fmt.Println("Error: You must provide a registry ID or slug.")
		os.Exit(1)
	}

	return val
}
