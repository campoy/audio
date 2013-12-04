audio
=====

_Highly volatile: currently working on this_

This package will be growing as I experiment with audio processing/playing in Go.

For now there's two subpackages that provide some basic functionalities.

* github.com/campoy/audio/notes

  This package provides parsing of notes from English format (e.g. A#4) to Notes, which provide the Freq method, returning the frequence. 
 
* github.com/campoy/audio/audio

  This package provides the Audio type that allows you to play a sound given a Sampler, which is called every time a new sample is needed.
