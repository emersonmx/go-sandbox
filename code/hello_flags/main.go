package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	makeCmd := flag.NewFlagSet("make", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	switch os.Args[1] {
	case "make":
		makeCmd.Parse(os.Args[2:])
		templateName := os.Args[2]
		fmt.Println("Making template", templateName)
	case "list":
		listCmd.Parse(os.Args[2:])
		fmt.Println("Listing templates...")
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}
