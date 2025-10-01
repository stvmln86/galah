// Package cell implements the Cell type and methods.
package cell

import (
	"github.com/gdamore/tcell/v2"
)

// Cell is a drawable rune with a foreground colour.
type Cell struct {
	Rune rune
	Tone rune
}

// tones is a map of all defined terminal colours.
var tones = map[rune]tcell.Color{
	'A': tcell.ColorAqua,
	'B': tcell.ColorBlue,
	'D': tcell.ColorDefault,
	'F': tcell.ColorFuchsia,
	'G': tcell.ColorGreen,
	'H': tcell.ColorGray,
	'K': tcell.ColorBlack,
	'L': tcell.ColorLime,
	'M': tcell.ColorMaroon,
	'N': tcell.ColorNavy,
	'O': tcell.ColorOlive,
	'P': tcell.ColorPurple,
	'R': tcell.ColorRed,
	'S': tcell.ColorSilver,
	'T': tcell.ColorTeal,
	'W': tcell.ColorWhite,
	'Y': tcell.ColorYellow,
}

// New returns a new Char.
func New(rune, tone rune) *Cell {
	return &Cell{rune, tone}
}

// Style returns the Char's terminal style.
func (c *Cell) Style() tcell.Style {
	return tcell.StyleDefault.Foreground(tones[c.Tone])
}
