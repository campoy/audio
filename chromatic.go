package main

import (
	"math"
	"time"

	"code.google.com/p/portaudio-go/portaudio"
)

const sampleRate = 44100

var step = math.Pow(2, 1.0/12)

func main() {
	portaudio.Initialize()
	defer portaudio.Terminate()

	// Multiple goroutines playing music, yay!
	Mchord := []float64{0, 4, 7, 12}

	ss := make([]*Sine, 0, len(Mchord))
	for _, f := range Mchord {
		s, err := newSine(220*math.Pow(step, f), sampleRate)
		if err != nil {
			panic(err)
		}
		defer s.End()
		ss = append(ss, s)
	}

	beat := time.Tick(500 * time.Millisecond)
	quit := time.After(10 * time.Second)

	for {
		select {
		case <-quit:
			return
		case <-beat:
			for _, s := range ss {
				s.Transpose(1)
			}
		}
	}
}

type Sine struct {
	str         *portaudio.Stream
	step, phase float64
}

func newSine(freq, sampleRate float64) (*Sine, error) {
	s := &Sine{step: freq / sampleRate}
	str, err := portaudio.OpenDefaultStream(0, 1, sampleRate, 0, s.processAudio)
	if err != nil {
		return nil, err
	}
	if err := str.Start(); err != nil {
		return nil, err
	}
	s.str = str
	return s, nil
}

func (s *Sine) Transpose(n int) {
	s.step *= math.Pow(step, float64(n))
}

func (s *Sine) End() error {
	if err := s.str.Stop(); err != nil {
		return err
	}
	return s.str.Close()
}

func (s *Sine) processAudio(out [][]float32) {
	for i := range out[0] {
		out[0][i] = float32(math.Sin(2 * math.Pi * s.phase))
		_, s.phase = math.Modf(s.phase + s.step)
	}
}
