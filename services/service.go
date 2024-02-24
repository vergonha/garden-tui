package services

import (
	"io"

	"github.com/faiface/beep"
	garden "github.com/vergonha/garden-tui/services/garden"
)

type Provider interface {
	Search(place string) garden.Search
	Stream(address string) io.ReadCloser
}

type Player interface {
	Play()
	VolumeUp()
	VolumeDown()
	GetVolume() float64
	Pause()
	IsPaused() bool
	SetStreamer(streamer beep.StreamSeeker)
}

type Service struct {
	API    Provider
	Player Player
}
