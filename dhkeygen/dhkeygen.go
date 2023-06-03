package dhkeygen

import (
	"errors"
	"flag"
	"fmt"
	"math/big"
	"math/rand"
	"os"
)

var ErrZeroModulus = errors.New("the modulus cannot be 0")
var ErrZeroBase = errors.New("the base cannot be 0")
var SecretKey struct {
	key int
}

// Generate a random secret key
func GenerateSecretKey() int {
	secret := rand.Intn(1000) + 1
	return secret
}

// Calculate the power of a number and return a big int
func Power(base *big.Int, x int) *big.Int {
	result := big.NewInt(1)
	for i := 0; i < x; i++ {
		result.Mul(result, base)
	}
	return result
}

// Parse a string and return a big int
func ParseBigInt(s string) (*big.Int, bool) {
	n := new(big.Int)
	return n.SetString(s, 10)
}

// Calculate the public key
func PublicKey(base int, modulus int) (*big.Int, error) {
	if modulus == 0 {
		return nil, ErrZeroModulus
	}
	if base == 0 {
		return nil, ErrZeroBase
	}
	SecretKey.key = GenerateSecretKey()
	fmt.Printf("Reminder, your generated secret key is %d\n", SecretKey.key)
	p := Power(big.NewInt(int64(base)), SecretKey.key)
	p.Mod(p, big.NewInt(int64(modulus)))
	return p, nil
}

// TODO:
// There should be a generated shared secret keys that the user can use twice
// Once in the public key,
// Then in the private key
// It needs to be the same and used twice.

// Calculate the shared key
func SharedKey(publicKey *big.Int, secret int, modulus int) (*big.Int, error) {
	if modulus == 0 {
		return nil, ErrZeroModulus
	}

	p := Power(publicKey, secret)
	p = p.Mod(p, big.NewInt(int64(modulus)))
	return p, nil
}

func Main() int {

	mod := flag.Int("m", 1, "The modulus is a prime number")
	base := flag.Int("b", 1, "base")
	pubKey := flag.String("k", "", "This is the public key")

	flag.Parse()

	if len(*pubKey) == 0 {

		pn1, err := PublicKey(*base, *mod)
		if err != nil {
			fmt.Println("Modulus cannot be 0")
			os.Exit(1)
		}
		fmt.Printf("This is your public key: %s.", pn1)
	} else {
		pk, ok := ParseBigInt(*pubKey)
		if !ok {
			fmt.Println("Your public key is not valid")
			os.Exit(1)
		}

		sk, err := SharedKey(pk, SecretKey.key, *mod)
		if err != nil {
			fmt.Println("Modulus cannot be 0")
			os.Exit(1)
		}
		fmt.Printf("This is your shared key %s", sk)
	}
	return 0
}
