package ccypher

import (
	"bytes"
	"fmt"
	"unicode"
)

type Cipher struct {
	Key int
}

func (c *Cipher) Encipher(s string) string {
	var b bytes.Buffer
	var ciphered string
	for _, r := range s {
		r = ShiftRune(r, 1)
		b.WriteRune(r)
		ciphered = b.String()
	}
	return fmt.Sprintf("Here's the ciphered message : %v", ciphered)
}

func ShiftRune(r rune, Shift int) rune {
	Shift %= 26
	if unicode.IsLetter(r) {
		return r + rune(Shift)
	}
	if r >= 'z' {
		return r - 26
	}
	return r
}

func (c *Cipher) Decipher(s string) string {
	var deciphered string
	var b bytes.Buffer
	for _, r := range s {
		r = ShiftRune(r, -1)
		b.WriteRune(r)
		deciphered = b.String()
	}
	return fmt.Sprintf("Here's the deciphered message : %v", deciphered)
}

func New(key int) *Cipher {
	return &Cipher{Key: key}
}
