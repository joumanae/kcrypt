package vigenere

import (
	"flag"
	"strings"

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

func (v *Vigenere) Shift(message []byte) string {
	k := v.key
	shift := make([]byte, len(message))

	// Repeat the key until it is the same length as the message
	for len(string(k)) < len(string(message)) {
		fmt.Println(len(string(k)))
		k = append(k, k...)
	}

	var plain []byte

	// Shift the message by the key
	for i := 0; i < len(message); i++ {
		shift[i] = ((message[i] - k[i]) % 66) + 65
		plain = append(plain, shift[i])

	}

	plainMessage := strings.ToUpper(string(plain))
	return plainMessage
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
	// message, err := os.ReadFile(flag.Args()[0])
	// if err != nil {
	// fmt.Fprintln(os.Stderr, err)
	// return 1
	// }

	message := "NSRZUKUFRR"
	cipher := NewVigenere(*key)
	var output string

	if *decipher {
		output = string(cipher.Shift([]byte(message)))
	}

	fmt.Println(output)
	return 0
}
