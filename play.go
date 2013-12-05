package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/campoy/audio/audio"
	"github.com/campoy/audio/note"
)

const sampleRate = 44100

func playNotes(ins audio.Instrument) error {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		for _, f := range strings.Fields(s.Text()) {
			n, err := note.FromString(f)
			if err != nil {
				fmt.Println("invalid note: ", f)
			}
			ins.SetFreq(n.Freq())
			time.Sleep(250 * time.Millisecond)
		}
	}
	return s.Err()
}

func main() {
	audio.Initialize()
	defer audio.Terminate()

	s := audio.NewSine(0, sampleRate)

	go func() {
		defer s.Stop()

		err := playNotes(s)
		if err != nil {
			panic(err)
		}
	}()

	err := s.Play()
	if err != nil {
		panic(err)
	}

	fmt.Println("good bye!")
}
