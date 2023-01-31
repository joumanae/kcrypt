package ccypher

type Cipher struct {
	Key         int
	PlainText   string
	EncodedText string
}

type CipherInterface interface {
	CipherText(s string) []string
	DecipherText(s string) []string
}

func (c *Cipher) CipherText(s string) []string {
	var ciphered []string
	for _, r := range s {
		r = ShiftRune(r, 1)
		ciphered = append(ciphered, string(r))
	}
	return ciphered
}

func ShiftRune(r rune, Shift int) rune {
	return r + rune(Shift)
}

func (c *Cipher) DecipherText(s string) []string {
	var deciphered []string
	for _, r := range s {
		r = ShiftRune(r, -1)
		deciphered = append(deciphered, string(r))
	}
	return deciphered
}

func NewCipher(key int, plainText string) *Cipher {
	return &Cipher{Key: key, PlainText: plainText}
}

func NewDecipher(key int, encodedText string) *Cipher {
	return &Cipher{Key: key, EncodedText: encodedText}
}
