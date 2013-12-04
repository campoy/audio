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

func playNotes(sine *audio.Sine) error {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		for _, f := range strings.Fields(s.Text()) {
			n, err := note.FromString(f)
			if err != nil {
				fmt.Println("invalid note: ", f)
			}
			sine.SetFreq(n.Freq())
			time.Sleep(250 * time.Millisecond)
		}
	}
	return s.Err()
}

func main() {
	audio.Initialize()
	defer audio.Terminate()

	sine := audio.NewSine(0, sampleRate)
	s := audio.NewSound(sine, sampleRate)

	go func() {
		defer s.Stop()

		err := playNotes(sine)
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
