package ccypher

import (
	"bytes"
	"unicode"
)

type Cipher struct {
	Key int
}

func (c *Cipher) Encipher(s string) string {
	var b bytes.Buffer
	var ciphered string
	for _, r := range s {
		r = ShiftRune(r, c.Key)
		b.WriteRune(r)
		ciphered = b.String()
	}
	return ciphered
}

func ShiftRune(r rune, Shift int) rune {
	if !unicode.IsLetter(r) {
		return r
	} else {
		r = unicode.ToUpper(r)
	}
	if r+rune(Shift) >= 'Z' && Shift != 13 {
		return r + rune(Shift) - 26
	}
	if r > 'M' && Shift == 13 {
		return r - 13
	}
	if r == 'M' && Shift == 13 {
		return 'Z'
	}
	return r + rune(Shift)
}

func (c *Cipher) Decipher(s string) string {
	var deciphered string
	var b bytes.Buffer
	for _, r := range s {
		r = ShiftRune(r, -c.Key)
		b.WriteRune(r)
		deciphered = b.String()
	}
	return deciphered
}

func New(key int) *Cipher {
	return &Cipher{Key: key}
}
