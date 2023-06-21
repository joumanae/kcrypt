package kcrypt

import (
	"errors"
	"flag"
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"strings"
)

// key word: vigenere
type Substitute struct {
	key []byte
}

// Return the Substitute key
func NewSubstitutionCipher(key string) *Substitute {
	return &Substitute{[]byte(key)}
}

func (v *Substitute) Cipher(message []byte) []byte {
	k := v.key
	fmt.Println("key", k, "message", message)
	shift := make([]byte, len(message))
	message = []byte(strings.ToUpper(string(message)))
	// Repeat the key until it is the same length as the message
	for len(string(k)) < len(message) {
		k = append(k, k...)
	}
	var plain []byte

	// Shift the message by the key
	for i := 0; i < len(message); i++ {

		// skip non-alphabetic characters and update the key
		if message[i] == 32 {
			shift[i] = message[i]
			plain = append(plain, shift[i])
			k = append((k[:i]), k[i-1:]...)
			continue
		}

		// skip punctuation and update the key
		if message[i] < 65 || message[i] > 90 {
			shift[i] = message[i]
			plain = append(plain, shift[i])
			k = append((k[:i]), k[i-1:]...)
			continue
		}

		// the key changes when there is a non-alphabetic character
		shift[i] = message[i] + (k[i] - 65)

		if shift[i] > 90 {
			shift[i] -= 26
		}
		plain = append(plain, shift[i])
	}
	return plain
}

func (v *Substitute) Decipher(message []byte) []byte {
	k := v.key
	shift := make([]byte, len(message))

	// Repeat the key until it is the same length as the message
	for len(string(k)) < len(message) {
		k = append(k, k...)
	}
	var plain []byte

	// Shift the message by the key
	for i := 0; i < len(message); i++ {

		// skip non-alphabetic characters and update the key
		if message[i] == 32 {
			shift[i] = message[i]
			plain = append(plain, shift[i])
			k = append((k[:i]), k[i-1:]...)
			continue
		}
		// the key changes when there is a non-alphabetic character
		shift[i] = message[i] - (k[i] - 65)

		if shift[i] < 65 {
			shift[i] += 26
		}
		plain = append(plain, shift[i])
	}
	return plain
}

var ErrZeroModulus = errors.New("the modulus cannot be 0")
var ErrZeroBase = errors.New("the base cannot be 0")

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
func PublicKey(base int, modulus int, secretKey int) (*big.Int, error) {
	if modulus == 0 {
		return nil, ErrZeroModulus
	}
	if base == 0 {
		return nil, ErrZeroBase
	}

	p := Power(big.NewInt(int64(base)), secretKey)
	p.Mod(p, big.NewInt(int64(modulus)))
	return p, nil
}

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

	const DefaultKey = "GO"
	scipher := flag.Bool("scipher", false, "Runs the substitute cipher code if true")
	dhkeygen := flag.Bool("dhkeygen", false, "Runs the Diffie Hillman key exchange")
	key := flag.String("key", DefaultKey, "key")
	decipher := flag.Bool("d", false, "decipher")
	cipher := flag.Bool("c", false, "cipher")
	mod := flag.Int("modulus", 1, "The modulus is a prime number")
	base := flag.Int("base", 1, "base")
	pubKey := flag.String("publicKey", "", "This is the public key")
	secretKey := GenerateSecretKey()
	secret := flag.Int("secret", 1, "This is your secret key")

	flag.Usage = func() {
		println("Usage: vigenere [-d] [-key key] string")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *scipher {
		message, err := os.ReadFile(flag.Args()[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}

		withKey := NewSubstitutionCipher(*key)
		var output string

		if *decipher {
			output = string(withKey.Decipher([]byte(message)))
		}
		if *cipher {
			output = string(withKey.Cipher([]byte(message)))
		}

		fmt.Println(output)

	}
	if *dhkeygen {
		if len(*pubKey) == 0 {

			pn1, err := PublicKey(*base, *mod, secretKey)
			if err != nil {
				fmt.Println("Modulus cannot be 0")
				os.Exit(1)
			}
			fmt.Printf("This is your public key: %s, & this is your secret key %v.", pn1, secretKey)
		} else {
			pk, ok := ParseBigInt(*pubKey)
			if !ok {
				fmt.Println("Your public key is not valid")
				os.Exit(1)
			}

			sk, err := SharedKey(pk, *secret, *mod)
			if err != nil {
				fmt.Println("Modulus cannot be 0")
				os.Exit(1)
			}
			fmt.Printf("This is your shared key %s", sk)
		}
	}

	return 0
}
