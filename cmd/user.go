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

// User provides the sub-command for the user API.
func User() cli.Command {
	return cli.Command{
		Name:    "user",
		Aliases: []string{"u"},
		Usage:   "User related sub-commands",
		Subcommands: []cli.Command{
			{
				Name:      "list",
				Aliases:   []string{"ls"},
				Usage:     "List all users",
				ArgsUsage: " ",
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
				},
				Action: func(c *cli.Context) error {
					return Handle(c, UserShow)
				},
			},
			{
				Name:      "update",
				Usage:     "Update a user",
				ArgsUsage: " ",
				Flags: append(
					[]cli.Flag{
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
					},
				),
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
				Flags: append(
					[]cli.Flag{
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
					},
				),
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

	if len(records) == 0 {
		fmt.Fprintf(os.Stderr, "Empty result\n")
		return nil
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{"ID", "Slug", "Username"})

	for _, record := range records {
		table.Append(
			[]string{
				strconv.FormatInt(record.ID, 10),
				record.Slug,
				record.Username,
			},
		)
	}

	table.Render()
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
			"Username",
			record.Username,
		},
	)

	table.Append(
		[]string{
			"Email",
			record.Email,
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

	fmt.Fprintf(os.Stderr, "Successfully appended to user\n")
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

	fmt.Fprintf(os.Stderr, "Successfully removed from user\n")
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

// UserOrgAppend provides the sub-command to append a org to the user.
func UserOrgAppend(c *cli.Context, client umschlag.ClientAPI) error {
	err := client.UserOrgAppend(
		umschlag.UserOrgParams{
			User:      GetIdentifierParam(c),
			Org: GetOrgParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully appended to user\n")
	return nil
}

// UserOrgRemove provides the sub-command to remove a org from the user.
func UserOrgRemove(c *cli.Context, client umschlag.ClientAPI) error {
	err := client.UserOrgDelete(
		umschlag.UserOrgParams{
			User:      GetIdentifierParam(c),
			Org: GetOrgParam(c),
		},
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully removed from user\n")
	return nil
}
