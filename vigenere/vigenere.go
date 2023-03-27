package vigenere

import (
	"flag"
	"fmt"
	"os"
	"strings"
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

func (v *Vigenere) DecipherLetter(message string) rune {
	const asciiA rune = 65

	messageIndex := len(message) - len(v.key)
	asciiLetter := (messageIndex+26)%26 + int(asciiA)

	return rune(asciiLetter)
}

// Decrypt decrypts a string
func (v *Vigenere) DecryptMessage(message string) string {
	var plain string
	newMessage := strings.ToUpper(message)
	for i, c := range newMessage {
		if unicode.IsLetter(c) && unicode.IsUpper(c) && c-rune(v.key[i%len(v.key)]) < 'A' {
			plain += string(v.DecipherLetter(newMessage))
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
	}
	fmt.Println(output)
	return 0
}
