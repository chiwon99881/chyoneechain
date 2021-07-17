package wallet

import (
	"crypto/ecdsa"
	"os"
)

type wallet struct {
	privateKey *ecdsa.PrivateKey
}

var w *wallet

func hasWalletFile() bool {
	_, err := os.Stat("chyoneecoin.wallet")
	// os.IsNotExist is return true when err indicate file not exists error.
	return !os.IsNotExist(err)
}

// Wallet is return wallet memory address
func Wallet() *wallet {
	if w == nil {
		if hasWalletFile() {
			//restore file
		}
		// generate new privateKey
	}
	return w
}
