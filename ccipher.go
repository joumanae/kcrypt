package ccipher

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"unicode"
)

type Cipher struct {
	Key      int
	Encipher bool
	Decipher bool
	file     *os.File
}

func (c *Cipher) EncipherDecipher(s string) string {
	var b bytes.Buffer
	var message string
	c.file, _ = os.Open("message.txt")
	defer c.file.Close()
	for _, r := range s {
		if c.Encipher {
			r = ShiftRune(r, c.Key)
			b.WriteRune(r)
			message = b.String()
			w := bufio.NewWriter(c.file)
			enc, err := w.WriteString(message)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("Enciphered: %d bytes", enc)
		}
		if c.Decipher {
			r = ShiftRune(r, -c.Key)
			b.WriteRune(r)
			message = b.String()
			w := bufio.NewWriter(c.file)
			dec, err := w.WriteString(message)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("Deciphered: %d bytes", dec)
		}
	}
	return message
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
		Key:      key,
		Encipher: true,
		Decipher: true,
		file:     nil,
	}
}
