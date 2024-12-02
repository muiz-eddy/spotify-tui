package main

import (
	"fmt"
	"log"
	"spotify-tui/internal/services"
	"spotify-tui/internal/spotifyUI"

	"github.com/jroimartin/gocui"
)

func main() {
	fmt.Println(services.GenerateToken())
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
