package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"io/fs"
	"reflect"
	"testing"
)

const (
	testKey       string = "307702010104207c75100933cb07c5edeff8ae031914e2461c99063414f0b7ca0594762d2c4bcea00a06082a8648ce3d030107a1440342000478cf84a90bfff209ceb66d05de3bc204345aae9df4c474d317410de832d69bf48a40de798fcb77340861f9b089bdf4a49aa55677cf963d0984b3d0b02f6370e7"
	testPayload   string = "0030bda52b7d93551fbf191c49f658c1b9a5f6913141e0226d32f57ee079ea91"
	testSignature string = "4b280dc235bba91961ed4451e0248f2f9d52aa1802e9184bb2f2c7584f8436e8eee4825437f640c1b6ce82b9360f64fadd3e2e2789d001725a6a98ad9a476ce2"
)

type fakeLayer struct {
	fakeHasWalletFile func() bool
}

func (f fakeLayer) hasWalletFile() bool {
	return f.fakeHasWalletFile()
}

func (fakeLayer) WriteFile(name string, data []byte, perm fs.FileMode) error {
	return nil
}

func (fakeLayer) ReadFile(name string) ([]byte, error) {
	return x509.MarshalECPrivateKey(makeTestWallet().privateKey)
}

func TestWallet(t *testing.T) {
	t.Run("New Wallet is created", func(t *testing.T) {
		files = fakeLayer{
			fakeHasWalletFile: func() bool { return false },
		}
		tw := Wallet()
		if reflect.TypeOf(tw) != reflect.TypeOf(&wallet{}) {
			t.Error("New Wallet should return a new wallet instance")
		}
	})
	t.Run("Wallet is restored", func(t *testing.T) {
		files = fakeLayer{
			fakeHasWalletFile: func() bool { return true },
		}
		w = nil
		tw := Wallet()
		if reflect.TypeOf(tw) != reflect.TypeOf(&wallet{}) {
			t.Error("Restored wallet should return a new wallet instance")
		}
	})
}

func makeTestWallet() *wallet {
	w := &wallet{}
	b, _ := hex.DecodeString(testKey)
	pk, _ := x509.ParseECPrivateKey(b)
	w.privateKey = pk
	w.Address = addrFromKey(pk)
	return w
}

func TestSign(t *testing.T) {
	s := Sign(testPayload, makeTestWallet())
	_, err := hex.DecodeString(s)
	if err != nil {
		t.Errorf("Sign() should return a hex encoded string, but got %s", s)
	}
}

func TestVerify(t *testing.T) {
	type test struct {
		input string
		ok    bool
	}
	tests := []test{
		{input: testPayload, ok: true},
		{input: "0230bda52b7d93551fbf191c49f658c1b9a5f6913141e0226d32f57ee079ea91", ok: false},
	}

	for _, tc := range tests {
		w := makeTestWallet()
		ok := Verify(testSignature, tc.input, w.Address)
		if ok != tc.ok {
			t.Error("Verify() could not verify testSignature and testPayload.")
		}
	}

}

func TestRestoreBigInts(t *testing.T) {
	_, _, err := restoreBigInts("xx")
	if err == nil {
		t.Error("restoreBigInts() should return error when payload is not hex.")
	}
}
