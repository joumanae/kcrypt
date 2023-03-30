package vigenere_test

import (
	"fmt"
	"testing"
	"vigenere"
)

func TestVigenere(t *testing.T) {
	fmt.Println("TestVigenere")
}

var byteTest = []struct {
	key     []byte
	message []byte
	want    string
}{
	{[]byte("GO"), []byte("RSSCT"), "LEMON"},
	{[]byte("GO"), []byte("NSRZU KUFRR"), "HELLO WORLD"},
	{[]byte("GO"), []byte("YIYVO"), "SUSHI"},
}

func TestShift(t *testing.T) {

	t.Parallel()
	for _, tt := range byteTest {
		//v.Shift
		got := vigenere.NewVigenere(string(tt.key)).Shift(tt.message)
		if got != tt.want {
			t.Errorf("got %q, want %q", got, tt.want)

		}
	}
}
