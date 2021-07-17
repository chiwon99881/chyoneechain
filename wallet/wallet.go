package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"math/big"

	"github.com/chiwon99881/chyocoin/utils"
)

const (
	hashedMessage string = "1c5863cd55b5a4413fd59f054af57ba3c75c0698b3851d70f99b8de2d5c7338f"
	privateKey    string = "30770201010420012b384f984f2f03499d1f216578a558b360a67d6288144de1ffd7c2e3835212a00a06082a8648ce3d030107a14403420004784949dab986b0d7c0aad6008bddd7b348bb4494901c2295e44c50aae9b818d89ef74cd37b6dc42ca42b885f3abb6918bbb2aa22a5c22172e5c177c87114401a"
	signature     string = "64b29e9469b203d9b345211c988415ae983fef1122f834f19db2d55e5f0d4ae518aea8fa55bed6654d8dadf3da3eabef0adf10ac2180b37b7e8fc32b2960ff74"
)

// Start is wallet start func
func Start() {

	privBytes, err := hex.DecodeString(privateKey)
	utils.HandleError(err)

	restoredKey, err := x509.ParseECPrivateKey(privBytes)
	utils.HandleError(err)

	sigBytes, err := hex.DecodeString(signature)

	rBytes := sigBytes[:len(sigBytes)/2]
	sBytes := sigBytes[len(sigBytes)/2:]

	var bigR, bigS = big.Int{}, big.Int{}

	bigR.SetBytes(rBytes)
	bigS.SetBytes(sBytes)

}
