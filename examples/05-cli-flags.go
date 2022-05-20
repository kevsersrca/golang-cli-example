package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	var version string

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "version",
				Aliases:     []string{"v"},
				Value:       "1.2.4",
				Usage:       "version of app",
				Destination: &version,
			},
		},
		Action: func(c *cli.Context) error {
			name := "someone"
			if version < "1.2.4" {
				fmt.Println("Please update your app, ", name)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
