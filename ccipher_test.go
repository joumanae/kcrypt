package ccypher

import (
	"testing"
)

func TestShiftRune(t *testing.T) {
	result := ShiftRune('a', 1)
	expected := 'b'
	if result != expected {
		t.Errorf("The expected rune is %v, when the shift is by %v", expected, 1)
	}
}

func TestCipherText(t *testing.T) {
	result := NewCipher(1, "abc").CipherText("abc")
	if result[0] != "b" || result[1] != "c" || result[2] != "d" {
		t.Errorf("The expected ciphered text is %v, when the shift is by %v", "b", 1)
	}
}

func TestDecipherText(t *testing.T) {
	result := NewDecipher(1, "bcd").DecipherText("bcd")
	if result[0] != "a" || result[1] != "b" || result[2] != "c" {
		t.Errorf("The expected deciphered text is %v, when the shift is by %v", "a", 1)
	}
}
