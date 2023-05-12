package vigenere_test

import (
	"testing"
	"vigenere"

	"github.com/google/go-cmp/cmp"
)

var decipherTableTest = []struct {
	key     []byte
	message []byte
	want    []byte
}{

	{
		key:     []byte("BITFIELD"),
		message: []byte("IMEQWAZUML"),
		want:    []byte("HELLOWORLD"),
	},
	{
		key:     []byte("BITFIELD"),
		message: []byte("IMEQW AZUML"),
		want:    []byte("HELLO WORLD"),
	},
}

func TestDecipher(t *testing.T) {

	t.Parallel()
	for _, tt := range decipherTableTest {

		got := vigenere.NewVigenere(string(tt.key)).Decipher(tt.message)
		if !cmp.Equal(got, tt.want) {
			t.Errorf("got %q, want %q", got, tt.want)
		}
	}
}

var cipherTableTest = []struct {
	key     []byte
	message []byte
	want    []byte
}{
	{
		key:     []byte("BITFIELD"),
		message: []byte("HELLOWORLD"),
		want:    []byte("IMEQWAZUML"),
	},
	{
		key: []byte("BITFIELD"),

		message: []byte("HELLO WORLD"),
		want:    []byte("IMEQW AZUML"),
	},
}

func TestCipher(t *testing.T) {

	t.Parallel()
	for _, tt := range cipherTableTest {
		//v.Shift
		got := vigenere.NewVigenere(string(tt.key)).Cipher(tt.message)
		if !cmp.Equal(got, tt.want) {
			t.Errorf("got %q, want %q", got, tt.want)
		}
	}
}
