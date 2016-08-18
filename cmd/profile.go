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

// profileFuncMap provides template helper functions.
var profileFuncMap = template.FuncMap{}

// tmplProfileShow represents a profile within details view.
var tmplProfileShow = "Slug: \x1b[33m{{ .Slug }} \x1b[0m" + `
ID: {{ .ID }}
Username: {{ .Username }}
Email: {{ .Email }}
Created: {{ .CreatedAt.Format "Mon Jan _2 15:04:05 MST 2006" }}
Updated: {{ .UpdatedAt.Format "Mon Jan _2 15:04:05 MST 2006" }}
`

// Profile provides the sub-command for the profile API.
func Profile() cli.Command {
	return cli.Command{
		Name:  "profile",
		Usage: "Profile related sub-commands",
		Subcommands: []cli.Command{
			{
				Name:  "show",
				Usage: "Show profile details",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "format",
						Value: tmplProfileShow,
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
					return Handle(c, ProfileShow)
				},
			},
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

// ProfileShow provides the sub-command to show profile details.
func ProfileShow(c *cli.Context, client umschlag.ClientAPI) error {
	record, err := client.ProfileGet()

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
		profileFuncMap,
	).Parse(
		fmt.Sprintf("%s\n", c.String("format")),
	)

	if err != nil {
		return err
	}

	return tmpl.Execute(os.Stdout, record)
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
