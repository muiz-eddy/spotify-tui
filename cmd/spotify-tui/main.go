package main

import (
	"fmt"
	/* "github.com/jroimartin/gocui" */
	"spotify-tui/internal/services"
	/* "spotify-tui/internal/spotifyUI" */)

// TODO: resp is returning bytes, change that to string
func main() {
	fmt.Println(services.GenerateToken())
	resp, err := services.Search("taylor swift", "album")
	fmt.Println(err)
	fmt.Println(resp.Albums)
	// g, err := gocui.NewGui(gocui.OutputNormal)
	// if err != nil {
	// 	log.Panicln(err)
	// }
	// defer g.Close()
	//
	// g.SetManagerFunc(spotifyUI.Layout)
	//
	// g.Update(func(g *gocui.Gui) error {
	// 	_, err := g.SetCurrentView("input")
	// 	return err
	// })
	//
	// if err := spotifyUI.Keybindings(g); err != nil {
	// 	log.Panicln(err)
	// }
	//
	// // Start the GUI
	// if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
	// 	log.Panicln(err)
	// }
}
