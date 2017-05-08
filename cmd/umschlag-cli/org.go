package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"text/template"

	"github.com/umschlag/umschlag-go/umschlag"
	"gopkg.in/urfave/cli.v2"
)

// orgFuncMap provides template helper functions.
var orgFuncMap = template.FuncMap{}

// tmplOrgList represents a row within org listing.
var tmplOrgList = "Slug: \x1b[33m{{ .Slug }} \x1b[0m" + `
ID: {{ .ID }}
Name: {{ .Name }}
`

// tmplOrgShow represents a org within details view.
var tmplOrgShow = "Slug: \x1b[33m{{ .Slug }} \x1b[0m" + `
ID: {{ .ID }}
Name: {{ .Name }}{{with .Registry}}
Registry: {{ .Name }}{{end}}{{with .Repos}}
Repos: {{ repoList . }}{{end}}{{with .Users}}
Users: {{ userList . }}{{end}}{{with .Teams}}
Teams: {{ teamList . }}{{end}}
Created: {{ .CreatedAt.Format "Mon Jan _2 15:04:05 MST 2006" }}
Updated: {{ .UpdatedAt.Format "Mon Jan _2 15:04:05 MST 2006" }}
`

// tmplOrgUserList represents a row within org user listing.
var tmplOrgUserList = "Slug: \x1b[33m{{ .User.Slug }} \x1b[0m" + `
ID: {{ .User.ID }}
Username: {{ .User.Username }}
Permission: {{ .Perm }}
`

// tmplOrgTeamList represents a row within org team listing.
var tmplOrgTeamList = "Slug: \x1b[33m{{ .Team.Slug }} \x1b[0m" + `
ID: {{ .Team.ID }}
Name: {{ .Team.Name }}
Permission: {{ .Perm }}
`

