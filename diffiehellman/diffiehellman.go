package diffiehellman

import (
	"fmt"
	"math/rand"
	"time"
)

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

func (a *alice) generateInteger() int {
	var df diffieHellman
	A := df.base ^ a.privateKey%df.modulus // A = g^a mod p
	return A
}

// B^a mod p
// func (a *alice) generatePrivateKey(modulus int, base int, publicKey int) int {
// 	publicKey = a.publicKey

// 	return 0
// }

// // need to figure out how to send it to alice then bob
// // need to add the tests
// // A^b mod p
// func (b *bob) generatePrivateKey(modulus int, base int, publicKey int) int {
// 	publicKey = b.publicKey

// 	return 0
// }

// Alice sends Integer to Bob
// Alice receives Bob's integer
// Alice computes it
// Same steps for Bob
// Alice computes s = B^a mod p

// Bob computes s = A^b mod p
// Next: add math big
//publicNumber

func Power(base int, x int) int {
	result := 1
	for i := 0; i < x; i++ {
		result *= base
	}
	return result
}

func Main() int {
	// var b bob
	// var a alice
	// mod := flag.Int("modulus", 0, "modulus")            // 0 means use a random prime
	// base := flag.Int("base", 0, "base")                 // 0 means use a random base
	// flag.Parse()
	base := 2
	modulus := 100
	rand.Seed(time.Now().Unix())
	aliceSecret := rand.Intn(10)
	alicePublic := Power(base, aliceSecret) % modulus
	bobSecret := rand.Intn(10)
	bobPublic := Power(base, bobSecret) % modulus
	aliceKey := Power(bobPublic, aliceSecret) % modulus
	bobKey := Power(alicePublic, bobSecret) % modulus
	fmt.Printf("Here's Alice's public number %v. Here's Bob's public number %v\n", alicePublic, bobPublic)
	fmt.Printf("Here's Alice's secret number %v, and here's Bob's %v\n", aliceSecret, bobSecret)
	fmt.Printf("Here's Alice's secret key:%v, here's Bob's %v\n", aliceKey, bobKey)
	return 0

}
