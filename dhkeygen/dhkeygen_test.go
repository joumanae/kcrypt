package dhkeygen_test // import "github.com/joumanae/cryptographywithgo/dhkeygen"

import (
	"math/big"
	"testing"

	dhkeygen "github.com/joumanae/cryptographywithgo/dhkeygen"
)

// Add fuzzing to the test
// go test -v -count=1000 -coverprofile=coverage.out -covermode=atomic
// -run=TestPublicKey

func FuzzTestPublicKey(f *testing.F) {
	f.Fuzz(func(t *testing.T, modulus int, base int) {
		dhkeygen.PublicKey(base, modulus)
	})
}

func FuzzTestSharedKey(f *testing.F) {
	f.Fuzz(func(t *testing.T, modulus int, base int, secret int) {
		dhkeygen.SharedKey(dhkeygen.PublicKey(base, modulus), secret, modulus)
	})
}

func TestParseBigInt(t *testing.T) {
	got, ok := dhkeygen.ParseBigInt("42")
	want := big.NewInt(42)
	if !ok {
		t.Fatal("problem parsing")
	}
	// cmp method
	if got.Cmp(want) != 0 {
		t.Errorf("want %v, got %v", want, got)
	}
}
