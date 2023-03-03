package ccipher

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"unicode"
)

type Cipher struct {
	Key int
}

func (c *Cipher) Encipher(s string) string {
	return Transform(s, c.Key)
}

func (c *Cipher) Decipher(s string) string {
	return Transform(s, -c.Key)
}

func Transform(s string, key int) string {
	var b bytes.Buffer
	for _, r := range s {
		r = ShiftRune(r, key)
		// fmt.Printf("Enciphered: %d bytes", enc)
		b.WriteRune(r)
	}
	return b.String()
}

func ShiftRune(r rune, Shift int) rune {
	if !unicode.IsLetter(r) {
		return r
	}

	r = unicode.ToUpper(r)

	if r >= 'A' && r <= 'Z' {
		r += rune(Shift)
		if r > 'Z' {
			r -= 26
		}
	}
	return r
}

func New(key int) *Cipher {
	return &Cipher{
		Key: key,
	}
}

const DefaultKey = 13

func Main() int {
	// encipher -k 1 message.txt
	key := flag.Int("k", DefaultKey, "the key to encipher/decipher with")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s -k KEY PATH\nEnciphers a given file with a given key\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if len(flag.Args()) < 1 {
		flag.Usage()
		return 1
	}
	data, err := os.ReadFile(flag.Args()[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	fmt.Println(New(*key).Encipher(string(data)))
	return 0
}
