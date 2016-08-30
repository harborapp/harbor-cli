package cmd

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"text/template"

	"github.com/umschlag/umschlag-go/umschlag"
	"github.com/urfave/cli"
)

// teamFuncMap provides template helper functions.
var teamFuncMap = template.FuncMap{}

// tmplTeamList represents a row within user listing.
var tmplTeamList = "Slug: \x1b[33m{{ .Slug }} \x1b[0m" + `
ID: {{ .ID }}
Name: {{ .Name }}
`

// tmplTeamShow represents a user within details view.
var tmplTeamShow = "Slug: \x1b[33m{{ .Slug }} \x1b[0m" + `
ID: {{ .ID }}
Name: {{ .Name }}{{with .Users}}
Users: {{ userList . }}{{end}}{{with .Orgs}}
Orgs: {{ orgList . }}{{end}}
Created: {{ .CreatedAt.Format "Mon Jan _2 15:04:05 MST 2006" }}
Updated: {{ .UpdatedAt.Format "Mon Jan _2 15:04:05 MST 2006" }}
`

// tmplTeamUserList represents a row within team user listing.
var tmplTeamUserList = "Slug: \x1b[33m{{ .User.Slug }} \x1b[0m" + `
ID: {{ .User.ID }}
Name: {{ .User.Name }}
Permission: {{ .Perm }}
`

// tmplTeamOrgList represents a row within team org listing.
var tmplTeamOrgList = "Slug: \x1b[33m{{ .Org.Slug }} \x1b[0m" + `
ID: {{ .Org.ID }}
Name: {{ .Org.Name }}
Permission: {{ .Perm }}
`

// Team provides the sub-command for the team API.
func Team() cli.Command {
	return cli.Command{
		Name:  "team",
		Usage: "Team related sub-commands",
		Subcommands: []cli.Command{
			{
				Name:      "list",
				Aliases:   []string{"ls"},
				Usage:     "List all teams",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "format",
						Value: tmplTeamList,
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
					return Handle(c, TeamList)
				},
			},
			{
				Name:      "show",
				Usage:     "Display a team",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Team ID or slug to show",
					},
					cli.StringFlag{
						Name:  "format",
						Value: tmplTeamShow,
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
					return Handle(c, TeamShow)
				},
			},
			{
				Name:      "delete",
				Aliases:   []string{"rm"},
				Usage:     "Delete a team",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Team ID or slug to show",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, TeamDelete)
				},
			},
			{
				Name:      "update",
				Usage:     "Update a team",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Team ID or slug to update",
					},
					cli.StringFlag{
						Name:  "slug",
						Value: "",
						Usage: "Provide a slug",
					},
					cli.StringFlag{
						Name:  "name",
						Value: "",
						Usage: "Provide a name",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, TeamUpdate)
				},
			},
			{
				Name:      "create",
				Usage:     "Create a team",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "slug",
						Value: "",
						Usage: "Provide a slug",
					},
					cli.StringFlag{
						Name:  "name",
						Value: "",
						Usage: "Provide a name",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, TeamCreate)
				},
			},
			{
				Name:  "user",
				Usage: "User assignments",
				Subcommands: []cli.Command{
					{
						Name:      "list",
						Aliases:   []string{"ls"},
						Usage:     "List assigned users",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "Team ID or slug to list users",
							},
							cli.StringFlag{
								Name:  "format",
								Value: tmplTeamUserList,
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
							return Handle(c, TeamUserList)
						},
					},
					{
						Name:      "append",
						Usage:     "Append a user to team",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "Team ID or slug to append to",
							},
							cli.StringFlag{
								Name:  "user, u",
								Value: "",
								Usage: "User ID or slug to append",
							},
							cli.StringFlag{
								Name:  "perm",
								Value: "user",
								Usage: "Permission for the user, can be user, admin or owner",
							},
						},
						Action: func(c *cli.Context) error {
							return Handle(c, TeamUserAppend)
						},
					},
					{
						Name:      "perm",
						Usage:     "Update team user permissions",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "Team ID or slug to update",
							},
							cli.StringFlag{
								Name:  "user, u",
								Value: "",
								Usage: "User ID or slug to update",
							},
							cli.StringFlag{
								Name:  "perm",
								Value: "user",
								Usage: "Permission for the user, can be user, admin or owner",
							},
						},
						Action: func(c *cli.Context) error {
							return Handle(c, TeamUserPerm)
						},
					},
					{
						Name:      "remove",
						Aliases:   []string{"rm"},
						Usage:     "Remove a user from team",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "Team ID or slug to remove from",
							},
							cli.StringFlag{
								Name:  "user, u",
								Value: "",
								Usage: "User ID or slug to remove",
							},
						},
						Action: func(c *cli.Context) error {
							return Handle(c, TeamUserRemove)
						},
					},
				},
			},
			{
				Name:  "org",
				Usage: "Org assignments",
				Subcommands: []cli.Command{
					{
						Name:      "list",
						Aliases:   []string{"ls"},
						Usage:     "List assigned orgs",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "Team ID or slug to list orgs",
							},
							cli.StringFlag{
								Name:  "format",
								Value: tmplTeamOrgList,
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
							return Handle(c, TeamOrgList)
						},
					},
					{
						Name:      "append",
						Usage:     "Append a org to team",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "Team ID or slug to append to",
							},
							cli.StringFlag{
								Name:  "org, u",
								Value: "",
								Usage: "Org ID or slug to append",
							},
							cli.StringFlag{
								Name:  "perm",
								Value: "user",
								Usage: "Permission for the team, can be user, admin or owner",
							},
						},
						Action: func(c *cli.Context) error {
							return Handle(c, TeamOrgAppend)
						},
					},
					{
						Name:      "perm",
						Usage:     "Update team org permissions",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "Team ID or slug to update",
							},
							cli.StringFlag{
								Name:  "user, u",
								Value: "",
								Usage: "Org ID or slug to update",
							},
							cli.StringFlag{
								Name:  "perm",
								Value: "user",
								Usage: "Permission for the team, can be user, admin or owner",
							},
						},
						Action: func(c *cli.Context) error {
							return Handle(c, TeamOrgPerm)
						},
					},
					{
						Name:      "remove",
						Aliases:   []string{"rm"},
						Usage:     "Remove a org from team",
						ArgsUsage: " ",
						Flags: []cli.Flag{
							cli.StringFlag{
								Name:  "id, i",
								Value: "",
								Usage: "Team ID or slug to remove from",
							},
							cli.StringFlag{
								Name:  "org, u",
								Value: "",
								Usage: "Org ID or slug to remove",
							},
						},
						Action: func(c *cli.Context) error {
							return Handle(c, TeamOrgRemove)
						},
					},
				},
			},
		},
	}
}

