package main

import (
	"github.com/chiwon99881/chyocoin/blockchain"
	"github.com/chiwon99881/chyocoin/cli"
)

func main() {
	blockchain.Blockchain()
	cli.Start()
}
