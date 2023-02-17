package ccypher_test

import (
	"ccypher"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var RuneTest = []struct {
	r     rune
	shift int
	want  rune
}{
	{'a', 1, 'B'},
	{'z', 1, 'A'},
	{'!', 1, '!'},
	{'1', 1, '1'},
	{'A', 2, 'C'},
	{'Z', 1, 'A'},
	{'M', 13, 'Z'},
	{'Z', 13, 'M'},
	{'Y', 2, 'A'},
}

func TestShiftRune(t *testing.T) {
	t.Parallel() // This is a new feature in Go 1.7

	for _, tt := range RuneTest {
		got := ccypher.ShiftRune(tt.r, tt.shift)
		want := tt.want
		if got != want {
			t.Errorf("With the following input rune %c and shift %d, Expected %c, but got %c", tt.r, tt.shift, want, got)
		}
	}
}

func TestEncipherWithKey1TurnsABCIntoBCD(t *testing.T) {
	t.Parallel()
	want := "BCD"
	got := ccypher.New(1).Encipher("abc")
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestEncipherThenDecipherReproducesOriginalOutput(t *testing.T) {
	t.Parallel()
	want := "HELLO WORLD"
	c := ccypher.New(1)
	ciphertext := c.Encipher(want)
	got := c.Decipher(ciphertext)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
