package diffiehellman

import "flag"

type alice struct {
	privateKey int
	publicKey  int
}

type bob struct {
	privateKey int
	publicKey  int
}

type diffieHellman struct {
	modulus int
	base    int
}

// Alice then chooses a secret integer a, and sends Bob A=g^a mod p

// Bob chooses a secret integer b and sends Alice B= g^b mod p

// Alice computes s = B^a mod p

// Bob computes s = A^b mod p

func (a *alice) sendsSecretInteger() int {
	var df diffieHellman
	A := df.base ^ a.privateKey%df.modulus // A = g^a mod p
	return A
}

func (b *bob) sendsSecretInteger() int {
	var df diffieHellman
	B := df.base ^ b.privateKey%df.modulus // B = g^b mod p
	return B
}

//B^a mod p
func (a *alice) generatePrivateKey(modulus int, base int, publicKey int) int {
	publicKey = a.publicKey

	return 0
}

// need to figure out how to send it to alice then bob
// need to add the tests
//A^b mod p
func (b *bob) generatePrivateKey(modulus int, base int, publicKey int) int {
	publicKey = b.publicKey

	return 0
}

func Main() int {
	var b bob
	var a alice
	mod := flag.Int("modulus", 0, "modulus")            // 0 means use a random prime
	base := flag.Int("base", 0, "base")                 // 0 means use a random base
	publicKey := flag.Int("publickey", 0, "public key") // 0 means use a random public key
	flag.Parse()

	// Generate a private key
	b.privateKey = b.generatePrivateKey(*mod, *base, *publicKey)
	a.privateKey = a.generatePrivateKey(*mod, *base, *publicKey)

	return 0

}
