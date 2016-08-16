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

// Profile provides the sub-command for the profile API.
func Profile() cli.Command {
	return cli.Command{
		Name:  "profile",
		Usage: "Profile related sub-commands",
		Subcommands: []cli.Command{
			{
				Name:  "token",
				Usage: "Show your token",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "username",
						Value: "",
						Usage: "Username for authentication",
					},
					cli.StringFlag{
						Name:  "password",
						Value: "",
						Usage: "Password for authentication",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, ProfileToken)
				},
			},
			{
				Name:  "show",
				Usage: "Show profile details",
				Action: func(c *cli.Context) error {
					return Handle(c, ProfileShow)
				},
			},
			{
				Name:  "update",
				Usage: "Update profile details",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "slug",
						Value: "",
						Usage: "Provide a slug",
					},
					cli.StringFlag{
						Name:  "username",
						Value: "",
						Usage: "Provide a username",
					},
					cli.StringFlag{
						Name:  "email",
						Value: "",
						Usage: "Provide a email",
					},
					cli.StringFlag{
						Name:  "password",
						Value: "",
						Usage: "Provide a password",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, ProfileUpdate)
				},
			},
		},
	}
}

// ProfileToken provides the sub-command to show your token.
func ProfileToken(c *cli.Context, client umschlag.ClientAPI) error {
	if !client.IsAuthenticated() {
		if !c.IsSet("username") {
			return fmt.Errorf("Please provide a username")
		}

		if !c.IsSet("password") {
			return fmt.Errorf("Please provide a password")
		}

		login, err := client.AuthLogin(
			c.String("username"),
			c.String("password"),
		)

		if err != nil {
			return err
		}

		client = umschlag.NewClientToken(
			c.GlobalString("server"),
			login.Token,
		)
	}

	record, err := client.ProfileToken()

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stdout, "%s\n", record.Token)
	return nil
}

// ProfileShow provides the sub-command to show profile details.
func ProfileShow(c *cli.Context, client umschlag.ClientAPI) error {
	record, err := client.ProfileGet()

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

// ProfileUpdate provides the sub-command to update the profile.
func ProfileUpdate(c *cli.Context, client umschlag.ClientAPI) error {
	record, err := client.ProfileGet()

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
		_, patch := client.ProfilePatch(
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
