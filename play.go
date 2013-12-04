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

func main() {
	audio.Initialize()
	defer audio.Terminate()

	sine := audio.NewSine(0, sampleRate)
	s := audio.NewSound(sine, sampleRate)

	go func() {
		defer s.Stop()

		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			for _, f := range strings.Fields(sc.Text()) {
				n, err := note.FromString(f)
				if err != nil {
					fmt.Println("invalid note: ", f)
				}
				sine.SetFreq(n.Freq())
				time.Sleep(250 * time.Millisecond)
			}
		}
		if sc.Err() != nil {
			panic(sc.Err())
		}
	}()

	err := s.Play()
	if err != nil {
		panic(err)
	}
	fmt.Println("good bye!")
}
