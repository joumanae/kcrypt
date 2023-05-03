package diffiehellman

import (
	"flag"
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

var keyExchange struct {
	modulus int
	base    int
	alice   int
	bob     int
}

// generate random secret key
func GenerateSecretKey() int {
	rand.Seed(time.Now().Unix())
	secret := rand.Intn(1000)
	return secret
}

func Power(base int, x int) *big.Int {
	result := big.NewInt(1)
	for i := 0; i < x; i++ {
		result.Mul(result, big.NewInt(int64(base)))
	}
	return result
}

// A=g^a mod p
func CalculatePublicNumber(base int, secret int, modulus int) *big.Int {
	p := Power(base, secret)
	p.Mod(p, big.NewInt(int64(modulus))) // p = p % modulus
	return p
}

func CalculatePrivateKey(publicNumber *big.Int, secret int, modulus int) *big.Int {
	b := keyExchange.base
	pn := CalculatePublicNumber(b, secret, modulus)
	p := Power(int(pn.Int64()), secret)
	p.Mod(p, big.NewInt(int64(modulus)))
	return p
}

func Main() int {
	// var b bob
	// var a alice
	mod := flag.Int("mod", 0, "mod")    // 0 means use a random prime
	base := flag.Int("base", 0, "base") // 0 means use a random base

	flag.Parse()

	keyExchange.modulus = *mod
	keyExchange.base = *base
	keyExchange.alice = GenerateSecretKey()
	keyExchange.bob = GenerateSecretKey()
	PNumberA := CalculatePublicNumber(keyExchange.base, keyExchange.alice, keyExchange.modulus)
	fmt.Println(CalculatePrivateKey(PNumberA, keyExchange.alice, keyExchange.modulus))
	PNumberB := CalculatePublicNumber(keyExchange.base, keyExchange.bob, keyExchange.modulus)
	fmt.Println(CalculatePrivateKey(PNumberB, keyExchange.bob, keyExchange.modulus))

	return 0
}
