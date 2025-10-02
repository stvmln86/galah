// Package icon implements the Icon interface and functions.
package icon

import (
	"github.com/gdamore/tcell/v2"
	"github.com/stvmln86/galah/galah/tools/tone"
)

// Icon is an interface for terminal display characters.
type Icon interface {
	// Char returns the Icon's display character.
	Char() rune

	// Tone returns the Icon's foreground colour.
	Tone() tone.Colour
}

// Style returns a terminal style from an Icon.
func Style(icon Icon) tcell.Style {
	return tcell.StyleDefault.Foreground(icon.Tone())
}