// Org provides the sub-command for the org API.
func Org() *cli.Command {
	return &cli.Command{
		Name:  "org",
		Usage: "Org related sub-commands",
		Subcommands: []*cli.Command{
			{
				Name:      "list",
				Aliases:   []string{"ls"},
				Usage:     "List all orgs",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "format",
						Value: tmplOrgList,
						Usage: "Custom output format",
					},
					&cli.BoolFlag{
						Name:  "json",
						Value: false,
						Usage: "Print in JSON format",
					},
					&cli.BoolFlag{
						Name:  "xml",
						Value: false,
						Usage: "Print in XML format",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, OrgList)
				},
			},
			{
				Name:      "show",
				Usage:     "Display a org",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Org ID or slug to show",
					},
					&cli.StringFlag{
						Name:  "format",
						Value: tmplOrgShow,
						Usage: "Custom output format",
					},
					&cli.BoolFlag{
						Name:  "json",
						Value: false,
						Usage: "Print in JSON format",
					},
					&cli.BoolFlag{
						Name:  "xml",
						Value: false,
						Usage: "Print in XML format",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, OrgShow)
				},
			},
			{
				Name:      "delete",
				Aliases:   []string{"rm"},
				Usage:     "Delete a org",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Org ID or slug to show",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, OrgDelete)
				},
			},
			{
				Name:      "update",
				Usage:     "Update a org",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Org ID or slug to update",
					},
					&cli.StringFlag{
						Name:  "slug",
						Value: "",
						Usage: "Provide a slug",
					},
					&cli.StringFlag{
						Name:  "name",
						Value: "",
						Usage: "Provide a name",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, OrgUpdate)
				},
			},
			{
				Name:      "create",
				Usage:     "Create a org",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "registry",
						Value: "",
						Usage: "Registry ID or slug",
					},
					&cli.StringFlag{
						Name:  "slug",
						Value: "",
						Usage: "Provide a slug",
					},
					&cli.StringFlag{
						Name:  "name",
						Value: "",
						Usage: "Provide a name",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, OrgCreate)
				},
			},
			{
				Name:  "user",
				Usage: "User assignments",
				Subcommands: []*cli.Command{
					{
						Name:      "list",
						Aliases:   []string{"ls"},
						Usage:     "List assigned users",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "Org ID or slug to list users",
							},
							&cli.StringFlag{
								Name:  "format",
								Value: tmplOrgUserList,
								Usage: "Custom output format",
							},
							&cli.BoolFlag{
								Name:  "json",
								Value: false,
								Usage: "Print in JSON format",
							},
							&cli.BoolFlag{
								Name:  "xml",
								Value: false,
								Usage: "Print in XML format",
							},
						},
						Action: func(c *cli.Context) error {
							return Handle(c, OrgUserList)
						},
					},
					{
						Name:      "append",
						Usage:     "Append a user to org",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "Org ID or slug to append to",
							},
							&cli.StringFlag{
								Name:  "user, u",
								Value: "",
								Usage: "User ID or slug to append",
							},
							&cli.StringFlag{
								Name:  "perm",
								Value: "user",
								Usage: "Permission for the user, can be user, admin or owner",
							},
						},
						Action: func(c *cli.Context) error {
							return Handle(c, OrgUserAppend)
						},
					},
					{
						Name:      "perm",
						Usage:     "Update org user permissions",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "Org ID or slug to update",
							},
							&cli.StringFlag{
								Name:  "user, u",
								Value: "",
								Usage: "User ID or slug to update",
							},
							&cli.StringFlag{
								Name:  "perm",
								Value: "user",
								Usage: "Permission for the user, can be user, admin or owner",
							},
						},
						Action: func(c *cli.Context) error {
							return Handle(c, OrgUserPerm)
						},
					},
					{
						Name:      "remove",
						Aliases:   []string{"rm"},
						Usage:     "Remove a user from org",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "Org ID or slug to remove from",
							},
							&cli.StringFlag{
								Name:  "user, u",
								Value: "",
								Usage: "User ID or slug to remove",
							},
						},
						Action: func(c *cli.Context) error {
							return Handle(c, OrgUserRemove)
						},
					},
				},
			},
			{
				Name:  "team",
				Usage: "Team assignments",
				Subcommands: []*cli.Command{
					{
						Name:      "list",
						Aliases:   []string{"ls"},
						Usage:     "List assigned teams",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "Org ID or slug to list teams",
							},
							&cli.StringFlag{
								Name:  "format",
								Value: tmplOrgTeamList,
								Usage: "Custom output format",
							},
							&cli.BoolFlag{
								Name:  "json",
								Value: false,
								Usage: "Print in JSON format",
							},
							&cli.BoolFlag{
								Name:  "xml",
								Value: false,
								Usage: "Print in XML format",
							},
						},
						Action: func(c *cli.Context) error {
							return Handle(c, OrgTeamList)
						},
					},
					{
						Name:      "append",
						Usage:     "Append a team to org",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "Org ID or slug to append to",
							},
							&cli.StringFlag{
								Name:  "team, t",
								Value: "",
								Usage: "Team ID or slug to append",
							},
							&cli.StringFlag{
								Name:  "perm",
								Value: "user",
								Usage: "Permission for the team, can be user, admin or owner",
							},
						},
						Action: func(c *cli.Context) error {
							return Handle(c, OrgTeamAppend)
						},
					},
					{
						Name:      "perm",
						Usage:     "Update org team permissions",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "Org ID or slug to update",
							},
							&cli.StringFlag{
								Name:  "team, t",
								Value: "",
								Usage: "Team ID or slug to update",
							},
							&cli.StringFlag{
								Name:  "perm",
								Value: "user",
								Usage: "Permission for the team, can be user, admin or owner",
							},
						},
						Action: func(c *cli.Context) error {
							return Handle(c, OrgTeamPerm)
						},
					},
					{
						Name:      "remove",
						Aliases:   []string{"rm"},
						Usage:     "Remove a team from org",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "Org ID or slug to remove from",
							},
							&cli.StringFlag{
								Name:  "team, t",
								Value: "",
								Usage: "Team ID or slug to remove",
							},
						},
						Action: func(c *cli.Context) error {
							return Handle(c, OrgTeamRemove)
						},
					},
				},
			},
		},
	}
}

