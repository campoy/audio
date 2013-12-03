package main

import (
	"math"
	"time"

	"code.google.com/p/portaudio-go/portaudio"
)

const sampleRate = 44100

func main() {
	portaudio.Initialize()
	defer portaudio.Terminate()

	freq, step := 55., math.Pow(2, 1.0/12)

	// Multiple goroutines playing music, yay!
	go func() {
		s := newSine(freq, sampleRate)
		chk(s.Start())
		for {
			time.Sleep(250 * time.Millisecond)
			s.step *= step
		}
		chk(s.Stop())
		s.Close()
	}()

	go func() {
		s := newSine(freq, sampleRate)
		chk(s.Start())
		for {
			time.Sleep(500 * time.Millisecond)
			s.step *= step
		}
		chk(s.Stop())
		s.Close()
	}()

	<-time.After(10 * time.Second)
}

type sine struct {
	*portaudio.Stream
	step, phase float64
}

func newSine(freq, sampleRate float64) *sine {
	s := &sine{step: freq / sampleRate}
	var err error
	s.Stream, err = portaudio.OpenDefaultStream(0, 1, sampleRate, 0, s.processAudio)
	chk(err)
	return s
}

func (g *sine) processAudio(out [][]float32) {
	for i := range out[0] {
		out[0][i] = float32(math.Sin(2 * math.Pi * g.phase))
		_, g.phase = math.Modf(g.phase + g.step)
	}
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
