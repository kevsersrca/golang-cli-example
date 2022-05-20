package main

import (
	"flag"
	"fmt"
	"os"
)

// USAGE
// <program> <command> <flag>
// ./3 apply
// ./3 apply --silent
// ./3 reset
// ./3 reset --loud
// ./3 reset --loud=false

func main() {
	args := os.Args

	f1 := flag.NewFlagSet("f1", flag.ContinueOnError)
	silent := f1.Bool("silent", false, "")

	f2 := flag.NewFlagSet("f2", flag.ContinueOnError)
	loud := f2.Bool("loud", true, "")

	switch args[1] {
	case "apply":
		if err := f1.Parse(args[2:]); err == nil {
			fmt.Println("apply", *silent)
		}
	case "reset":
		if err := f2.Parse(args[2:]); err == nil {
			fmt.Println("reset", *loud)
		}
	}
}
