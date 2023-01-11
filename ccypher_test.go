package ccypher

import (
	"testing"
)

var testCases = []struct {
	r        rune
	shift    int
	expected rune
}{
	{'a', 1, 'b'},
	{'z', 1, 'a'},
	{'a', 26, 'a'},
	{'z', 26, 'z'},
}

func TestShift(t *testing.T) {
	for _, tc := range testCases {
		if ShiftRune(tc.r, tc.shift) != tc.expected {
			t.Errorf("The rune is incorrect %v", string(ShiftRune(tc.r, tc.shift)))
		}

	}
}
