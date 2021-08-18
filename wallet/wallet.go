package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"io/fs"
	"math/big"
	"os"

	"github.com/chiwon99881/chyocoin/utils"
)

const (
	fileName string = "chyoneecoin.wallet"
)

type fileLayer interface {
	hasWalletFile() bool
	WriteFile(name string, data []byte, perm fs.FileMode) error
	ReadFile(name string) ([]byte, error)
}

type layer struct {
}

func (layer) hasWalletFile() bool {
	_, err := os.Stat(fileName)
	// os.IsNotExist is return true when err indicate file not exists error.
	return !os.IsNotExist(err)
}

func (layer) WriteFile(name string, data []byte, perm fs.FileMode) error {
	return os.WriteFile(name, data, perm)
}

func (layer) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

var files fileLayer = layer{}

type wallet struct {
	privateKey *ecdsa.PrivateKey
	Address    string
}

var w *wallet

func createPrivateKey() *ecdsa.PrivateKey {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleError(err)
	return privateKey
}

func persistKey(key *ecdsa.PrivateKey) {
	bytes, err := x509.MarshalECPrivateKey(key)
	utils.HandleError(err)
	err = files.WriteFile(fileName, bytes, 0644)
	utils.HandleError(err)
}

func restoreKey() *ecdsa.PrivateKey {
	bytes, err := files.ReadFile(fileName)
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

// Sign is make signature with private Key
func Sign(payload string, w *wallet) string {
	payloadAsBytes, err := hex.DecodeString(payload)
	utils.HandleError(err)
	r, s, err := ecdsa.Sign(rand.Reader, w.privateKey, payloadAsBytes)
	utils.HandleError(err)
	signature := append(r.Bytes(), s.Bytes()...)
	return fmt.Sprintf("%x", signature)
}

func restoreBigInts(payload string) (*big.Int, *big.Int, error) {
	bytes, err := hex.DecodeString(payload)
	if err != nil {
		return nil, nil, err
	}
	aBytes := bytes[:len(bytes)/2]
	bBytes := bytes[len(bytes)/2:]
	bigA, bigB := big.Int{}, big.Int{}
	bigA.SetBytes(aBytes)
	bigB.SetBytes(bBytes)
	return &bigA, &bigB, nil
}

// Verify is verify signature
func Verify(signature, payload, address string) bool {
	r, s, err := restoreBigInts(signature)
	utils.HandleError(err)
	x, y, err := restoreBigInts(address)
	utils.HandleError(err)
	publicKey := ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X:     x,
		Y:     y,
	}
	hash, err := hex.DecodeString(payload)
	utils.HandleError(err)
	ok := ecdsa.Verify(&publicKey, hash, r, s)
	return ok
}

// Wallet is return wallet memory address
func Wallet() *wallet {
	if w == nil {
		w = &wallet{}
		if files.hasWalletFile() {
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
