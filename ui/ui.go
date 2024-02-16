package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/vergonha/garden-tui/services"
	garden "github.com/vergonha/garden-tui/services/garden"
)

type UI struct {
	App        *tview.Application
	Station    Station
	NowPlaying NowPlaying
}

func newPrimitive(title string) *tview.TextView {
	view := tview.NewTextView().SetTextAlign(tview.AlignCenter).SetText(title)
	return view
}

func drawSearchGrid() (*tview.Grid, *tview.InputField, *tview.List) {
	// Function to Draw the main search box in the middle of screen.
	drawSearchBox := func() *tview.InputField {
		searchGridFormInputField := tview.NewInputField().
			SetLabel("🔎 ").
			SetPlaceholderTextColor(tcell.ColorWhite).
			SetLabelColor(tcell.ColorWhite)

		return searchGridFormInputField
	}

	search := newPrimitive("Garden Radio 🌎")

	// Create a grid to hold the title, search box and the list of results.
	searchGrid := tview.NewGrid().
		// 2 - 2 - 0
		// 2 rows for title, 2 rows for search box and the rest of div for the list.
		SetRows(2, 2, 0).
		// 0 - 32 - 0
		// 32 columns of space in the middle of screen, to centralize the search stuff.
		SetColumns(0, 32, 0)

	// Add the title to the grid.
	searchGrid.AddItem(search, 0, 1, 1, 1, 0, 0, false)

	// The list of results.
	searchGridFormList := tview.NewList()
	// Add the list of results to grid.
	searchGrid.AddItem(searchGridFormList, 2, 1, 1, 1, 0, 0, false)

	// The search box.
	searchGridFormInputField := drawSearchBox()
	// Add the search box to grid.
	searchGrid.AddItem(searchGridFormInputField, 1, 1, 1, 1, 0, 0, true)
	return searchGrid, searchGridFormInputField, searchGridFormList
}

func drawPlayerGrid() (*tview.Grid, *tview.TextView) {

	// The initial texts for the player.
	playerText := newPrimitive("🎶 Waiting...")
	instructions := newPrimitive("🔍 Search for a station with 's' or play/pause with 'p'")

	// All centered.
	playerGrid := tview.NewGrid().
		SetRows(0).
		SetColumns(0)

	playerGrid.AddItem(playerText, 0, 0, 1, 3, 0, 0, false)
	playerGrid.AddItem(instructions, 1, 0, 1, 3, 0, 0, false)

	return playerGrid, playerText
}

// Run starts the UI
func Run(app *tview.Application, service *services.Service) {

	// Main grid.
	grid := tview.NewGrid().
		// 0 - 2
		// All rows for the search grid and 2 last rows for the player.
		SetRows(0, 2).
		// No columns.
		SetColumns(0).
		SetBorders(true)

	searchGrid, searchGridFormInputField, searchGridFormList := drawSearchGrid()
	playerGrid, playerText := drawPlayerGrid()

	grid.AddItem(playerGrid, 1, 0, 1, 3, 0, 0, false)
	grid.AddItem(searchGrid, 0, 0, 1, 3, 0, 0, true)

	core := UI{
		App: app,
		Station: Station{
			Form: StationForm{
				Search: searchGridFormInputField,
			},
			Results: StationResults{
				Results: searchGridFormList,
			},
		},
		NowPlaying: NowPlaying{
			State: NowPlayingState{
				CurrentStation: playerText,
			},
		},
	}

	SetupKeyboardInputHandlers(&core, service)

	if err := app.SetRoot(grid, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}

func (u *UI) SetFocusOnSearch() {
	u.App.SetFocus(u.Station.Form.Search)
	u.Station.Form.Search.SetText("")
}

func (UI *UI) DisplayAndSetFocusOnResults(currentResults *garden.Search) {
	UI.Station.Form.Search.SetText("")

	UI.Station.Results.Results.Clear()
	for i := range currentResults.Hits.Hits {
		idx := len(currentResults.Hits.Hits) - 1 - i
		UI.Station.Results.Results.InsertItem(0, currentResults.Hits.Hits[idx].Source.Title, currentResults.Hits.Hits[idx].Source.Subtitle, '📻', nil).
			SetShortcutColor(tcell.ColorWhite)
	}

	UI.App.SetFocus(UI.Station.Results.Results.SetCurrentItem(0))
}

type StationForm struct {
	Search *tview.InputField
}

type StationResults struct {
	Results *tview.List
}

type Station struct {
	Form    StationForm
	Results StationResults
}

type NowPlayingState struct {
	CurrentStation *tview.TextView
	Title          string
}

type NowPlaying struct {
	State NowPlayingState
}
