package caesarcipher_test

import (
	"os"
	"testing"

	c "caesarcipher"

	"github.com/google/go-cmp/cmp"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"encipher": c.Main,
	}))
}

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
	{'Y', 13, 'L'},
	{'A', 13, 'N'},
	{'Y', 2, 'A'},
	{'Z', 14, 'N'},
	{'Y', 14, 'M'},
	{'Z', 15, 'O'},
	{'Y', 5, 'D'},
}

func TestShiftRune(t *testing.T) {
	t.Parallel() // This is a new feature in Go 1.7

	for _, tt := range RuneTest {
		got := c.ShiftRune(tt.r, tt.shift)
		want := tt.want
		if got != want {
			t.Errorf("With the following input rune %c and shift %d, Expected %c, but got %c", tt.r, tt.shift, want, got)
		}
	}
}

func TestEncipherWithKey1TurnsABCIntoBCD(t *testing.T) {
	t.Parallel()
	want := "BCD"
	got := c.New(1).Encipher("ABC")
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestEncipherThenDecipherReproducesOriginalOutput(t *testing.T) {
	t.Parallel()
	want := "HELLO WORLD"
	c := c.New(1)
	ciphertext := c.Encipher(want)
	got := c.Decipher(ciphertext)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestScript(t *testing.T) {
	t.Parallel()
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}
