package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("Welcome to chyoneecoin\n\n")
	fmt.Printf("Please use the following commands:\n\n")
	fmt.Printf("explorer:    Start the HTML explorer\n")
	fmt.Printf("rest:        Start the REST API(recommanded)\n")
	os.Exit(0)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	switch os.Args[1] {
	case "rest":
		fmt.Println("HI REST")
	case "explorer":
		fmt.Println("HI EXPLORER")
	default:
		usage()
	}
}
