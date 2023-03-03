package ccipher

import (
	"bytes"
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