// TeamList provides the sub-command to list all teams.
func TeamList(c *cli.Context, client umschlag.ClientAPI) error {
	records, err := client.TeamList()

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
		teamFuncMap,
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

// TeamShow provides the sub-command to show team details.
func TeamShow(c *cli.Context, client umschlag.ClientAPI) error {
	record, err := client.TeamGet(
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
		teamFuncMap,
	).Parse(
		fmt.Sprintf("%s\n", c.String("format")),
	)

	if err != nil {
		return err
	}

	return tmpl.Execute(os.Stdout, record)
}

// TeamDelete provides the sub-command to delete a team.
func TeamDelete(c *cli.Context, client umschlag.ClientAPI) error {
	err := client.TeamDelete(
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully delete\n")
	return nil
}

// TeamUpdate provides the sub-command to update a team.
func TeamUpdate(c *cli.Context, client umschlag.ClientAPI) error {
	record, err := client.TeamGet(
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
		_, patch := client.TeamPatch(
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

// TeamCreate provides the sub-command to create a team.
func TeamCreate(c *cli.Context, client umschlag.ClientAPI) error {
	record := &umschlag.Team{}

	if val := c.String("slug"); c.IsSet("slug") && val != "" {
		record.Slug = val
	}

	if val := c.String("name"); c.IsSet("name") && val != "" {
		record.Name = val
	} else {
		return fmt.Errorf("You must provide a name.")
	}

	_, err := client.TeamPost(
		record,
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully created\n")
	return nil
}

// TeamUserList provides the sub-command to list users of the team.
func TeamUserList(c *cli.Context, client umschlag.ClientAPI) error {
	records, err := client.TeamUserList(
		umschlag.TeamUserParams{
			Team: GetIdentifierParam(c),
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
		teamFuncMap,
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

// TeamUserAppend provides the sub-command to append a user to the team.
func TeamUserAppend(c *cli.Context, client umschlag.ClientAPI) error {
	err := client.TeamUserAppend(
		umschlag.TeamUserParams{
			Team: GetIdentifierParam(c),
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

// TeamUserPerm provides the sub-command to update team user permissions.
func TeamUserPerm(c *cli.Context, client umschlag.ClientAPI) error {
	err := client.TeamUserPerm(
		umschlag.TeamUserParams{
			Team: GetIdentifierParam(c),
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

// TeamUserRemove provides the sub-command to remove a user from the team.
func TeamUserRemove(c *cli.Context, client umschlag.ClientAPI) error {
	err := client.TeamUserDelete(
		umschlag.TeamUserParams{
			Team: GetIdentifierParam(c),
			User: GetUserParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully removed from user\n")
	return nil
}

// TeamOrgList provides the sub-command to list orgs of the team.
func TeamOrgList(c *cli.Context, client umschlag.ClientAPI) error {
	records, err := client.TeamOrgList(
		umschlag.TeamOrgParams{
			Team: GetIdentifierParam(c),
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
		teamFuncMap,
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

// TeamOrgAppend provides the sub-command to append a org to the team.
func TeamOrgAppend(c *cli.Context, client umschlag.ClientAPI) error {
	err := client.TeamOrgAppend(
		umschlag.TeamOrgParams{
			Team: GetIdentifierParam(c),
			Org:  GetOrgParam(c),
			Perm: GetPermParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully appended to org\n")
	return nil
}

// TeamOrgPerm provides the sub-command to update team org permissions.
func TeamOrgPerm(c *cli.Context, client umschlag.ClientAPI) error {
	err := client.TeamOrgPerm(
		umschlag.TeamOrgParams{
			Team: GetIdentifierParam(c),
			Org:  GetOrgParam(c),
			Perm: GetPermParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully updated permissions\n")
	return nil
}

// TeamOrgRemove provides the sub-command to remove a org from the team.
func TeamOrgRemove(c *cli.Context, client umschlag.ClientAPI) error {
	err := client.TeamOrgDelete(
		umschlag.TeamOrgParams{
			Team: GetIdentifierParam(c),
			Org:  GetOrgParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully removed from org\n")
	return nil
}
