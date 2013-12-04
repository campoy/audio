package note

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
		{"C0", 0, nil},
		{"D0", 2, nil},
		{"E0", 4, nil},
		{"F0", 5, nil},
		{"G0", 7, nil},
		{"A0", 9, nil},
		{"B0", 11, nil},
		{"H0", 0, errors.New("unknown note name H")},

		// Note octaves
		{"C0", 0, nil},
		{"C1", 12, nil},
		{"C2", 24, nil},
		{"C3", 36, nil},
		{"C-1", 0, errors.New("lowest note is C0")},
		{"Coctave", 0, errors.New("octave \"octave\" is not a number")},
		{"C", 0, errors.New("missing note octave")},

		// Accidentals
		{"Cb", 0, errors.New("missing note octave")},
		{"Cb0", 0, errors.New("lowest note is C0")},
		{"C&0", 0, errors.New("lowest note is C0")},
		{"C#0", 1, nil},
		{"Db0", 1, nil},
		{"D#0", 3, nil},
		{"Cb1", 11, nil},
		{"C&1", 10, nil},
		{"B#0", 12, nil},
		{"Bx0", 13, nil},
	}

	for _, c := range cases {
		n, err := FromString(c.text)
		if !sameError(c.err, err) {
			t.Errorf("expected error %v; got %v", c.err, err)
			continue
		}
		if c.note != n {
			t.Errorf("expected note %v; got %v", c.note, n)
		}
	}
}
