package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/umschlag/umschlag-go/umschlag"
	"github.com/urfave/cli"
)

// Team provides the sub-command for the team API.
func Team() cli.Command {
	return cli.Command{
		Name:    "team",
		Aliases: []string{"t"},
		Usage:   "Team related sub-commands",
		Subcommands: []cli.Command{
			{
				Name:      "list",
				Aliases:   []string{"ls"},
				Usage:     "List all teams",
				ArgsUsage: " ",
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
				},
				Action: func(c *cli.Context) error {
					return Handle(c, TeamShow)
				},
			},
			{
				Name:      "update",
				Usage:     "Update a team",
				ArgsUsage: " ",
				Flags: append(
					[]cli.Flag{
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
				),
				Action: func(c *cli.Context) error {
					return Handle(c, TeamUpdate)
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
				Name:      "create",
				Usage:     "Create a team",
				ArgsUsage: " ",
				Flags: append(
					[]cli.Flag{
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
				),
				Action: func(c *cli.Context) error {
					return Handle(c, TeamCreate)
				},
			},
			{
				Name:      "user-list",
				Usage:     "List assigned users",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Team ID or slug to list users",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, TeamUserList)
				},
			},
			{
				Name:      "user-append",
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
				},
				Action: func(c *cli.Context) error {
					return Handle(c, TeamUserAppend)
				},
			},
			{
				Name:      "user-remove",
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
			{
				Name:      "org-list",
				Usage:     "List assigned orgs",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Team ID or slug to list orgs",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, TeamOrgList)
				},
			},
			{
				Name:      "org-append",
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
				},
				Action: func(c *cli.Context) error {
					return Handle(c, TeamOrgAppend)
				},
			},
			{
				Name:      "org-remove",
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
	}
}

// TeamList provides the sub-command to list all teams.
func TeamList(c *cli.Context, client umschlag.ClientAPI) error {
	records, err := client.TeamList()

	if err != nil {
		return err
	}

	if len(records) == 0 {
		fmt.Fprintf(os.Stderr, "Empty result\n")
		return nil
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{"ID", "Slug", "Name"})

	for _, record := range records {
		table.Append(
			[]string{
				strconv.FormatInt(record.ID, 10),
				record.Slug,
				record.Name,
			},
		)
	}

	table.Render()
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

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{"Key", "Value"})

	table.Append(
		[]string{
			"ID",
			strconv.FormatInt(record.ID, 10),
		},
	)

	table.Append(
		[]string{
			"Slug",
			record.Slug,
		},
	)

	table.Append(
		[]string{
			"Name",
			record.Name,
		},
	)

	table.Append(
		[]string{
			"Created",
			record.CreatedAt.Format(time.UnixDate),
		},
	)

	table.Append(
		[]string{
			"Updated",
			record.UpdatedAt.Format(time.UnixDate),
		},
	)

	table.Render()
	return nil
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

	if len(records) == 0 {
		fmt.Fprintf(os.Stderr, "Empty result\n")
		return nil
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{"User"})

	for _, record := range records {
		table.Append(
			[]string{
				record.Slug,
			},
		)
	}

	table.Render()
	return nil
}

// TeamUserAppend provides the sub-command to append a user to the team.
func TeamUserAppend(c *cli.Context, client umschlag.ClientAPI) error {
	err := client.TeamUserAppend(
		umschlag.TeamUserParams{
			Team: GetIdentifierParam(c),
			User: GetUserParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully appended to team\n")
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

	fmt.Fprintf(os.Stderr, "Successfully removed from team\n")
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

	if len(records) == 0 {
		fmt.Fprintf(os.Stderr, "Empty result\n")
		return nil
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{"Org"})

	for _, record := range records {
		table.Append(
			[]string{
				record.Slug,
			},
		)
	}

	table.Render()
	return nil
}

// TeamOrgAppend provides the sub-command to append a org to the team.
func TeamOrgAppend(c *cli.Context, client umschlag.ClientAPI) error {
	err := client.TeamOrgAppend(
		umschlag.TeamOrgParams{
			Team: GetIdentifierParam(c),
			Org:  GetOrgParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully appended to team\n")
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

	fmt.Fprintf(os.Stderr, "Successfully removed from team\n")
	return nil
}
