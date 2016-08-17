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

// Registry provides the sub-command for the registry API.
func Registry() cli.Command {
	return cli.Command{
		Name:    "registry",
		Aliases: []string{"r"},
		Usage:   "Registry related sub-commands",
		Subcommands: []cli.Command{
			{
				Name:      "list",
				Aliases:   []string{"ls"},
				Usage:     "List all registries",
				ArgsUsage: " ",
				Action: func(c *cli.Context) error {
					return Handle(c, RegistryList)
				},
			},
			{
				Name:      "show",
				Usage:     "Display a registry",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Registry ID or slug to show",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, RegistryShow)
				},
			},
			{
				Name:      "update",
				Usage:     "Update a registry",
				ArgsUsage: " ",
				Flags: append(
					[]cli.Flag{
						cli.StringFlag{
							Name:  "id, i",
							Value: "",
							Usage: "Registry ID or slug to update",
						},
						cli.StringFlag{
							Name:  "slug",
							Value: "",
							Usage: "Provide a slug",
						},
						cli.StringFlag{
							Name:  "name",
							Value: "",
							Usage: "Provide an name",
						},
						cli.StringFlag{
							Name:  "host",
							Value: "",
							Usage: "Provide an host",
						},
						cli.BoolFlag{
							Name:  "use-ssl",
							Usage: "Should use SSL",
						},
						cli.BoolFlag{
							Name:  "no-ssl",
							Usage: "Should not use SSL",
						},
					},
				),
				Action: func(c *cli.Context) error {
					return Handle(c, RegistryUpdate)
				},
			},
			{
				Name:      "delete",
				Aliases:   []string{"rm"},
				Usage:     "Delete a registry",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Registry ID or slug to show",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, RegistryDelete)
				},
			},
			{
				Name:      "create",
				Usage:     "Create a registry",
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
							Usage: "Provide an name",
						},
						cli.StringFlag{
							Name:  "host",
							Value: "",
							Usage: "Provide an host",
						},
						cli.BoolFlag{
							Name:  "use-ssl",
							Usage: "Should use SSL",
						},
						cli.BoolFlag{
							Name:  "no-ssl",
							Usage: "Should not use SSL",
						},
					},
				),
				Action: func(c *cli.Context) error {
					return Handle(c, RegistryCreate)
				},
			},
		},
	}
}

// RegistryList provides the sub-command to list all registries.
func RegistryList(c *cli.Context, client umschlag.ClientAPI) error {
	records, err := client.RegistryList()

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

// RegistryShow provides the sub-command to show registry details.
func RegistryShow(c *cli.Context, client umschlag.ClientAPI) error {
	record, err := client.RegistryGet(
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
			"Host",
			record.Host,
		},
	)

	table.Append(
		[]string{
			"Use SSL",
			strconv.FormatBool(record.UseSSL),
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

// RegistryDelete provides the sub-command to delete a registry.
func RegistryDelete(c *cli.Context, client umschlag.ClientAPI) error {
	err := client.RegistryDelete(
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully delete\n")
	return nil
}

// RegistryUpdate provides the sub-command to update a registry.
func RegistryUpdate(c *cli.Context, client umschlag.ClientAPI) error {
	record, err := client.RegistryGet(
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

	if val := c.String("host"); c.IsSet("host") && val != record.Host {
		record.Host = val
		changed = true
	}

	if c.IsSet("use-ssl") && c.IsSet("no-ssl") {
		return fmt.Errorf("Conflict, you can mark it only use-ssl OR no-ssl!")
	}

	if c.IsSet("use-ssl") {
		record.UseSSL = true
		changed = true
	}

	if c.IsSet("no-ssl") {
		record.UseSSL = false
		changed = true
	}

	if changed {
		_, patch := client.RegistryPatch(
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

// RegistryCreate provides the sub-command to create a registry.
func RegistryCreate(c *cli.Context, client umschlag.ClientAPI) error {
	record := &umschlag.Registry{}

	if val := c.String("slug"); c.IsSet("slug") && val != "" {
		record.Slug = val
	}

	if val := c.String("name"); c.IsSet("name") && val != "" {
		record.Name = val
	} else {
		return fmt.Errorf("You must provide an name.")
	}

	if val := c.String("host"); c.IsSet("host") && val != "" {
		record.Host = val
	} else {
		return fmt.Errorf("You must provide an host.")
	}

	if c.IsSet("use-ssl") && c.IsSet("no-ssl") {
		return fmt.Errorf("Conflict, you can mark it only use-ssl OR no-ssl!")
	}

	if c.IsSet("use-ssl") {
		record.UseSSL = true
	}

	if c.IsSet("no-ssl") {
		record.UseSSL = false
	}

	_, err := client.RegistryPost(
		record,
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully created\n")
	return nil
}
