package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/chiwon99881/chyocoin/utils"
)

// Start is wallet start func
func Start() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleError(err)

	message := "i love you"
	// hash message into hex string
	hashMessage := utils.Hash(message)
	// convert byte from hex string
	hashAsBytes, err := hex.DecodeString(hashMessage)
	utils.HandleError(err)

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashAsBytes)
	utils.HandleError(err)

	fmt.Printf("R:%d\nS:%d\n", r, s)
}
