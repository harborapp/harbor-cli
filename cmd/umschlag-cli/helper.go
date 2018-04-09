package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/umschlag/umschlag-go/umschlag"
	"gopkg.in/urfave/cli.v2"
)

// sprigFuncMap provides template helpers provided by sprig.
var sprigFuncMap = sprig.TxtFuncMap()

// globalFuncMap provides global template helper functions.
var globalFuncMap = template.FuncMap{
	"taglist": func(s []*umschlag.Tag) string {
		res := []string{}

		for _, row := range s {
			res = append(res, row.String())
		}

		return strings.Join(res, ", ")
	},
	"orglist": func(s []*umschlag.Org) string {
		res := []string{}

		for _, row := range s {
			res = append(res, row.String())
		}

		return strings.Join(res, ", ")
	},
	"teamlist": func(s []*umschlag.Team) string {
		res := []string{}

		for _, row := range s {
			res = append(res, row.String())
		}

		return strings.Join(res, ", ")
	},
	"userlist": func(s []*umschlag.User) string {
		res := []string{}

		for _, row := range s {
			res = append(res, row.String())
		}

		return strings.Join(res, ", ")
	},
	"repolist": func(s []*umschlag.Repo) string {
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
		fmt.Println("error: you must provide an id or a slug.")
		os.Exit(1)
	}

	return val
}

// GetUserParam checks and returns the user id/slug parameter.
func GetUserParam(c *cli.Context) string {
	val := c.String("user")

	if val == "" {
		fmt.Println("error: you must provide a user id or slug.")
		os.Exit(1)
	}

	return val
}

// GetTeamParam checks and returns the team id/slug parameter.
func GetTeamParam(c *cli.Context) string {
	val := c.String("team")

	if val == "" {
		fmt.Println("error: you must provide a team id or slug.")
		os.Exit(1)
	}

	return val
}

// GetOrgParam checks and returns the org id/slug parameter.
func GetOrgParam(c *cli.Context) string {
	val := c.String("org")

	if val == "" {
		fmt.Println("error: <ou must provide a org id or slug.")
		os.Exit(1)
	}

	return val
}

// GetPermParam checks and returns the permission parameter.
func GetPermParam(c *cli.Context) string {
	val := c.String("perm")

	if val == "" {
		fmt.Println("error: you must provide a permission.")
		os.Exit(1)
	}

	for _, perm := range []string{"user", "admin", "owner"} {
		if perm == val {
			return val
		}
	}

	fmt.Println("error: invalid permission, can be user, admin or owner.")
	os.Exit(1)

	return ""
}
