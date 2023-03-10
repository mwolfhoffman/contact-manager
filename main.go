package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mwolfhoffman/contact-manager/src"
	"github.com/urfave/cli/v2"
)

var service *src.Service

func enrichContext(ctx context.Context) context.Context {
	devConnString := os.Getenv("DEV_DB_CONN_STRING")
	c := context.WithValue(ctx, "db", src.ConnectToDb(devConnString))
	return c
}

func init() {
	godotenv.Load(".env")
	ctx := context.Background()
	ctx = enrichContext(ctx)
	repo := src.NewRepository()
	service = src.NewService(ctx, repo)
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
					err := service.AddContact(cCtx)
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
					res, err := service.List(cCtx)
					fmt.Println(res, err)
					return nil
				},
			},
			{
				Name:    "search",
				Aliases: []string{"s"},
				Usage:   "search contacts",
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
					err := service.Search(cCtx)
					if err != nil {
						fmt.Println(err)
					}
					return nil
				},
			},
			{
				Name:    "edit",
				Aliases: []string{"e"},
				Usage:   "edit contact",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "id",
						Usage: "id of contact",
					},
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
					err := service.Edit(cCtx)
					if err != nil {
						fmt.Println(err)
					}
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// add, edit, list, remove, find, help
