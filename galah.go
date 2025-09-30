package main

import (
	"time"

	"github.com/stvmln86/galah/galah/games/grid"
	"github.com/stvmln86/galah/galah/games/tile"
	"github.com/stvmln86/galah/galah/maths/pair"
	"github.com/stvmln86/galah/galah/maths/plot"
	"github.com/stvmln86/galah/galah/nodes/wall"
	"github.com/stvmln86/galah/galah/terms/term"
)

func main() {
	// initialise globals
	grid := grid.New(13)
	term, _ := term.New()

	// initialise gameworld objects
	wall := wall.New("#WD")
	for _, pair := range plot.Rectangle(pair.New(0, 0), pair.New(12, 12)) {
		grid.Set(pair, tile.New(" DD", wall))
	}

	// draw grid
	term.Clear()
	term.SetRender(grid.Render())
	term.SetString(pair.New(0, 1), "hello world", " BD")
	term.Show()

	// sleep and quit
	time.Sleep(3 * time.Second)
	term.Clear()
	term.Close()
}
