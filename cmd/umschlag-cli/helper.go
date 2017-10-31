package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/umschlag/umschlag-cli/pkg/sdk"
	"gopkg.in/urfave/cli.v2"
)

// globalFuncMap provides global template helper functions.
var globalFuncMap = template.FuncMap{
	"split":    strings.Split,
	"join":     strings.Join,
	"toUpper":  strings.ToUpper,
	"toLower":  strings.ToLower,
	"contains": strings.Contains,
	"replace":  strings.Replace,
	"tagList": func(s []*sdk.Tag) string {
		res := []string{}

		for _, row := range s {
			res = append(res, row.String())
		}

		return strings.Join(res, ", ")
	},
	"orgList": func(s []*sdk.Org) string {
		res := []string{}

		for _, row := range s {
			res = append(res, row.String())
		}

		return strings.Join(res, ", ")
	},
	"teamList": func(s []*sdk.Team) string {
		res := []string{}

		for _, row := range s {
			res = append(res, row.String())
		}

		return strings.Join(res, ", ")
	},
	"userList": func(s []*sdk.User) string {
		res := []string{}

		for _, row := range s {
			res = append(res, row.String())
		}

		return strings.Join(res, ", ")
	},
	"repoList": func(s []*sdk.Repo) string {
		res := []string{}

		for _, row := range s {
			res = append(res, row.FullName)
		}

		return strings.Join(res, ", ")
	},
}

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

// GetOrgParam checks and returns the org id/slug parameter.
func GetOrgParam(c *cli.Context) string {
	val := c.String("org")

	if val == "" {
		fmt.Println("Error: You must provide a org ID or slug.")
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

// GetPermParam checks and returns the permission parameter.
func GetPermParam(c *cli.Context) string {
	val := c.String("perm")

	if val == "" {
		fmt.Println("Error: You must provide a permission.")
		os.Exit(1)
	}

	for _, perm := range []string{"user", "admin", "owner"} {
		if perm == val {
			return val
		}
	}

	fmt.Println("Error: Invalid permission, can be user, admin or owner.")
	os.Exit(1)

	return ""
}
