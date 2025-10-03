///////////////////////////////////////////////////////////////////////////////////////
//                  galah · a maze runner in Go · by Stephen Malone                  //
///////////////////////////////////////////////////////////////////////////////////////

package main

import (
	"math/rand"

	"github.com/gdamore/tcell/v2"
)

///////////////////////////////////////////////////////////////////////////////////////
//                          part one · constants and globals                         //
///////////////////////////////////////////////////////////////////////////////////////

// 1.1 · global math constants
///////////////////////////////

// GridArea is the total area of the MainGrid.
const GridArea = GridSize * GridSize

// GridEdge is the furthest edge co-ordinate of the MainGrid.
const GridEdge = GridSize - 1

// GridSize is the height and width of the MainGrid.
const GridSize = 100

// 1.2 · global system objects
///////////////////////////////

// MainGrid is the main gameworld grid.
var MainGrid [GridArea]rune

// MainScreen is the main terminal display controller.
var MainScreen tcell.Screen

// 1.3 · terminal lookup tables
////////////////////////////////

// KeyNames is a map of all defined terminal keys and their names.
var KeyNames = map[tcell.Key]string{
	tcell.KeyUp:    "up",
	tcell.KeyDown:  "down",
	tcell.KeyLeft:  "left",
	tcell.KeyRight: "right",
}

// RuneTones is a map of all defined runes and their foreground colours.
var RuneTones = map[rune]tcell.Color{
	'@': tcell.ColorBlue,
	'#': tcell.ColorGrey,
	'·': tcell.ColorYellow,
}

///////////////////////////////////////////////////////////////////////////////////////
//                       part two · terminal control functions                       //
///////////////////////////////////////////////////////////////////////////////////////

// DrawMainGrid draws the MainGrid onto the MainScreen.
func DrawMainGrid(px, py int) {
	wx, hy := MainScreen.Size()
	sx, sy := px-wx/4, py-hy/2

	for y := range hy {
		for x := range wx / 2 {
			gx, gy := sx+x, sy+y
			if gx >= 0 && gx < GridSize && gy >= 0 && gy < GridSize {
				rune := MainGrid[gy*GridSize+gx]
				DrawRune(x*2, y, rune, RuneTones[rune])
			}
		}
	}
}

// DrawRune draws a rune onto the MainScreen.
func DrawRune(x, y int, rune rune, tone tcell.Color) {
	styl := tcell.StyleDefault.Foreground(tone)
	MainScreen.SetContent(x, y, rune, nil, styl)
}

// DrawText draws a string onto the MainScreen.
func DrawText(x, y int, text string, tone tcell.Color) {
	for i, rune := range text {
		DrawRune(x+i, y, rune, tone)
	}
}

// PollEvent returns a new MainScreen event string.
func PollEvent() string {
	switch evnt := MainScreen.PollEvent().(type) {
	case *tcell.EventKey:
		if name, ok := KeyNames[evnt.Key()]; ok {
			return name
		} else {
			return string(evnt.Rune())
		}
	default:
		return "unknown"
	}
}

///////////////////////////////////////////////////////////////////////////////////////
//                             part ??? · main functions                             //
///////////////////////////////////////////////////////////////////////////////////////

// main runs the main Galah program.
func main() {
	// Create and initialise MainScreen.
	MainScreen, _ = tcell.NewScreen()
	MainScreen.Init()

	// Create and populated MainGrid.
	MainGrid = [GridArea]rune{}
	for i := range GridArea {
		x, y := i%GridSize, i/GridSize

		switch {
		case x == 0 || x == GridEdge || y == 0 || y == GridEdge:
			MainGrid[i] = '#'
		case rand.Intn(100) < 10:
			MainGrid[i] = '·'
		default:
			MainGrid[i] = ' '
		}
	}

	// Enter main drawing loop.
loop:
	for {
		// Draw screen.
		MainScreen.Clear()
		DrawMainGrid(50, 50)
		DrawText(0, 0, "hello world", tcell.ColorGreen)
		MainScreen.Show()

		// Poll for event.
		switch PollEvent() {
		case "q":
			break loop
		}
	}

	// Close screen.
	MainScreen.Fini()
}
