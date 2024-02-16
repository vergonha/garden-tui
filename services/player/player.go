package services

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/speaker"
)

// thanks, beep documentation

type Player struct {
	Panel *audioPanel
}

func (p *Player) Play() {
	p.Panel.Play()
}

func (p *Player) Pause() {
	p.Panel.Pause()
}

type audioPanel struct {
	sampleRate beep.SampleRate
	streamer   beep.StreamSeeker
	ctrl       *beep.Ctrl
	resampler  *beep.Resampler
	volume     *effects.Volume
}

func (ap *audioPanel) Play() {
	speaker.Play(ap.volume)
}

func (ap *audioPanel) Pause() {
	speaker.Lock()
	ap.ctrl.Paused = !ap.ctrl.Paused
	speaker.Unlock()
}

func (ap *audioPanel) IsPaused() bool {
	return ap.ctrl.Paused
}

func (ap *audioPanel) SetStreamer(streamer beep.StreamSeeker) {
	ap.streamer = streamer
	ap.ctrl = &beep.Ctrl{Streamer: beep.Loop(-1, streamer)}
	ap.volume = &effects.Volume{Streamer: beep.ResampleRatio(4, 1, ap.ctrl), Base: 2}

}

func NewAudioPanel(sampleRate beep.SampleRate, streamer beep.StreamSeeker) *audioPanel {
	ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer)}
	resampler := beep.ResampleRatio(4, 1, ctrl)
	volume := &effects.Volume{Streamer: resampler, Base: 2}
	return &audioPanel{sampleRate, streamer, ctrl, resampler, volume}
}
