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
		PK, err := dhkeygen.PublicKey(base, modulus)
		if err != nil {
			t.Skip()
		}
		dhkeygen.SharedKey(PK, secret, modulus)

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

func TestGenerateSecretKey(t *testing.T) {
	secret1 := dhkeygen.GenerateSecretKey()
	secret2 := dhkeygen.GenerateSecretKey()
	pk1, err := dhkeygen.PublicKey(2, 5)
	if err != nil {
		t.Fatal(err)
	}
	pk2, err := dhkeygen.PublicKey(2, 5)
	if err != nil {
		t.Fatal(err)
	}
	t1 := dhkeygen.SharedKey(big.NewInt(pk1.Int64()), secret1, 5)
	t2 := dhkeygen.SharedKey(big.NewInt(pk2.Int64()), secret2, 5)
	if t1.Cmp(t2) != 0 {
		t.Errorf("the two users do not have the same shared key: want %v, got %v", t1, t2)
	}
}
