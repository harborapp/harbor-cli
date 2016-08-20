package cmd

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/umschlag/umschlag-go/umschlag"
	"github.com/urfave/cli"
)

// userFuncMap provides template helper functions.
var userFuncMap = template.FuncMap{
	"teamList": func(s []*umschlag.Team) string {
		res := []string{}

		for _, row := range s {
			res = append(res, row.String())
		}

		return strings.Join(res, ", ")
	},
	"orgList": func(s []*umschlag.Org) string {
		res := []string{}

		for _, row := range s {
			res = append(res, row.String())
		}

		return strings.Join(res, ", ")
	},
}

// tmplUserList represents a row within user listing.
var tmplUserList = "Slug: \x1b[33m{{ .Slug }} \x1b[0m" + `
ID: {{ .ID }}
Username: {{ .Username }}
`

// tmplUserShow represents a user within details view.
var tmplUserShow = "Slug: \x1b[33m{{ .Slug }} \x1b[0m" + `
ID: {{ .ID }}
Username: {{ .Username }}
Email: {{ .Email }}
Active: {{ .Active }}
Admin: {{ .Admin }}{{with .Teams}}
Teams: {{ teamList . }}{{end}}{{with .Orgs}}
Orgs: {{ orgList . }}{{end}}
Created: {{ .CreatedAt.Format "Mon Jan _2 15:04:05 MST 2006" }}
Updated: {{ .UpdatedAt.Format "Mon Jan _2 15:04:05 MST 2006" }}
`

// tmplUserTeamList represents a row within user team listing.
var tmplUserTeamList = "Slug: \x1b[33m{{ .Slug }} \x1b[0m" + `
ID: {{ .ID }}
Name: {{ .Name }}
`

// tmplUserOrgList represents a row within user org listing.
var tmplUserOrgList = "Slug: \x1b[33m{{ .Slug }} \x1b[0m" + `
ID: {{ .ID }}
Name: {{ .Name }}
`

