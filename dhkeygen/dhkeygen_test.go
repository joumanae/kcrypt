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
		secret1 := dhkeygen.GenerateSecretKey()
		secret2 := dhkeygen.GenerateSecretKey()
		pk1 := dhkeygen.Power(big.NewInt(int64(base)), secret1)
		pk2 := dhkeygen.Power(big.NewInt(int64(base)), secret2)
		key1, err := dhkeygen.SharedKey(pk2, secret1, modulus)
		if err != nil {
			t.Fatal(err)
		}
		key2, err := dhkeygen.SharedKey(pk1, secret2, modulus)
		if err != nil {
			t.Fatal(err)
		}
		if key1.Cmp(key2) != 0 {
			t.Errorf("the two users do not have the same shared key: key 1: %v, key 2: %v", key1, key2)
		}
	})
}
func TestModulusCannotBeZero(t *testing.T) {
	_, err := dhkeygen.PublicKey(2, 0)
	if err != dhkeygen.ErrModulusCannotBeZero {
		t.Fatal("An unexpected error occured")
	}
	_, err = dhkeygen.SharedKey(big.NewInt(2), 2, 0)
	if err != dhkeygen.ErrModulusCannotBeZero {
		t.Fatal("An unexpected error occured")
	}
}

func TestBaseCannotBeZero(t *testing.T) {
	_, err := dhkeygen.PublicKey(0, 2)
	if err != dhkeygen.ErrBaseCannotBeZero {
		t.Fatal("An unexpected error occured")
	}
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
