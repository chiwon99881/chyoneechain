package main

import (
	"github.com/chiwon99881/chyocoin/cli"
	"github.com/chiwon99881/chyocoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
