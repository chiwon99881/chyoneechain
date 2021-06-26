package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/chiwon99881/chyocoin/explorer"
	"github.com/chiwon99881/chyocoin/rest"
)

func usage() {
	fmt.Printf("Welcome to chyoneecoin\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port [port number]:     Set port of the server\n")
	fmt.Printf("-mode [set mode]:        Set mode between 'html' and 'rest'\n")
	fmt.Println("-bothMode [true]:      If you send 'true', html and rest running both.")
	// 모든 함수를 제거하지만 그 전에 defer로 선언된 문을 먼저 실행하게 함
	runtime.Goexit()
}

// Start of the cli.go
func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")
	bothMode := flag.Bool("bothMode", false, "If you send 'true', html and rest are running both.")

	flag.Parse()
	if *bothMode == true {
		go rest.Start(*port)
		explorer.Start(*port + 1000)
	} else {
		switch *mode {
		case "rest":
			rest.Start(*port)
		case "html":
			explorer.Start(*port)
		default:
			usage()
		}
	}
}
