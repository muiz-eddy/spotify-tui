package spotifyUI

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("input", maxX/4, maxY/2-1, maxX*3/4, maxY/2+1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "Input"
		v.Editable = true
		v.Wrap = true
	}
	return nil
}

func Keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("input", gocui.KeyEnter, gocui.ModNone, SubmitInput); err != nil {
		return err
	}

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	return nil
}

func SubmitInput(g *gocui.Gui, v *gocui.View) error {
	input := v.Buffer()
	v.Clear()
	v.SetCursor(0, 0)

	fmt.Println("User Input: ", input)
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
