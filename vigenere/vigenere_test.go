package vigenere_test

import (
	"fmt"
	"testing"
	"vigenere"

	"github.com/google/go-cmp/cmp"
)

func TestVigenere(t *testing.T) {
	fmt.Println("TestVigenere")
}

var byteTest = []struct {
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
		key:     []byte("BITFIELD"),
		message: []byte("HELLO WORLD"),
		want:    []byte("IMEQW AZUML"),
	},
}

func TestShift(t *testing.T) {

	t.Parallel()
	for _, tt := range byteTest {
		//v.Shift
		got := vigenere.NewVigenere(string(tt.key)).Shift(tt.message)
		if !cmp.Equal(got, tt.want) {
			t.Errorf("got %q, want %q", got, tt.want)
		}
	}
}
