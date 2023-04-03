package vigenere

import (
	"flag"
	"os"

	"fmt"
)

// key word: vigenere

type Vigenere struct {
	key []byte
}

// NewVigenere returns a new Vigenere
func NewVigenere(key string) *Vigenere {
	return &Vigenere{[]byte(key)}
}

func (v *Vigenere) Shift(message []byte) []byte {
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
		shift[i] = message[i] + (k[i] - 65)

		if shift[i] > 90 {
			shift[i] -= 26
		}
		plain = append(plain, shift[i])
	}
	return plain
}

func (v *Vigenere) Unshift(message []byte) []byte {
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

func Main() int {

	const DefaultKey = "GO"
	key := flag.String("key", DefaultKey, "key")
	decipher := flag.Bool("d", false, "decipher")
	cipher := flag.Bool("c", false, "cipher")

	flag.Usage = func() {
		println("Usage: vigenere [-d] [-key key] string")
		flag.PrintDefaults()
	}

	flag.Parse()
	message, err := os.ReadFile(flag.Args()[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	withKey := NewVigenere(*key)
	var output string

	if *decipher {
		output = string(withKey.Shift([]byte(message)))
	}
	if *cipher {
		output = string(withKey.Unshift([]byte(message)))
	}

	fmt.Println(output)
	return 0
}
