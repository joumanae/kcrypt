package ccypher

import (
	"fmt"
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

var invalidShift = []struct {
	r        rune
	shift    int
	expected error
}{
	{'a', -1, fmt.Errorf("The shift number has to be between 0 and 26:%v", -1)},
	{'z', 27, fmt.Errorf("The shift number has to be between 0 and 26:%v", 27)},
}

func TestShift(t *testing.T) {

	for _, is := range invalidShift {
		_, err := ShiftRune(is.r, is.shift)
		if err != is.expected {
			t.Errorf("Expected %v, got %v", is.expected, err)
		}
	}

	for _, tc := range testCases {
		if tc.r != tc.expected {
			t.Errorf("Expected %v, got %v", tc.expected, tc.r)
		}
	}
}
