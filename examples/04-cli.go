package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{}

	app.Copyright = "2022 Ankara Gophers"
	app.Description = "Basic CRUD"
	app.Authors = []*cli.Author{
		{
			Name:  "Kevser Sirca",
			Email: "kevser.sirca@gmail.com",
		},
	}
	app.EnableBashCompletion = true
	app.Usage = "<program> <command> <options>"

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
