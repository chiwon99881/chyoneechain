package main

import (
	"github.com/chiwon99881/chyocoin/explorer"
	"github.com/chiwon99881/chyocoin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
