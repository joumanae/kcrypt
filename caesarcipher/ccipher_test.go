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
	shift rune
	want  rune
}{
	{'A', 'B', 'C'},
	{'z', 'A', 'A'},
	{'!', 'A', '!'},
	{'1', 'A', '1'},
	{'Z', 'A', 'A'},
	{'M', 'M', 'Z'},
	{'Z', 'M', 'M'},
	{'Y', 'M', 'L'},
	{'A', 'M', 'N'},
	{'Y', 'B', 'A'},
	{'Z', 'N', 'N'},
	{'Y', 'N', 'M'},
	{'B', -'E', 'W'},
	{'D', -'E', 'Y'},
}

func TestShiftRune(t *testing.T) {
	t.Parallel() // This is a new feature in Go 1.7

	for _, tt := range RuneTest {
		got := c.ShiftRune(tt.r, tt.shift)
		want := tt.want
		if got != want {
			t.Errorf("With the following input rune %c and shift %c, Expected %c, but got %c", tt.r, tt.shift, want, got)
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
