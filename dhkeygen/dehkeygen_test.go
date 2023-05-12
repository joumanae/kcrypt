package diffiehellman_test

import (
	"math/big"
	"testing"

	diffiehellman "github.com/joumanae/dhkeygen"
)

// func TestSharedKey(t *testing.T) {

// 	input1 := []struct {
// 		modulus        int
// 		base           int
// 		aliceSecretKey int
// 	}{
// 		{17, 11, 7},
// 	}

// 	input2 := []struct {
// 		modulus      int
// 		base         int
// 		bobSecretKey int
// 	}{
// 		{17, 11, 15},
// 	}

// 	a := (SharedKey(PNA, input1[0].aliceSecretKey, input1[0].modulus))
// 	b := SharedKey(PNB, input2[0].bobSecretKey, input2[0].modulus)
// 	if a.Cmp(b) != 0 {
// 		t.Errorf("Expected %v, got %v", a, b)
// 	}
// }

func TestParseBigInt(t *testing.T) {
	got, ok := diffiehellman.ParseBigInt("42")
	want := big.NewInt(42)
	if !ok {
		t.Fatal("problem parsing")
	}
	// cmp method
	if got.Cmp(want) != 0 {
		t.Errorf("want %v, got %v", want, got)
	}
}
