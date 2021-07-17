package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/chiwon99881/chyocoin/utils"
)

const (
	fileName string = "chyoneecoin.wallet"
)

type wallet struct {
	privateKey *ecdsa.PrivateKey
	Address    string
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

func restoreKey() *ecdsa.PrivateKey {
	bytes, err := os.ReadFile(fileName)
	utils.HandleError(err)
	restorePrivateKey, err := x509.ParseECPrivateKey(bytes)
	utils.HandleError(err)
	return restorePrivateKey
}

func addrFromKey(key *ecdsa.PrivateKey) string {
	x := key.X.Bytes()
	y := key.Y.Bytes()
	z := append(x, y...)
	return fmt.Sprintf("%x", z)
}

func sign(payload string, w *wallet) string {
	payloadAsBytes, err := hex.DecodeString(payload)
	utils.HandleError(err)
	r, s, err := ecdsa.Sign(rand.Reader, w.privateKey, payloadAsBytes)
	utils.HandleError(err)
	signature := append(r.Bytes(), s.Bytes()...)
	return fmt.Sprintf("%x", signature)
}

func verify(signature, payload, publicKey string) bool {
	//ecdsa.Verify()
}

// Wallet is return wallet memory address
func Wallet() *wallet {
	if w == nil {
		w = &wallet{}
		if hasWalletFile() {
			key := restoreKey()
			w.privateKey = key
		} else {
			key := createPrivateKey()
			persistKey(key)
			w.privateKey = key
		}
		w.Address = addrFromKey(w.privateKey)
	}
	return w
}
