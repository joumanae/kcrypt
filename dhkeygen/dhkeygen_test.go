package dhkeygen_test // import "github.com/joumanae/cryptographywithgo/dhkeygen"

import (
	"math/big"
	"testing"

	dhkeygen "github.com/joumanae/cryptographywithgo/dhkeygen"
)

func FuzzTestPublicKey(f *testing.F) {
	f.Fuzz(func(t *testing.T, modulus int, base int) {
		dhkeygen.PublicKey(base, modulus)
	})
}

func FuzzTestSharedKey(f *testing.F) {
	f.Fuzz(func(t *testing.T, modulus int, base int, secret int) {

		pk1, err1 := dhkeygen.PublicKey(base, modulus)
		if err1 != nil {
			return
		}
		pk2, err2 := dhkeygen.PublicKey(base, modulus)
		if err2 != nil {
			return
		}
		secret1 := dhkeygen.GenerateSecretKey()
		secret2 := dhkeygen.GenerateSecretKey()
		key1 := dhkeygen.SharedKey(pk2, secret1, modulus)
		if modulus == 0 && base == 00 {
			t.Skip()
		}

		key2 := dhkeygen.SharedKey(pk1, secret2, modulus)

		if key1.Cmp(key2) != 0 {
			t.Errorf("the two users do not have the same shared key: key 1: %v, key 2: %v", key1, key2)
		}
	})
}

func TestParseBigInt(t *testing.T) {

	got, ok := dhkeygen.ParseBigInt("52")
	want := big.NewInt(52)
	if !ok {
		t.Fatal("problem parsing")
	}
	// cmp method
	if got.Cmp(want) != 0 {
		t.Errorf("want %v, got %v", want, got)
	}
}
