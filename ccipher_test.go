package ccypher

import (
	"testing"
)

// var testCases = []struct {
// 	r        rune
// 	shift    int
// 	expected rune
// }{
// 	{'a', 1, 'b'},
// 	{'z', 1, 'a'},
// 	{'a', 26, 'a'},
// 	{'z', 26, 'z'},
// }

func TestShiftRune(t *testing.T) {

	// Shift takes a rune and an integer, and shifts the given rune by the given integer and returns the result

	// for _, tc := range testCases {
	// if r != tc.expected {
	// 	t.Errorf("Expected %v, got %v", tc.expected, tc.r)
	// }
	result := ShiftRune('a', 1)
	expected := 'b'
	if result != expected {
		t.Errorf("The expected rune is %v, when the shift is by %v", expected, 1)
	}
}
