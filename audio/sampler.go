package audio

import (
	"math"

	"code.google.com/p/portaudio-go/portaudio"
)

var (
	Initialize = portaudio.Initialize
	Terminate  = portaudio.Terminate
)

type Sampler interface {
	Sample() []float64
}

type Sound struct {
	sampler    Sampler
	sampleRate float64
	quit       chan error
}

func NewSound(sampler Sampler, sampleRate float64) *Sound {
	return &Sound{
		sampler:    sampler,
		sampleRate: sampleRate,
		quit:       make(chan error),
	}
}

func (s *Sound) Play() error {
	chans := len(s.sampler.Sample())
	str, err := portaudio.OpenDefaultStream(0, chans, s.sampleRate, 0, s.processAudio)
	if err != nil {
		return err
	}
	if err = str.Start(); err != nil {
		return err
	}

	<-s.quit
	if err := str.Stop(); err != nil {
		return err
	}
	return str.Close()
}

func (s *Sound) Stop() {
	close(s.quit)
}

func (s *Sound) processAudio(out [][]float32) {
	for i := range out[0] {
		for c, s := range s.sampler.Sample() {
			out[c][i] = float32(s)
		}
	}
}

type Sine struct {
	freq, phase, rate float64
}

func NewSine(freq, rate float64) *Sine {
	return &Sine{freq: freq, rate: rate}
}

func (s *Sine) Sample() []float64 {
	_, s.phase = math.Modf(s.phase + s.freq/s.rate)
	return []float64{math.Sin(2 * math.Pi * s.phase)}
}

func (s *Sine) SetFreq(freq float64) {
	s.freq = freq
}
