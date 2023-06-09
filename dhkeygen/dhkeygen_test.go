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
			t.Error(err1)
		}
		pk2, err2 := dhkeygen.PublicKey(base, modulus)
		if err2 != nil {
			t.Error(err2)
		}
		s1 := dhkeygen.GenerateSecretKey()
		s2 := dhkeygen.GenerateSecretKey()

		key1, err1 := dhkeygen.SharedKey(pk2, s1, modulus)
		if err1 != nil {
			t.Errorf("error %v", err1)
		}
		if modulus == 0 && base == 00 {
			t.Skip()
		}

		key2, err2 := dhkeygen.SharedKey(pk1, s2, modulus)
		if err2 != nil {
			t.Errorf("error %v", err2)
		}

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
