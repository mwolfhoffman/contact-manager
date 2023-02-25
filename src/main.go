package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

var service *Service

func enrichContext(ctx context.Context) context.Context {
	devConnString := os.Getenv("DEV_DB_CONN_STRING")
	c := context.WithValue(ctx, "db", ConnectToDb(devConnString))
	return c
}

func init() {
	godotenv.Load(".env")
	ctx := context.Background()
	ctx = enrichContext(ctx)
	repo := NewRepository()
	service = NewService(ctx, repo)
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
					newContact := Contact{
						Name:  cCtx.Value("name").(string),
						Email: cCtx.Value("email").(string),
						Phone: cCtx.Value("phone").(string),
					}
					err := service.AddContact(newContact)
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