// OrgList provides the sub-command to list all orgs.
func OrgList(c *cli.Context, client umschlag.ClientAPI) error {
	records, err := client.OrgList()

	if err != nil {
		return err
	}

	if c.IsSet("json") && c.IsSet("xml") {
		return fmt.Errorf("Conflict, you can only use JSON or XML at once")
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
		globalFuncMap,
	).Funcs(
		orgFuncMap,
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

// OrgShow provides the sub-command to show org details.
func OrgShow(c *cli.Context, client umschlag.ClientAPI) error {
	record, err := client.OrgGet(
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	if c.IsSet("json") && c.IsSet("xml") {
		return fmt.Errorf("Conflict, you can only use JSON or XML at once")
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
		globalFuncMap,
	).Funcs(
		orgFuncMap,
	).Parse(
		fmt.Sprintf("%s\n", c.String("format")),
	)

	if err != nil {
		return err
	}

	return tmpl.Execute(os.Stdout, record)
}

// OrgDelete provides the sub-command to delete a org.
func OrgDelete(c *cli.Context, client umschlag.ClientAPI) error {
	err := client.OrgDelete(
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully delete\n")
	return nil
}

// OrgUpdate provides the sub-command to update a org.
func OrgUpdate(c *cli.Context, client umschlag.ClientAPI) error {
	record, err := client.OrgGet(
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

	if val := c.String("name"); c.IsSet("name") && val != record.Name {
		record.Name = val
		changed = true
	}

	if changed {
		_, patch := client.OrgPatch(
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

// OrgCreate provides the sub-command to create a org.
func OrgCreate(c *cli.Context, client umschlag.ClientAPI) error {
	record := &umschlag.Org{}

	if c.String("registry") == "" {
		return fmt.Errorf("You must provide a registry ID or slug")
	}

	if c.IsSet("registry") {
		if match, _ := regexp.MatchString("^([0-9]+)$", c.String("registry")); match {
			if val, err := strconv.ParseInt(c.String("registry"), 10, 64); err == nil && val != 0 {
				record.RegistryID = val
			}
		} else {
			if c.String("registry") != "" {
				related, err := client.RegistryGet(
					c.String("registry"),
				)

				if err != nil {
					return err
				}

				if related.ID != record.RegistryID {
					record.RegistryID = related.ID
				}
			}
		}
	}

	if val := c.String("slug"); c.IsSet("slug") && val != "" {
		record.Slug = val
	}

	if val := c.String("name"); c.IsSet("name") && val != "" {
		record.Name = val
	} else {
		return fmt.Errorf("You must provide a name")
	}

	_, err := client.OrgPost(
		record,
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully created\n")
	return nil
}

// OrgUserList provides the sub-command to list users of the org.
func OrgUserList(c *cli.Context, client umschlag.ClientAPI) error {
	records, err := client.OrgUserList(
		umschlag.OrgUserParams{
			Org: GetIdentifierParam(c),
		},
	)

	if err != nil {
		return err
	}

	if c.IsSet("json") && c.IsSet("xml") {
		return fmt.Errorf("Conflict, you can only use JSON or XML at once")
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
		globalFuncMap,
	).Funcs(
		orgFuncMap,
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

// OrgUserAppend provides the sub-command to append a user to the org.
func OrgUserAppend(c *cli.Context, client umschlag.ClientAPI) error {
	err := client.OrgUserAppend(
		umschlag.OrgUserParams{
			Org:  GetIdentifierParam(c),
			User: GetUserParam(c),
			Perm: GetPermParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully appended to user\n")
	return nil
}

// OrgUserPerm provides the sub-command to update org user permissions.
func OrgUserPerm(c *cli.Context, client umschlag.ClientAPI) error {
	err := client.OrgUserPerm(
		umschlag.OrgUserParams{
			Org:  GetIdentifierParam(c),
			User: GetUserParam(c),
			Perm: GetPermParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully updated permissions\n")
	return nil
}

// OrgUserRemove provides the sub-command to remove a user from the org.
func OrgUserRemove(c *cli.Context, client umschlag.ClientAPI) error {
	err := client.OrgUserDelete(
		umschlag.OrgUserParams{
			Org:  GetIdentifierParam(c),
			User: GetUserParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully removed from user\n")
	return nil
}

// OrgTeamList provides the sub-command to list teams of the org.
func OrgTeamList(c *cli.Context, client umschlag.ClientAPI) error {
	records, err := client.OrgTeamList(
		umschlag.OrgTeamParams{
			Org: GetIdentifierParam(c),
		},
	)

	if err != nil {
		return err
	}

	if c.IsSet("json") && c.IsSet("xml") {
		return fmt.Errorf("Conflict, you can only use JSON or XML at once")
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
		globalFuncMap,
	).Funcs(
		orgFuncMap,
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

// OrgTeamAppend provides the sub-command to append a team to the org.
func OrgTeamAppend(c *cli.Context, client umschlag.ClientAPI) error {
	err := client.OrgTeamAppend(
		umschlag.OrgTeamParams{
			Org:  GetIdentifierParam(c),
			Team: GetTeamParam(c),
			Perm: GetPermParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully appended to team\n")
	return nil
}

// OrgTeamPerm provides the sub-command to update org team permissions.
func OrgTeamPerm(c *cli.Context, client umschlag.ClientAPI) error {
	err := client.OrgTeamPerm(
		umschlag.OrgTeamParams{
			Org:  GetIdentifierParam(c),
			Team: GetTeamParam(c),
			Perm: GetPermParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully updated permissions\n")
	return nil
}

// OrgTeamRemove provides the sub-command to remove a team from the org.
func OrgTeamRemove(c *cli.Context, client umschlag.ClientAPI) error {
	err := client.OrgTeamDelete(
		umschlag.OrgTeamParams{
			Org:  GetIdentifierParam(c),
			Team: GetTeamParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully removed from team\n")
	return nil
}
