package main

import (
	services "github.com/vergonha/garden-tui/services"
	garden "github.com/vergonha/garden-tui/services/garden"
	"github.com/vergonha/garden-tui/ui"

	"github.com/rivo/tview"
)

func main() {

	app := tview.NewApplication()

	service := services.Service{
		API:    &garden.Garden{},
		Player: nil,
	}

	ui.Run(app, &service)

}
