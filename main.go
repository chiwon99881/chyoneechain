package main

import (
	"fmt"

	"github.com/chiwon99881/chyocoin/blockchain"
)

func main() {
	chain := blockchain.GetBlockchain()
	fmt.Println(chain)
}
