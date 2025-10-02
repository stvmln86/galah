// Package draw implements terminal drawing functions.
package draw

import (
	"github.com/gdamore/tcell/v2"
	"github.com/stvmln86/galah/galah/elems/grid"
	"github.com/stvmln86/galah/galah/elems/icon"
	"github.com/stvmln86/galah/galah/tools/tone"
)

// Grid draws a Grid centered onto a Screen.
func Grid(scrn tcell.Screen, grid *grid.Grid) {
	wX, hY := scrn.Size()
	mX, mY := grid.Size*2, grid.Size
	oX, oY := (wX-mX)/2, (hY-mY)/2

	for i, icon := range grid.Icons() {
		x, y := i%grid.Size, i/grid.Size
		Icon(scrn, (x*2)+oX, y+oY, icon)
	}
}

// Icon draws an Icon onto a Screen.
func Icon(scrn tcell.Screen, x, y int, icon_ icon.Icon) {
	styl := icon.Style(icon_)
	scrn.SetContent(x, y, icon_.Char(), nil, styl)
}

// String draws a string onto a Screen.
func String(scrn tcell.Screen, x, y int, text string, tone tone.Colour) {
	styl := tcell.StyleDefault.Foreground(tone)
	for i, char := range text {
		scrn.SetContent(x+i, y, char, nil, styl)
	}
}

// StringPad draws a centered string onto a Screen.
func StringPad(scrn tcell.Screen, x, y int, text string, tone tone.Colour) {
	wX, _ := scrn.Size()
	mX := len(text)
	oX := (wX - mX) / 2

	for i, char := range text {
		String(scrn, oX+i, y, string(char), tone)
	}
}
