package diffiehellman

import (
	"fmt"
	"testing"
)

func TestCalculatePrivateKey(t *testing.T) {

	input1 := []struct {
		modulus        int
		base           int
		aliceSecretKey int
	}{
		{17, 11, 7},
	}

	input2 := []struct {
		modulus      int
		base         int
		bobSecretKey int
	}{
		{17, 11, 15},
	}

	PNA := CalculatePublicNumber(input1[0].base, input1[0].aliceSecretKey, input1[0].modulus)
	fmt.Println("PNA: ", PNA)
	a := (CalculatePrivateKey(PNA, input1[0].aliceSecretKey, input1[0].modulus))
	PNB := CalculatePublicNumber(input2[0].base, input2[0].bobSecretKey, input2[0].modulus)
	b := CalculatePrivateKey(PNB, input2[0].bobSecretKey, input2[0].modulus)
	if a.Cmp(b) != 0 {
		t.Errorf("Expected %v, got %v", a, b)
	}
}
