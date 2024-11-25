package main

import (
	"github.com/jroimartin/gocui"
	"log"
	"spotify-tui/internal/spotifyUI"
)

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetManagerFunc(spotifyUI.Layout)

	g.Update(func(g *gocui.Gui) error {
		_, err := g.SetCurrentView("input")
		return err
	})

	if err := spotifyUI.Keybindings(g); err != nil {
		log.Panicln(err)
	}

	// Start the GUI
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
