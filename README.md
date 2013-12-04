audio
=====

_Highly volatile: currently working on this_

This package will be growing as I experiment with audio processing/playing in Go.

For now there's two subpackages that provide some basic functionalities.

<h3>github.com/campoy/audio/notes</h3>

This package provides parsing of notes from English format (e.g. A#4) to Notes.
Notes provide the Freq method, returning the frequence. 

<h3>github.com/campoy/audio/audio</h3>

This package provides the Audio type that allows you to play a sound given a Sampler,
which is called every time a new sample is needed.

<h2>Try it now</h2>

Just 'go get' the package and enjoy the beautiful music!

<pre>
  go get github.com/campoy/audio
  go run play.go < music.sample
</pre>