// User provides the sub-command for the user API.
func User() cli.Command {
	return cli.Command{
		Name:  "user",
		Usage: "User related sub-commands",
		Subcommands: []cli.Command{
			{
				Name:      "list",
				Aliases:   []string{"ls"},
				Usage:     "List all users",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "format",
						Value: tmplUserList,
						Usage: "Custom output format",
					},
					cli.BoolFlag{
						Name:  "json",
						Usage: "Print in JSON format",
					},
					cli.BoolFlag{
						Name:  "xml",
						Usage: "Print in XML format",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, UserList)
				},
			},
			{
				Name:      "show",
				Usage:     "Display a user",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "User ID or slug to show",
					},
					cli.StringFlag{
						Name:  "format",
						Value: tmplUserShow,
						Usage: "Custom output format",
					},
					cli.BoolFlag{
						Name:  "json",
						Usage: "Print in JSON format",
					},
					cli.BoolFlag{
						Name:  "xml",
						Usage: "Print in XML format",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, UserShow)
				},
			},
			{
				Name:      "update",
				Usage:     "Update a user",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "User ID or slug to update",
					},
					cli.StringFlag{
						Name:  "slug",
						Value: "",
						Usage: "Provide a slug",
					},
					cli.StringFlag{
						Name:  "username",
						Value: "",
						Usage: "Provide an username",
					},
					cli.StringFlag{
						Name:  "email",
						Value: "",
						Usage: "Provide an email",
					},
					cli.StringFlag{
						Name:  "password",
						Value: "",
						Usage: "Provide a password",
					},
					cli.BoolFlag{
						Name:  "active",
						Usage: "Mark user as active",
					},
					cli.BoolFlag{
						Name:  "blocked",
						Usage: "Mark user as blocked",
					},
					cli.BoolFlag{
						Name:  "admin",
						Usage: "Mark user as admin",
					},
					cli.BoolFlag{
						Name:  "user",
						Usage: "Mark user as user",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, UserUpdate)
				},
			},
			{
				Name:      "delete",
				Aliases:   []string{"rm"},
				Usage:     "Delete a user",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "User ID or slug to show",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, UserDelete)
				},
			},
			{
				Name:      "create",
				Usage:     "Create a user",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "slug",
						Value: "",
						Usage: "Provide a slug",
					},
					cli.StringFlag{
						Name:  "username",
						Value: "",
						Usage: "Provide an username",
					},
					cli.StringFlag{
						Name:  "email",
						Value: "",
						Usage: "Provide an email",
					},
					cli.StringFlag{
						Name:  "password",
						Value: "",
						Usage: "Provide a password",
					},
					cli.BoolFlag{
						Name:  "active",
						Usage: "Mark user as active",
					},
					cli.BoolFlag{
						Name:  "blocked",
						Usage: "Mark user as blocked",
					},
					cli.BoolFlag{
						Name:  "admin",
						Usage: "Mark user as admin",
					},
					cli.BoolFlag{
						Name:  "user",
						Usage: "Mark user as user",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, UserCreate)
				},
			},
			{
				Name:      "team-list",
				Usage:     "List assigned teams",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "User ID or slug to list teams",
					},
					cli.StringFlag{
						Name:  "format",
						Value: tmplUserTeamList,
						Usage: "Custom output format",
					},
					cli.BoolFlag{
						Name:  "json",
						Usage: "Print in JSON format",
					},
					cli.BoolFlag{
						Name:  "xml",
						Usage: "Print in XML format",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, UserTeamList)
				},
			},
			{
				Name:      "team-append",
				Usage:     "Append a team to user",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "User ID or slug to append to",
					},
					cli.StringFlag{
						Name:  "team, t",
						Value: "",
						Usage: "Team ID or slug to append",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, UserTeamAppend)
				},
			},
			{
				Name:      "team-remove",
				Usage:     "Remove a team from user",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "User ID or slug to remove from",
					},
					cli.StringFlag{
						Name:  "team, t",
						Value: "",
						Usage: "Team ID or slug to remove",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, UserTeamRemove)
				},
			},
			{
				Name:      "org-list",
				Usage:     "List assigned orgs",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "User ID or slug to list orgs",
					},
					cli.StringFlag{
						Name:  "format",
						Value: tmplUserOrgList,
						Usage: "Custom output format",
					},
					cli.BoolFlag{
						Name:  "json",
						Usage: "Print in JSON format",
					},
					cli.BoolFlag{
						Name:  "xml",
						Usage: "Print in XML format",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, UserOrgList)
				},
			},
			{
				Name:      "org-append",
				Usage:     "Append a org to user",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "User ID or slug to append to",
					},
					cli.StringFlag{
						Name:  "org, t",
						Value: "",
						Usage: "Org ID or slug to append",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, UserOrgAppend)
				},
			},
			{
				Name:      "org-remove",
				Usage:     "Remove a org from user",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "User ID or slug to remove from",
					},
					cli.StringFlag{
						Name:  "org, t",
						Value: "",
						Usage: "Org ID or slug to remove",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, UserOrgRemove)
				},
			},
		},
	}
}

// UserList provides the sub-command to list all users.
func UserList(c *cli.Context, client umschlag.ClientAPI) error {
	records, err := client.UserList()

	if err != nil {
		return err
	}

	if c.IsSet("json") && c.IsSet("xml") {
		return fmt.Errorf("Conflict, you can only use JSON or XML at once!")
	}

	if c.Bool("xml") {
		res, err := xml.MarshalIndent(records, "", "  ")

		if err != nil {
			return err
		}

		fmt.Fprintf(os.Stdout, "%s\n", res)
		return nil
	}

	if c.Bool("json") {
		res, err := json.MarshalIndent(records, "", "  ")

		if err != nil {
			return err
		}

		fmt.Fprintf(os.Stdout, "%s\n", res)
		return nil
	}

	if len(records) == 0 {
		fmt.Fprintf(os.Stderr, "Empty result\n")
		return nil
	}

	tmpl, err := template.New(
		"_",
	).Funcs(
		userFuncMap,
	).Parse(
		fmt.Sprintf("%s\n", c.String("format")),
	)

	if err != nil {
		return err
	}

	for _, record := range records {
		err := tmpl.Execute(os.Stdout, record)

		if err != nil {
			return err
		}
	}

	return nil
}

