// https://cli.urfave.org/v2/getting-started/
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mwolfhoffman/contact-manager/commands"
	"github.com/mwolfhoffman/contact-manager/db"
	"github.com/urfave/cli/v2"
)

func init() {
	db.ConnectToDb()
}

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a contact",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "name",
						Usage: "name of contact",
					},
					&cli.StringFlag{
						Name:  "email",
						Usage: "email of contact",
					}, &cli.StringFlag{
						Name:  "phone",
						Usage: "phone number of contact",
					},
				},
				Action: func(cCtx *cli.Context) error {
					err := commands.AddContact(cCtx)
					if err != nil {
						fmt.Println(err)
					}
					return nil
				},
			},
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "list contacts",
				Action: func(cCtx *cli.Context) error {
					res, err := commands.List(cCtx)
					fmt.Println(res, err)
					return nil
				},
			},
			{
				Name:    "edit",
				Aliases: []string{"e"},
				Usage:   "edit contact",
				// Subcommands: []*cli.Command{
				// 	{
				// 		Name:  "add",
				// 		Usage: "add a new template",
				// 		Action: func(cCtx *cli.Context) error {
				// 			fmt.Println("new task template: ", cCtx.Args().First())
				// 			return nil
				// 		},
				// 	},
				// 	{
				// 		Name:  "remove",
				// 		Usage: "remove an existing template",
				// 		Action: func(cCtx *cli.Context) error {
				// 			fmt.Println("removed task template: ", cCtx.Args().First())
				// 			return nil
				// 		},
				// 	},
				// },
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// add, edit, list, remove, find, help
