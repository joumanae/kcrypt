package vigenere

import (
	"flag"
	"fmt"
	"os"
	"unicode"
)

// key word: vigenere
type Vigenere struct {
	key string
}

// NewVigenere returns a new Vigenere
func NewVigenere(key string) *Vigenere {
	return &Vigenere{key}
}

// Encrypt encrypts a string
func (v *Vigenere) EncryptMessage(message string) string {
	var cipher string
	for i, c := range message {
		if unicode.IsLetter(c) && unicode.IsUpper(c) && c+rune(v.key[i%len(v.key)]) > 'Z' {
			fmt.Println(c, rune(v.key[i%len(v.key)]), c+rune(v.key[i%len(v.key)]))
			cipher += string(c + rune(v.key[i%len(v.key)]) - 26)
		}
	}
	return cipher
}

// Decrypt decrypts a string
func (v *Vigenere) DecryptMessage(message string) string {
	var plain string
	for i, c := range message {
		if unicode.IsLetter(c) && unicode.IsUpper(c) && c-rune(v.key[i%len(v.key)]) < 'A' {
			plain += string(c - rune(v.key[i%len(v.key)]))
		}
	}
	return plain
}

func Main() int {
	const DefaultKey = "GO"
	key := flag.String("key", DefaultKey, "key")
	decipher := flag.Bool("d", false, "decipher")
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
	cipher := NewVigenere(*key)
	var output string
	if *decipher {
		output = cipher.DecryptMessage(string(message))
	} else {
		output = cipher.EncryptMessage(string(message))
	}
	fmt.Println(output)
	return 0
}
