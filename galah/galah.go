package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/stvmln86/galah/galah/games/grid"
	"github.com/stvmln86/galah/galah/games/node"
	"github.com/stvmln86/galah/galah/nodes/base"
	"github.com/stvmln86/galah/galah/nodes/wall"
	"github.com/stvmln86/galah/galah/terms/cell"
	"github.com/stvmln86/galah/galah/tools/plot"
)

var (
	// cells
	cellBase = cell.New(' ', 'D')
	cellPlyr = cell.New('@', 'B')
	cellWall = cell.New('#', 'H')

	// nodes
	nodePlayer   = base.New(cellPlyr, "player", true)
	nodeBaseFunc = func() node.Node { return base.New(cellBase, "", true) }
	nodeWallFunc = func() node.Node { return wall.New(cellWall, "wall") }

	// variables
	pX, pY = 10, 10
)

func try(scrn tcell.Screen, err error) {
	if err != nil {
		scrn.Fini()
		panic(err)
	}
}

func main() {
	// initialise display
	scrn, err := tcell.NewScreen()
	try(scrn, err)
	try(scrn, scrn.Init())

	// initialise grid
	grid := grid.New(20, nodeBaseFunc)

	// generate walls on grid
	for _, pairs := range plot.Rectangle(0, 0, 19, 19) {
		grid.SetNode(pairs[0], pairs[1], nodeWallFunc())
	}

	// generate entities on grid
	grid.SetNode(pX, pY, nodePlayer)

	// enter main loop
loop:
	for {
		// prepare screen
		wX, hY := scrn.Size()
		mX, mY := grid.Size*2, grid.Size
		oX, oY := (wX-mX)/2, (hY-mY)/2
		scrn.Clear()

		// draw grid
		for i, cell := range grid.Cells() {
			if cell != nil {
				x, y := i%grid.Size, i/grid.Size
				scrn.SetContent((x*2)+oX, y+oY, cell.Rune(), nil, cell.Style())
			}
		}

		// draw text
		for i, char := range "hello world" {
			scrn.SetContent(oX+i, oY+grid.Size, char, nil, tcell.StyleDefault)
		}

		// flip screen
		scrn.Show()

		// enter event poll
		switch e := scrn.PollEvent().(type) {
		case *tcell.EventKey:
			switch e.Key() {
			case tcell.KeyEscape:
				break loop
			case tcell.KeyUp:
				dest := plot.Neighbours(pX, pY)[0]
				node := grid.GetNode(dest[0], dest[1])
				if grid.In(dest[0], dest[1]) && node.Open() {
					grid.SetNode(dest[0], dest[1], nodePlayer)
					grid.SetNode(pX, pY, node)
					pX, pY = dest[0], dest[1]
				}
			}
		}
	}

	// close screen
	scrn.Fini()
}
