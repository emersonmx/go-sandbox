package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	makeCmd := flag.NewFlagSet("make", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	Usage := func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		makeCmd.PrintDefaults()
		listCmd.PrintDefaults()
	}

	if len(os.Args) < 2 {
		Usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "make":
		makeCmd.Parse(os.Args[2:])
		templateName := os.Args[2]
		fmt.Println("Making template", templateName)
	case "list":
		listCmd.Parse(os.Args[2:])
		fmt.Println("Listing templates...")
	default:
		Usage()
		os.Exit(1)
	}
}
