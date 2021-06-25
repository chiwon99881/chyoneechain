package main

import (
	"flag"
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

	// flag는 standard library에서 가져오는 녀석인데, 특정 os.Args의 입력값을 받아서 그 입력값의 flag들을 FlagSet으로 저장할 수 있다.
	rest := flag.NewFlagSet("rest", flag.ExitOnError)
	portFlag := rest.Int("port", 4000, "Server port")

	switch os.Args[1] {
	case "rest":
		// os.Args[2:]는 Slice또는 Array의 인덱스 2번 부터 끝을 의미
		rest.Parse(os.Args[2:])
	case "explorer":
		fmt.Println("HI EXPLORER")
	default:
		usage()
	}

	if rest.Parsed() {
		fmt.Println(*portFlag)
	}
}