// UserShow provides the sub-command to show user details.
func UserShow(c *cli.Context, client umschlag.ClientAPI) error {
	record, err := client.UserGet(
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	if c.IsSet("json") && c.IsSet("xml") {
		return fmt.Errorf("Conflict, you can only use JSON or XML at once!")
	}

	if c.Bool("xml") {
		res, err := xml.MarshalIndent(record, "", "  ")

		if err != nil {
			return err
		}

		fmt.Fprintf(os.Stdout, "%s\n", res)
		return nil
	}

	if c.Bool("json") {
		res, err := json.MarshalIndent(record, "", "  ")

		if err != nil {
			return err
		}

		fmt.Fprintf(os.Stdout, "%s\n", res)
		return nil
	}

	tmpl, err := template.New(
		"_",
	).Funcs(
		userFuncMap,
	).Parse(
		fmt.Sprintf("%s\n", c.String("format")),
	)

	if err != nil {
		return err
	}

	return tmpl.Execute(os.Stdout, record)
}

// UserDelete provides the sub-command to delete a user.
func UserDelete(c *cli.Context, client umschlag.ClientAPI) error {
	err := client.UserDelete(
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully delete\n")
	return nil
}

// UserUpdate provides the sub-command to update a user.
func UserUpdate(c *cli.Context, client umschlag.ClientAPI) error {
	record, err := client.UserGet(
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	changed := false

	if val := c.String("slug"); c.IsSet("slug") && val != record.Slug {
		record.Slug = val
		changed = true
	}

	if val := c.String("username"); c.IsSet("username") && val != record.Username {
		record.Username = val
		changed = true
	}

	if val := c.String("email"); c.IsSet("email") && val != record.Email {
		record.Email = val
		changed = true
	}

	if val := c.String("password"); c.IsSet("password") {
		record.Password = val
		changed = true
	}

	if c.IsSet("active") && c.IsSet("blocked") {
		return fmt.Errorf("Conflict, you can mark it only active OR blocked!")
	}

	if c.IsSet("active") {
		record.Active = true
		changed = true
	}

	if c.IsSet("blocked") {
		record.Active = false
		changed = true
	}

	if c.IsSet("admin") && c.IsSet("user") {
		return fmt.Errorf("Conflict, you can mark it only admin OR user!")
	}

	if c.IsSet("admin") {
		record.Admin = true
		changed = true
	}

	if c.IsSet("user") {
		record.Admin = false
		changed = true
	}

	if changed {
		_, patch := client.UserPatch(
			record,
		)

		if patch != nil {
			return patch
		}

		fmt.Fprintf(os.Stderr, "Successfully updated\n")
	} else {
		fmt.Fprintf(os.Stderr, "Nothing to update...\n")
	}

	return nil
}

// UserCreate provides the sub-command to create a user.
func UserCreate(c *cli.Context, client umschlag.ClientAPI) error {
	record := &umschlag.User{}

	if val := c.String("slug"); c.IsSet("slug") && val != "" {
		record.Slug = val
	}

	if val := c.String("username"); c.IsSet("username") && val != "" {
		record.Username = val
	} else {
		return fmt.Errorf("You must provide an username.")
	}

	if val := c.String("email"); c.IsSet("email") && val != "" {
		record.Email = val
	} else {
		return fmt.Errorf("You must provide an email.")
	}

	if val := c.String("password"); c.IsSet("password") && val != "" {
		record.Password = val
	} else {
		return fmt.Errorf("You must provide a password.")
	}

	if c.IsSet("active") && c.IsSet("blocked") {
		return fmt.Errorf("Conflict, you can mark it only active OR blocked!")
	}

	if c.IsSet("active") {
		record.Active = true
	}

	if c.IsSet("blocked") {
		record.Active = false
	}

	if c.IsSet("admin") && c.IsSet("user") {
		return fmt.Errorf("Conflict, you can mark it only admin OR user!")
	}

	if c.IsSet("admin") {
		record.Admin = true
	}

	if c.IsSet("user") {
		record.Admin = false
	}

	_, err := client.UserPost(
		record,
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully created\n")
	return nil
}

// UserTeamList provides the sub-command to list teams of the user.
func UserTeamList(c *cli.Context, client umschlag.ClientAPI) error {
	records, err := client.UserTeamList(
		umschlag.UserTeamParams{
			User: GetIdentifierParam(c),
		},
	)

	if err != nil {
		return err
	}

	if c.IsSet("json") && c.IsSet("xml") {
		return fmt.Errorf("Conflict, you can only use JSON or XML at once!")
	}

	if c.Bool("xml") {
		res, err := xml.MarshalIndent(records, "", "  ")

		if err != nil {
			return err
		}

		fmt.Fprintf(os.Stdout, "%s\n", res)
		return nil
	}

	if c.Bool("json") {
		res, err := json.MarshalIndent(records, "", "  ")

		if err != nil {
			return err
		}

		fmt.Fprintf(os.Stdout, "%s\n", res)
		return nil
	}

	if len(records) == 0 {
		fmt.Fprintf(os.Stderr, "Empty result\n")
		return nil
	}

	tmpl, err := template.New(
		"_",
	).Funcs(
		userFuncMap,
	).Parse(
		fmt.Sprintf("%s\n", c.String("format")),
	)

	if err != nil {
		return err
	}

	for _, record := range records {
		err := tmpl.Execute(os.Stdout, record)

		if err != nil {
			return err
		}
	}

	return nil
}

// UserTeamAppend provides the sub-command to append a team to the user.
func UserTeamAppend(c *cli.Context, client umschlag.ClientAPI) error {
	err := client.UserTeamAppend(
		umschlag.UserTeamParams{
			User: GetIdentifierParam(c),
			Team: GetTeamParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully appended to team\n")
	return nil
}

// UserTeamRemove provides the sub-command to remove a team from the user.
func UserTeamRemove(c *cli.Context, client umschlag.ClientAPI) error {
	err := client.UserTeamDelete(
		umschlag.UserTeamParams{
			User: GetIdentifierParam(c),
			Team: GetTeamParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully removed from team\n")
	return nil
}

// UserOrgList provides the sub-command to list orgs of the user.
func UserOrgList(c *cli.Context, client umschlag.ClientAPI) error {
	records, err := client.UserOrgList(
		umschlag.UserOrgParams{
			User: GetIdentifierParam(c),
		},
	)

	if err != nil {
		return err
	}

	if c.IsSet("json") && c.IsSet("xml") {
		return fmt.Errorf("Conflict, you can only use JSON or XML at once!")
	}

	if c.Bool("xml") {
		res, err := xml.MarshalIndent(records, "", "  ")

		if err != nil {
			return err
		}

		fmt.Fprintf(os.Stdout, "%s\n", res)
		return nil
	}

	if c.Bool("json") {
		res, err := json.MarshalIndent(records, "", "  ")

		if err != nil {
			return err
		}

		fmt.Fprintf(os.Stdout, "%s\n", res)
		return nil
	}

	if len(records) == 0 {
		fmt.Fprintf(os.Stderr, "Empty result\n")
		return nil
	}

	tmpl, err := template.New(
		"_",
	).Funcs(
		userFuncMap,
	).Parse(
		fmt.Sprintf("%s\n", c.String("format")),
	)

	if err != nil {
		return err
	}

	for _, record := range records {
		err := tmpl.Execute(os.Stdout, record)

		if err != nil {
			return err
		}
	}

	return nil
}

// UserOrgAppend provides the sub-command to append a org to the user.
func UserOrgAppend(c *cli.Context, client umschlag.ClientAPI) error {
	err := client.UserOrgAppend(
		umschlag.UserOrgParams{
			User: GetIdentifierParam(c),
			Org:  GetOrgParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully appended to org\n")
	return nil
}

// UserOrgRemove provides the sub-command to remove a org from the user.
func UserOrgRemove(c *cli.Context, client umschlag.ClientAPI) error {
	err := client.UserOrgDelete(
		umschlag.UserOrgParams{
			User: GetIdentifierParam(c),
			Org:  GetOrgParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully removed from org\n")
	return nil
}
