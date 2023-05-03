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
	a := CalculatePrivateKey(PNA, input1[0].aliceSecretKey, input1[0].modulus)
	fmt.Println("Alice's private key: ", a)
	PNB := CalculatePublicNumber(input2[0].base, input2[0].bobSecretKey, input2[0].modulus)
	fmt.Println("PNB: ", PNB)
	b := CalculatePrivateKey(PNB, input2[0].bobSecretKey, input2[0].modulus)
	fmt.Println("Bob's private key: ", b)
	if a != b {
		t.Errorf("Alice's private key %v != Bob's private key %v", a, b)
	}
}
