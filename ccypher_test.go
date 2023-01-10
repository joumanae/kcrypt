package ccypher

import "testing"

func TestShift(t *testing.T) {
	if ShiftRune('a', 1) != 'b' {
		t.Errorf("The rune is incorrect %v", ShiftRune('a', 1))
	}
	if ShiftRune('z', 1) != 'a' {
		t.Errorf("The rune is incorrect %v", ShiftRune('z', 1))
	}
}
