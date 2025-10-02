// Package main runs the main Galah program.
package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/stvmln86/galah/galah/elems/grid"
	"github.com/stvmln86/galah/galah/elems/icon"
	"github.com/stvmln86/galah/galah/elems/node"
	"github.com/stvmln86/galah/galah/icons/baseicon"
	"github.com/stvmln86/galah/galah/nodes/basenode"
	"github.com/stvmln86/galah/galah/tools/plot"
	"github.com/stvmln86/galah/galah/tools/poll"
	"github.com/stvmln86/galah/galah/tools/tone"
)

var (
	// Icon definitions.
	iconFloor  = baseicon.New(' ', tone.Default)
	iconWall   = baseicon.New('#', tone.Grey)
	iconPlayer = baseicon.New('@', tone.Blue)

	// Node definitions.
	nodeFloor  = func() node.Node { return basenode.New(iconFloor, "floor", true) }
	nodeWall   = func() node.Node { return basenode.New(iconWall, "wall", true) }
	nodePlayer = basenode.New(iconPlayer, "player", true)

	// Global variable definitions.
	mainScreen       tcell.Screen
	mainGrid         = grid.New(21, nodeFloor)
	playerX, playerY = 10, 10
)

func drawGrid(scrn tcell.Screen, grid *grid.Grid) {
	wX, hY := scrn.Size()
	mX, mY := grid.Size*2, grid.Size
	oX, oY := (wX-mX)/2, (hY-mY)/2

	for i, icon_ := range grid.Icons() {
		if icon_ != nil {
			x, y := i%grid.Size, i/grid.Size
			scrn.SetContent((x*2)+oX, y+oY, icon_.Char(), nil, icon.Style(icon_))
		}
	}
}

func drawText(scrn tcell.Screen, x, y int, text string, tone tone.Colour) {
	styl := tcell.StyleDefault.Foreground(tone)

	for i, char := range text {
		scrn.SetContent(x+i, y, char, nil, styl)
	}
}

// main runs the main Galah program.
func main() {
	// Initialise the screen.
	mainScreen, _ = tcell.NewScreen()
	mainScreen.Init()

	// Generate grid content.
	walls := plot.Rectangle(0, 0, mainGrid.Size-1, mainGrid.Size-1)
	mainGrid.SetFunc(walls, nodeWall)
	mainGrid.Set(playerX, playerY, nodePlayer)

	// Enter main interactive loop.
loop:
	for {
		// Draw grid on screen.
		mainScreen.Clear()
		drawGrid(mainScreen, mainGrid)
		drawText(mainScreen, 0, 0, "hello world", tone.Blue)
		mainScreen.Show()

		// Receive poll event.
		switch poll.Poll(mainScreen) {
		case "escape":
			break loop
		case "resize":
			continue loop
		}
	}

	// Close the screen.
	mainScreen.Fini()
}
