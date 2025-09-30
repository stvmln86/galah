// Package flag implements the Flag type and methods.
package flag

import (
	"github.com/gdamore/tcell/v2"
)

// Flag is a three-character string containing a rune, foreground and background.
type Flag string

// Colours is a map of all defined TCell colours.
var Colours = map[rune]tcell.Color{
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

// Back returns the Flag's background colour.
func (f Flag) Back() tcell.Color {
	return Colours[rune(f[2])]
}

// Fore returns the Flag's foreground colour.
func (f Flag) Fore() tcell.Color {
	return Colours[rune(f[1])]
}

// Rune returns the Flag's drawing rune.
func (f Flag) Rune() rune {
	return rune(f[0])
}
