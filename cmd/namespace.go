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

// Namespace provides the sub-command for the namespace API.
func Namespace() cli.Command {
	return cli.Command{
		Name:    "namespace",
		Aliases: []string{"n"},
		Usage:   "Namespace related sub-commands",
		Subcommands: []cli.Command{
			{
				Name:      "list",
				Aliases:   []string{"ls"},
				Usage:     "List all namespaces",
				ArgsUsage: " ",
				Action: func(c *cli.Context) error {
					return Handle(c, NamespaceList)
				},
			},
			{
				Name:      "show",
				Usage:     "Display a namespace",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Namespace ID or slug to show",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, NamespaceShow)
				},
			},
			{
				Name:      "update",
				Usage:     "Update a namespace",
				ArgsUsage: " ",
				Flags: append(
					[]cli.Flag{
						cli.StringFlag{
							Name:  "id, i",
							Value: "",
							Usage: "Namespace ID or slug to update",
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
					return Handle(c, NamespaceUpdate)
				},
			},
			{
				Name:      "delete",
				Aliases:   []string{"rm"},
				Usage:     "Delete a namespace",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Namespace ID or slug to show",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, NamespaceDelete)
				},
			},
			{
				Name:      "create",
				Usage:     "Create a namespace",
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
					return Handle(c, NamespaceCreate)
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
						Usage: "Namespace ID or slug to list users",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, NamespaceUserList)
				},
			},
			{
				Name:      "user-append",
				Usage:     "Append a user to namespace",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Namespace ID or slug to append to",
					},
					cli.StringFlag{
						Name:  "user, u",
						Value: "",
						Usage: "User ID or slug to append",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, NamespaceUserAppend)
				},
			},
			{
				Name:      "user-remove",
				Usage:     "Remove a user from namespace",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Namespace ID or slug to remove from",
					},
					cli.StringFlag{
						Name:  "user, u",
						Value: "",
						Usage: "User ID or slug to remove",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, NamespaceUserRemove)
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
						Usage: "Namespace ID or slug to list teams",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, NamespaceTeamList)
				},
			},
			{
				Name:      "team-append",
				Usage:     "Append a team to namespace",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Namespace ID or slug to append to",
					},
					cli.StringFlag{
						Name:  "team, t",
						Value: "",
						Usage: "Team ID or slug to append",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, NamespaceTeamAppend)
				},
			},
			{
				Name:      "team-remove",
				Usage:     "Remove a team from namespace",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Namespace ID or slug to remove from",
					},
					cli.StringFlag{
						Name:  "team, t",
						Value: "",
						Usage: "Team ID or slug to remove",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, NamespaceTeamRemove)
				},
			},
		},
	}
}

// NamespaceList provides the sub-command to list all namespaces.
func NamespaceList(c *cli.Context, client umschlag.ClientAPI) error {
	records, err := client.NamespaceList()

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

// NamespaceShow provides the sub-command to show namespace details.
func NamespaceShow(c *cli.Context, client umschlag.ClientAPI) error {
	record, err := client.NamespaceGet(
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

// NamespaceDelete provides the sub-command to delete a namespace.
func NamespaceDelete(c *cli.Context, client umschlag.ClientAPI) error {
	err := client.NamespaceDelete(
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully delete\n")
	return nil
}

// NamespaceUpdate provides the sub-command to update a namespace.
func NamespaceUpdate(c *cli.Context, client umschlag.ClientAPI) error {
	record, err := client.NamespaceGet(
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
		_, patch := client.NamespacePatch(
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

// NamespaceCreate provides the sub-command to create a namespace.
func NamespaceCreate(c *cli.Context, client umschlag.ClientAPI) error {
	record := &umschlag.Namespace{}

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

	_, err := client.NamespacePost(
		record,
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully created\n")
	return nil
}

// NamespaceUserList provides the sub-command to list users of the namespace.
func NamespaceUserList(c *cli.Context, client umschlag.ClientAPI) error {
	records, err := client.NamespaceUserList(
		umschlag.NamespaceUserParams{
			Namespace: GetIdentifierParam(c),
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

// NamespaceUserAppend provides the sub-command to append a user to the namespace.
func NamespaceUserAppend(c *cli.Context, client umschlag.ClientAPI) error {
	err := client.NamespaceUserAppend(
		umschlag.NamespaceUserParams{
			Namespace: GetIdentifierParam(c),
			User:      GetUserParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully appended to namespace\n")
	return nil
}

// NamespaceUserRemove provides the sub-command to remove a user from the namespace.
func NamespaceUserRemove(c *cli.Context, client umschlag.ClientAPI) error {
	err := client.NamespaceUserDelete(
		umschlag.NamespaceUserParams{
			Namespace: GetIdentifierParam(c),
			User:      GetUserParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully removed from namespace\n")
	return nil
}

// NamespaceTeamList provides the sub-command to list teams of the namespace.
func NamespaceTeamList(c *cli.Context, client umschlag.ClientAPI) error {
	records, err := client.NamespaceTeamList(
		umschlag.NamespaceTeamParams{
			Namespace: GetIdentifierParam(c),
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

// NamespaceTeamAppend provides the sub-command to append a team to the namespace.
func NamespaceTeamAppend(c *cli.Context, client umschlag.ClientAPI) error {
	err := client.NamespaceTeamAppend(
		umschlag.NamespaceTeamParams{
			Namespace: GetIdentifierParam(c),
			Team:      GetTeamParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully appended to namespace\n")
	return nil
}

// NamespaceTeamRemove provides the sub-command to remove a team from the namespace.
func NamespaceTeamRemove(c *cli.Context, client umschlag.ClientAPI) error {
	err := client.NamespaceTeamDelete(
		umschlag.NamespaceTeamParams{
			Namespace: GetIdentifierParam(c),
			Team:      GetTeamParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully removed from namespace\n")
	return nil
}
