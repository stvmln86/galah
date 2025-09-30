package main

import (
	"math/rand"

	"github.com/gdamore/tcell/v2"
	"github.com/stvmln86/galah/galah/games/grid"
	"github.com/stvmln86/galah/galah/games/tile"
	"github.com/stvmln86/galah/galah/maths/pair"
	"github.com/stvmln86/galah/galah/maths/plot"
	"github.com/stvmln86/galah/galah/nodes/player"
	"github.com/stvmln86/galah/galah/nodes/wall"
	"github.com/stvmln86/galah/galah/terms/term"
)

/*
TODOS:
- add Pair.Random(), Pair.Sub(), Pair.SubXY().
- add grid.SetGen(pairs, func() *tile.Tile { tile.New(" DD", wall.New("#WD")) })
- add terms/keys to return all keys as strings ("a", "enter", etc).
*/

func main() {
	// initialise globals
	size := 25
	grid := grid.New(size)
	plyr := pair.New(size/2, size/2)
	term, err := term.New()
	if err != nil {
		panic(err)
	}

	// insert blank tiles into grid
	for _, pair := range plot.Block(pair.New(0, 0), pair.New(size-1, size-1)) {
		grid.Set(pair, tile.New(" DD"))
	}

	// insert random walls into grid
	for range size * 2 {
		pair := pair.New(rand.Intn(size-1)+1, rand.Intn(size-1)+1)
		grid.Get(pair).Push(wall.New("•HD"))
	}

	// insert border walls into grid
	for _, pair := range plot.Rectangle(pair.New(0, 0), pair.New(size-1, size-1)) {
		grid.Get(pair).Push(wall.New("#HD"))
	}

	// insert player into grid
	grid.Get(plyr).Push(player.New("@WD"))

loop:
	for {
		// calculate pairs
		full := term.Size()
		orig := pair.New((full.X-(size*2))/2, (full.Y-size)/2)
		head := orig.AddXY(0, -1)
		foot := head.AddXY(0, size+1)
		nbors := plot.Neighbours(plyr)

		// draw grid
		term.Clear()
		term.SetGrid(grid)
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

			switch evnt.Rune() {
			case 'w':
				orig := grid.Get(plyr)
				dest := grid.Get(nbors[0])
				if !dest.Wall() {
					dest.Push(orig.Pop())
					plyr = nbors[0]
				}
			}

		case *tcell.EventResize:
			continue loop
		}
	}

	// close screen
	term.Close()
}
