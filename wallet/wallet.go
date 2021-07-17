package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"os"

	"github.com/chiwon99881/chyocoin/utils"
)

const (
	fileName string = "chyoneecoin.wallet"
)

type wallet struct {
	privateKey *ecdsa.PrivateKey
}

var w *wallet

func hasWalletFile() bool {
	_, err := os.Stat(fileName)
	// os.IsNotExist is return true when err indicate file not exists error.
	return !os.IsNotExist(err)
}

func createPrivateKey() *ecdsa.PrivateKey {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleError(err)
	return privateKey
}

func persistKey(key *ecdsa.PrivateKey) {
	bytes, err := x509.MarshalECPrivateKey(key)
	utils.HandleError(err)
	err = os.WriteFile(fileName, bytes, 0644)
	utils.HandleError(err)
}

// Wallet is return wallet memory address
func Wallet() *wallet {
	if w == nil {
		w = &wallet{}
		if hasWalletFile() {
			//restore file
		} else {
			key := createPrivateKey()
			persistKey(key)
			w.privateKey = key
		}
	}
	return w
}
