// Package main runs the main Galah program.
package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/stvmln86/galah/galah/elems/grid"
	"github.com/stvmln86/galah/galah/elems/node"
	"github.com/stvmln86/galah/galah/icons/baseicon"
	"github.com/stvmln86/galah/galah/nodes/basenode"
	"github.com/stvmln86/galah/galah/tools/draw"
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
		draw.Grid(mainScreen, mainGrid)
		draw.String(mainScreen, 0, 0, "hello world", tone.Blue)
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
