package audio

import (
	"math"

	"code.google.com/p/portaudio-go/portaudio"
)

var (
	Initialize = portaudio.Initialize
	Terminate  = portaudio.Terminate
)

type Instrument interface {
	SetFreq(float64)
	Play() error
	Stop()
}

type sound struct {
	sample     func() []float64
	sampleRate float64
	quit       chan error
}

func newSound(sample func() []float64, sampleRate float64) *sound {
	return &sound{
		sample:     sample,
		sampleRate: sampleRate,
		quit:       make(chan error),
	}
}

func (s *sound) Play() error {
	chans := len(s.sample())
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

func (s *sound) Stop() {
	close(s.quit)
}

func (s *sound) processAudio(out [][]float32) {
	for i := range out[0] {
		for c, s := range s.sample() {
			out[c][i] = float32(s)
		}
	}
}

type phaser struct {
	freq, phase, rate float64
}

func (s *phaser) next() float64 {
	_, s.phase = math.Modf(s.phase + s.freq/s.rate)
	return s.phase
}

func (s *phaser) SetFreq(f float64) { s.freq = f }

// Sine wave sound, this sounds smooth
type Sine struct {
	*sound
	phaser
}

func NewSine(freq, rate float64) Instrument {
	s := &Sine{phaser: phaser{
		freq: freq,
		rate: rate,
	}}
	s.sound = newSound(s.sample, rate)
	return s
}

func (s *Sine) sample() []float64 {
	return []float64{math.Sin(2 * math.Pi * s.next())}
}

// Square wave sound, this sounds strong
type Square struct {
	*sound
	phaser
}

func NewSquare(freq, rate float64) Instrument {
	s := &Square{phaser: phaser{
		freq: freq,
		rate: rate,
	}}
	s.sound = newSound(s.sample, rate)
	return s
}

func (s *Square) sample() []float64 {
	if v := s.next(); v > 0.5 {
		return []float64{1}
	}
	return []float64{-1}
}

// Chainsaw wave sound, this is just ugly
type Saw struct {
	*sound
	phaser
}

func NewSaw(freq, rate float64) Instrument {
	s := &Saw{phaser: phaser{
		freq: freq,
		rate: rate,
	}}
	s.sound = newSound(s.sample, rate)
	return s
}

func (s *Saw) sample() []float64 {
	return []float64{s.next()}
}
