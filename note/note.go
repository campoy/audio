package note

import (
	"fmt"
	"math"
	"strconv"
)

var (
	notes    = map[int32]byte{0: 'C', 2: 'D', 4: 'E', 5: 'F', 7: 'G', 9: 'A', 11: 'B'}
	notesInv = inverse(notes)

	accidentals    = map[int32]byte{-2: '&', -1: 'b', 1: '#', 2: 'x'}
	accidentalsInv = inverse(accidentals)
)

func inverse(in map[int32]byte) map[byte]int32 {
	out := make(map[byte]int32)
	for k, v := range in {
		out[v] = k
	}
	return out
}

type Note int32

func (n Note) String() string {
	o, n := n/12, n%12
	name, ok := notes[int32(n)]
	acc := byte(0)
	if !ok {
		name, ok = notes[int32(n)-1]
		acc = accidentals[1]
	}
	return fmt.Sprintf("%c%c%d", name, acc, o)
}

func FromString(s string) (Note, error) {
	if len(s) == 0 {
		return 0, fmt.Errorf("empty note")
	}

	// Note name
	n, ok := notesInv[s[0]]
	if !ok {
		return 0, fmt.Errorf("unknown note name %c", s[0])
	}
	s = s[1:]

	// Accidentals
	if len(s) > 0 {
		a, ok := accidentalsInv[s[0]]
		if ok {
			s = s[1:]
			n += a
		}
	}

	// Octave
	if len(s) == 0 {
		return 0, fmt.Errorf("missing note octave")
	}
	o, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("octave %q is not a number", s)
	}
	note := Note(int32(o)*12 + n)
	if note < 0 {
		return 0, fmt.Errorf("lowest note is C0")
	}
	return note, nil
}

func (n Note) Freq() float64 {
	return 440 * math.Pow(2, float64(n-48)/12)
}
