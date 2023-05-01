package diffiehellman

import (
	"flag"
	"math/big"
)

var keyExchange struct {
	modulus int
	base    int
	alice   int
	bob     int
}

func Power(base int, x int) *big.Int {
	result := big.NewInt(1)
	for i := 0; i < x; i++ {
		result.Mul(result, big.NewInt(int64(base)))
	}
	return result
}

func CalculatePublicNumber(base int, secret int, modulus int) *big.Int {
	p := Power(base, secret)
	p.Mod(p, big.NewInt(int64(modulus)))
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
	mod := flag.Int("modulus", 0, "modulus") // 0 means use a random prime
	base := flag.Int("base", 0, "base")      // 0 means use a random base
	ak := flag.Int("alicekey", 0, "alice key")
	bk := flag.Int("bobkey", 0, "bob key")
	flag.Parse()

	keyExchange.modulus = *mod
	keyExchange.base = *base
	keyExchange.alice = *ak
	keyExchange.bob = *bk
	PNumberA := CalculatePublicNumber(keyExchange.base, keyExchange.alice, keyExchange.modulus)
	CalculatePrivateKey(PNumberA, keyExchange.alice, keyExchange.modulus)
	PNumberB := CalculatePublicNumber(keyExchange.base, keyExchange.bob, keyExchange.modulus)
	CalculatePrivateKey(PNumberB, keyExchange.bob, keyExchange.modulus)

	return 0
}
