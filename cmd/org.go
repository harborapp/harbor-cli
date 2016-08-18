package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/umschlag/umschlag-go/umschlag"
	"github.com/urfave/cli"
)

// Org provides the sub-command for the org API.
func Org() cli.Command {
	return cli.Command{
		Name:    "org",
		Aliases: []string{"o"},
		Usage:   "Org related sub-commands",
		Subcommands: []cli.Command{
			{
				Name:      "list",
				Aliases:   []string{"ls"},
				Usage:     "List all orgs",
				ArgsUsage: " ",
				Action: func(c *cli.Context) error {
					return Handle(c, OrgList)
				},
			},
			{
				Name:      "show",
				Usage:     "Display a org",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Org ID or slug to show",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, OrgShow)
				},
			},
			{
				Name:      "update",
				Usage:     "Update a org",
				ArgsUsage: " ",
				Flags: append(
					[]cli.Flag{
						cli.StringFlag{
							Name:  "id, i",
							Value: "",
							Usage: "Org ID or slug to update",
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
					return Handle(c, OrgUpdate)
				},
			},
			{
				Name:      "delete",
				Aliases:   []string{"rm"},
				Usage:     "Delete a org",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
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
				Name:      "create",
				Usage:     "Create a org",
				ArgsUsage: " ",
				Flags: append(
					[]cli.Flag{
						cli.StringFlag{
							Name:  "registry",
							Value: "",
							Usage: "Registry ID or slug",
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
					return Handle(c, OrgCreate)
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
						Usage: "Org ID or slug to list users",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, OrgUserList)
				},
			},
			{
				Name:      "user-append",
				Usage:     "Append a user to org",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Org ID or slug to append to",
					},
					cli.StringFlag{
						Name:  "user, u",
						Value: "",
						Usage: "User ID or slug to append",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, OrgUserAppend)
				},
			},
			{
				Name:      "user-remove",
				Usage:     "Remove a user from org",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Org ID or slug to remove from",
					},
					cli.StringFlag{
						Name:  "user, u",
						Value: "",
						Usage: "User ID or slug to remove",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, OrgUserRemove)
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
						Usage: "Org ID or slug to list teams",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, OrgTeamList)
				},
			},
			{
				Name:      "team-append",
				Usage:     "Append a team to org",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Org ID or slug to append to",
					},
					cli.StringFlag{
						Name:  "team, t",
						Value: "",
						Usage: "Team ID or slug to append",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, OrgTeamAppend)
				},
			},
			{
				Name:      "team-remove",
				Usage:     "Remove a team from org",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Org ID or slug to remove from",
					},
					cli.StringFlag{
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
	}
}

// OrgList provides the sub-command to list all orgs.
func OrgList(c *cli.Context, client umschlag.ClientAPI) error {
	records, err := client.OrgList()

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

// OrgShow provides the sub-command to show org details.
func OrgShow(c *cli.Context, client umschlag.ClientAPI) error {
	record, err := client.OrgGet(
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

	if record.Registry != nil {
		table.Append(
			[]string{
				"Registry",
				record.Registry.String(),
			},
		)
	}

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
		return fmt.Errorf("You must provide a registry ID or slug.")
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
		return fmt.Errorf("You must provide a name.")
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

// OrgUserAppend provides the sub-command to append a user to the org.
func OrgUserAppend(c *cli.Context, client umschlag.ClientAPI) error {
	err := client.OrgUserAppend(
		umschlag.OrgUserParams{
			Org: GetIdentifierParam(c),
			User:      GetUserParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully appended to org\n")
	return nil
}

// OrgUserRemove provides the sub-command to remove a user from the org.
func OrgUserRemove(c *cli.Context, client umschlag.ClientAPI) error {
	err := client.OrgUserDelete(
		umschlag.OrgUserParams{
			Org: GetIdentifierParam(c),
			User:      GetUserParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully removed from org\n")
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

	if len(records) == 0 {
		fmt.Fprintf(os.Stderr, "Empty result\n")
		return nil
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{"Team"})

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

// OrgTeamAppend provides the sub-command to append a team to the org.
func OrgTeamAppend(c *cli.Context, client umschlag.ClientAPI) error {
	err := client.OrgTeamAppend(
		umschlag.OrgTeamParams{
			Org: GetIdentifierParam(c),
			Team:      GetTeamParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully appended to org\n")
	return nil
}

// OrgTeamRemove provides the sub-command to remove a team from the org.
func OrgTeamRemove(c *cli.Context, client umschlag.ClientAPI) error {
	err := client.OrgTeamDelete(
		umschlag.OrgTeamParams{
			Org: GetIdentifierParam(c),
			Team:      GetTeamParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully removed from org\n")
	return nil
}
