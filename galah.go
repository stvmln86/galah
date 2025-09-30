package main

import (
	"math/rand"

	"github.com/gdamore/tcell/v2"
	"github.com/stvmln86/galah/galah/games/grid"
	"github.com/stvmln86/galah/galah/games/tile"
	"github.com/stvmln86/galah/galah/maths/pair"
	"github.com/stvmln86/galah/galah/maths/plot"
	"github.com/stvmln86/galah/galah/nodes/wall"
	"github.com/stvmln86/galah/galah/terms/term"
)

/*
TODOS:
- change grid.New to take different X/Y co-ordinates.
- add Pair.Random(), Pair.Sub(), Pair.SubXY().
*/

func main() {
	// initialise globals
	size := 25
	grid := grid.New(size)
	term, err := term.New()
	if err != nil {
		panic(err)
	}

	// draw random walls on grid
	for range size * 2 {
		pair := pair.New(rand.Intn(size-1)+1, rand.Intn(size-1)+1)
		grid.Set(pair, tile.New(" DD", wall.New("•WD")))
	}

	// draw border walls on grid
	for _, pair := range plot.Rectangle(pair.New(0, 0), pair.New(size-1, size-1)) {
		grid.Set(pair, tile.New(" DD", wall.New("#WD")))
	}

loop:
	for {
		// calculate pairs
		full := term.Size()
		head := pair.New((full.X-(size*2))/2, (full.Y-size)/2)
		foot := head.AddXY(0, size+1)

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
