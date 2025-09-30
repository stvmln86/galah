package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/stvmln86/galah/galah/games/grid"
	"github.com/stvmln86/galah/galah/games/tile"
	"github.com/stvmln86/galah/galah/maths/pair"
	"github.com/stvmln86/galah/galah/maths/plot"
	"github.com/stvmln86/galah/galah/nodes/wall"
	"github.com/stvmln86/galah/galah/terms/term"
)

func main() {
	// initialise globals
	size := 15
	grid := grid.New(size)
	term, err := term.New()
	if err != nil {
		panic(err)
	}

	// draw border walls on grid
	for _, pair := range plot.Rectangle(pair.New(0, 0), pair.New(size-1, size-1)) {
		grid.Set(pair, tile.New(" DD", wall.New("·WD")))
	}

loop:
	for {
		// calculate pairs
		full := term.Size()
		orig := pair.New((full.X-28)/2, (full.Y-13)/2)
		head := orig.AddXY(0, -1)
		foot := head.AddXY(0, 16)

		// draw grid
		term.Clear()
		term.SetRender(grid.Render())
		term.SetStringPad(head, "header", size*2, " BD")
		term.SetStringPad(foot, "footer", size*2, " GD")
		term.Show()

		// run event poll
		switch evnt := term.Poll().(type) {
		case *tcell.EventKey:
			switch evnt.Key() {
			case tcell.KeyEscape:
				break loop
			}

		case *tcell.EventResize:
			continue loop
		}
	}

	// close screen
	term.Clear()
	term.Close()
}
