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
		message: []byte("GWSGGSDOKV"),
		want:    []byte("HELLOWORLD"),
	},
	{
		key:     []byte("BITFIELD"),
		message: []byte("GWSGG SDOKV"),
		want:    []byte("HELLO WORLD"),
	},
}

func TestShift(t *testing.T) {

	t.Parallel()
	for _, tt := range decipherTableTest {
		//v.Shift
		got := vigenere.NewVigenere(string(tt.key)).Shift(tt.message)
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
		want:    []byte("GWSGGSDOKV"),
	},
	{
		key: []byte("BITFIELD"),

		message: []byte("HELLO WORLD"),
		want:    []byte("GWSGG SDOKV"),
	},
}

func TestUnshift(t *testing.T) {

	t.Parallel()
	for _, tt := range cipherTableTest {
		//v.Shift
		got := vigenere.NewVigenere(string(tt.key)).Unshift(tt.message)
		if !cmp.Equal(got, tt.want) {
			t.Errorf("got %q, want %q", got, tt.want)
		}
	}
}
