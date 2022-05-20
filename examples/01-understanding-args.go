package main

import (
	"fmt"
	"os"
)

// USAGE
// <program> <arguments>
// ./1 -help
// ./1 -test
// ./1 command -h

func main() {
	// get all params
	params := os.Args

	// set args
	args := params[1:]

	// count of args
	count := len(args)

	// Program Name is always the first (implicit) argument
	programName := params[0]

	// List args
	for i, arg := range args {
		fmt.Printf("%d is: %s\n", i, arg)
	}

	fmt.Printf("Program Name: %s\n", programName)

	fmt.Printf("Total Arguments %d\n", count)

}

// ./1 help -flag=test -test 1
