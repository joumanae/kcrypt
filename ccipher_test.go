package ccypher_test

import (
	"ccypher"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestShiftRune_GivenAAnd1ReturnB(t *testing.T) {
	t.Parallel() // This is a new feature in Go 1.7
	result := ccypher.ShiftRune('a', 1)
	expected := 'b'
	if result != expected {
		t.Errorf("The expected rune is %v, when the shift is by %v", expected, 1)
	}
}

func TestEncipherWithKey1TurnsABCIntoBCD(t *testing.T) {
	t.Parallel()
	want := "bcd"
	got := ccypher.New(1).Encipher("abc")
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestEncipherThenDecipherReproducesOriginalOutput(t *testing.T) {
	t.Parallel()
	want := "hello world"
	c := ccypher.New(1)
	ciphertext := c.Encipher(want)
	got := c.Decipher(ciphertext)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestShiftRune_OnlyShiftsLetters(t *testing.T) {
	t.Parallel()
	result := ccypher.ShiftRune('1', 1)
	expected := '1'
	if result != expected {
		t.Errorf("The expected rune is %v, when the shift is by %v", expected, 1)
	}
}

func TestShiftRune_ShiftsBackToA(t *testing.T) {
	t.Parallel()
	got := ccypher.ShiftRune('z', 1)
	want := 'a'
	if got != want {
		t.Errorf("The expected rune is %c, when the shift is by %v", want, 1)
	}
}
