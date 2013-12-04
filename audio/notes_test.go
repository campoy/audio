package audio

import (
	"errors"
	"testing"
)

func sameError(a, b error) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == nil {
		return true
	}
	return a.Error() == b.Error()
}

func TestParseNote(t *testing.T) {
	cases := []struct {
		text string
		note Note
		err  error
	}{
		{"", 0, errors.New("empty note")},

		// Note names
		{"A0", 0, nil},
		{"B0", 2, nil},
		{"C0", 3, nil},
		{"D0", 5, nil},
		{"E0", 7, nil},
		{"F0", 8, nil},
		{"G0", 10, nil},
		{"H0", 0, errors.New("unknown note name H")},

		// Note octaves
		{"A0", 0, nil},
		{"A1", 12, nil},
		{"A10", 120, nil},
		{"A100", 1200, nil},
		{"A-1", 0, errors.New("lowest note is A0")},
		{"Aoctave", 0, errors.New("octave \"octave\" is not a number")},
		{"A", 0, errors.New("missing note octave")},

		// Accidentals
		{"Ab", 0, errors.New("missing note octave")},
		{"Ab0", 0, errors.New("lowest note is A0")},
		{"A&0", 0, errors.New("lowest note is A0")},
		{"A#0", 1, nil},
		{"Bb0", 1, nil},
		{"B#0", 3, nil},
		{"Cb0", 2, nil},
		{"E#0", 8, nil},
		{"Fb0", 7, nil},
		{"Fx0", 10, nil},
		{"A&1", 10, nil},
	}

	for _, c := range cases {
		n, err := ParseNote(c.text)
		if !sameError(c.err, err) {
			t.Errorf("expected error %v; got %v", c.err, err)
			continue
		}
		if c.note != n {
			t.Errorf("expected note %v; got %v", c.note, n)
		}
	}
}
