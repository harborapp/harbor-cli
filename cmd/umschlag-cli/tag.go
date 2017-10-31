package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"text/template"

	"github.com/umschlag/umschlag-cli/pkg/sdk"
	"gopkg.in/urfave/cli.v2"
)

// tagFuncMap provides template helper functions.
var tagFuncMap = template.FuncMap{}

// tmplTagList represents a row within tag listing.
var tmplTagList = "Slug: \x1b[33m{{ .Slug }} \x1b[0m" + `
ID: {{ .ID }}
Name: {{ .FullName }}
`

// tmplTagShow represents a tag within details view.
var tmplTagShow = "Slug: \x1b[33m{{ .Slug }} \x1b[0m" + `
ID: {{ .ID }}
Name: {{ .FullName }}
Created: {{ .CreatedAt.Format "Mon Jan _2 15:04:05 MST 2006" }}
Updated: {{ .UpdatedAt.Format "Mon Jan _2 15:04:05 MST 2006" }}
`

// Tag provides the sub-command for the tag API.
func Tag() *cli.Command {
	return &cli.Command{
		Name:  "tag",
		Usage: "Tag related sub-commands",
		Subcommands: []*cli.Command{
			{
				Name:      "list",
				Aliases:   []string{"ls"},
				Usage:     "List all tags",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "format",
						Value: tmplTagList,
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
					return Handle(c, TagList)
				},
			},
			{
				Name:      "show",
				Usage:     "Display a tag",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Tag ID or slug to show",
					},
					&cli.StringFlag{
						Name:  "format",
						Value: tmplTagShow,
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
					return Handle(c, TagShow)
				},
			},
			{
				Name:      "delete",
				Aliases:   []string{"rm"},
				Usage:     "Delete a tag",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "id, i",
						Value: "",
						Usage: "Tag ID or slug to show",
					},
				},
				Action: func(c *cli.Context) error {
					return Handle(c, TagDelete)
				},
			},
		},
	}
}

// TagList provides the sub-command to list all tags.
func TagList(c *cli.Context, client sdk.ClientAPI) error {
	records, err := client.TagList()

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
		tagFuncMap,
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

// TagShow provides the sub-command to show tag details.
func TagShow(c *cli.Context, client sdk.ClientAPI) error {
	record, err := client.TagGet(
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
		tagFuncMap,
	).Parse(
		fmt.Sprintf("%s\n", c.String("format")),
	)

	if err != nil {
		return err
	}

	return tmpl.Execute(os.Stdout, record)
}

// TagDelete provides the sub-command to delete a tag.
func TagDelete(c *cli.Context, client sdk.ClientAPI) error {
	err := client.TagDelete(
		GetIdentifierParam(c),
	)

	if err != nil {
		return err
	}

	fmt.Fprintf(os.Stderr, "Successfully delete\n")
	return nil
}
