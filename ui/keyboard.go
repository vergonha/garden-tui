package ui

import (
	"time"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/gdamore/tcell/v2"
	services "github.com/vergonha/garden-tui/services"
	garden "github.com/vergonha/garden-tui/services/garden"
	player "github.com/vergonha/garden-tui/services/player"
)

// Definitely isn't a correct name for this struct, but i'll use it to avoid making unnecessary API calls.
type Cache struct {
	CurrentResults garden.Search
}

// SetupKeyboardInputHandlers sets up the keyboard input handlers for the UI.
// It captures keyboard events and performs actions based on the key pressed.
func SetupKeyboardInputHandlers(UI *UI, Service *services.Service) {
	// Initialize a new Cache to store the current search results
	state := Cache{}

	// Set the input capture function for the application
	UI.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// If 's' is pressed, set focus on the search form
		if event.Rune() == 's' {
			if UI.App.GetFocus() != UI.Station.Form.Search {
				UI.SetFocusOnSearch()
			}
			return event
		}

		// If 'p' is pressed, pause or play the current audio
		if event.Rune() == 'p' {
			if Service.Player != nil && UI.App.GetFocus() != UI.Station.Form.Search {
				Service.Player.Pause()
				if Service.Player.IsPaused() {
					UI.NowPlaying.State.CurrentStation.SetText(" ⏸️ " + UI.NowPlaying.State.Title)
				} else {
					UI.NowPlaying.State.CurrentStation.SetText(" 🎶 " + UI.NowPlaying.State.Title)
				}
			}
		}

		// If Esc is pressed, remove focus from all elements
		if event.Key() == tcell.KeyEsc {
			UI.App.SetFocus(nil)
		}

		// If Enter is pressed, perform a search or play a selected audio
		if event.Key() == tcell.KeyEnter {
			// If focus is on the search form, perform a search
			if UI.App.GetFocus() == UI.Station.Form.Search {
				state.CurrentResults = Service.API.Search(UI.Station.Form.Search.GetText())
				UI.DisplayAndSetFocusOnResults(&state.CurrentResults)
				return event
			}

			// If focus is on the results list, play the selected audio
			if UI.App.GetFocus() == UI.Station.Results.Results {
				idx := UI.Station.Results.Results.GetCurrentItem()
				selected := state.CurrentResults.Hits.Hits[idx].Source.URL
				stream := Service.API.Stream(selected)

				// Decode the MP3 stream and initialize the speaker
				streamer, format, _ := mp3.Decode(stream)
				speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
				ap := player.NewAudioPanel(format.SampleRate, streamer)
				Service.Player = ap
				Service.Player.Play()

				// Update the Now Playing information
				UI.NowPlaying.State.Title = state.CurrentResults.Hits.Hits[idx].Source.Title + " - " + state.CurrentResults.Hits.Hits[idx].Source.Subtitle
				UI.NowPlaying.State.CurrentStation.SetText(" 🎶 " + state.CurrentResults.Hits.Hits[idx].Source.Title + " - " + state.CurrentResults.Hits.Hits[idx].Source.Subtitle)
			}
		}

		// Return the event to allow further processing
		return event
	})
}
